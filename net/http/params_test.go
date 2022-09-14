package http

import (
	"context"
	"errors"
	"io"
	"net/url"
	"reflect"
	"strings"
	"testing"
)

func TestPathParam(t *testing.T) {
	type args struct {
		ctx       context.Context
		Param     func(ctx context.Context, paramName string) string
		paramName string
		required  bool
		dt        string
	}
	tests := []struct {
		name    string
		args    args
		wantP   any
		wantErr bool
	}{
		{
			name: "test int64 parse",
			args: args{
				context.WithValue(context.Background(), contextKey("int64id"), "123"),
				Param,
				"int64id",
				true,
				"int64",
			},
			wantP:   int64(123),
			wantErr: false,
		},
		{
			name: "test int32 parse",
			args: args{
				context.WithValue(context.Background(), contextKey("int32id"), "123"),
				Param,
				"int32id",
				true,
				"int32",
			},
			wantP:   int32(123),
			wantErr: false,
		},
		{
			name: "test string parse",
			args: args{
				context.WithValue(context.Background(), contextKey("stringid"), "foo"),
				Param,
				"stringid",
				true,
				"string",
			},
			wantP:   string("foo"),
			wantErr: false,
		},
		{
			name: "test missing required parameter",
			args: args{
				context.WithValue(context.Background(), contextKey("stringid"), ""),
				Param,
				"stringid",
				true,
				"string",
			},
			wantP:   nil,
			wantErr: true,
		},
		{
			name: "test unknown type parameter",
			args: args{
				context.WithValue(context.Background(), contextKey("stringid"), "foo"),
				Param,
				"stringid",
				true,
				"not_a_real_type",
			},
			wantP:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := PathParam(tt.args.ctx, tt.args.Param, tt.args.paramName, tt.args.required, tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("PathParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("PathParam() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func toValues(s string) url.Values {
	v, _ := url.ParseQuery(s)
	return v
}
func TestMappedParam(t *testing.T) {

	type args struct {
		query     url.Values
		paramName string
		required  bool
		dt        string
	}
	tests := []struct {
		name    string
		args    args
		wantP   any
		wantErr bool
	}{
		{
			name: "test int64 parse",
			args: args{
				toValues("x=123"),
				"x",
				true,
				"int64",
			},
			wantP:   int64(123),
			wantErr: false,
		},
		{
			name: "test int32 parse",
			args: args{
				toValues("x=123"),
				"x",
				true,
				"int32",
			},
			wantP:   int32(123),
			wantErr: false,
		},
		{
			name: "test bool parse",
			args: args{
				toValues("x=true"),
				"x",
				true,
				"bool",
			},
			wantP:   bool(true),
			wantErr: false,
		},
		{
			name: "test string parse",
			args: args{
				toValues("x=foobar"),
				"x",
				true,
				"string",
			},
			wantP:   string("foobar"),
			wantErr: false,
		},
		{
			name: "test []int64 parse",
			args: args{
				toValues("x=123&x=456"),
				"x",
				true,
				"[]int64",
			},
			wantP:   []int64{int64(123), int64(456)},
			wantErr: false,
		},
		{
			name: "test []int64 bad parse",
			args: args{
				toValues("x=123&x=4q56"),
				"x",
				true,
				"[]int64",
			},
			wantP:   nil,
			wantErr: true,
		},
		{
			name: "test []int32 parse",
			args: args{
				toValues("x=123&x=456"),
				"x",
				true,
				"[]int32",
			},
			wantP:   []int32{int32(123), int32(456)},
			wantErr: false,
		},
		{
			name: "test []int32 bad parse",
			args: args{
				toValues("x=123&x=4q56"),
				"x",
				true,
				"[]int32",
			},
			wantP:   nil,
			wantErr: true,
		},
		{
			name: "test []string parse",
			args: args{
				toValues("x=foo&x=bar"),
				"x",
				true,
				"[]string",
			},
			wantP:   []string{"foo", "bar"},
			wantErr: false,
		},
		{
			name: "test missing required parameter",
			args: args{
				toValues("y=hello"),
				"x",
				true,
				"string",
			},
			wantP:   nil,
			wantErr: true,
		},
		{
			name: "test unknown type parameter",
			args: args{
				toValues("x=hello"),
				"x",
				true,
				"not_a_real_type",
			},
			wantP:   nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotP, err := mappedParam(tt.args.query, tt.args.paramName, tt.args.required, tt.args.dt)
			if (err != nil) != tt.wantErr {
				t.Errorf("QueryParam() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotP, tt.wantP) {
				t.Errorf("QueryParam() = %v, want %v", gotP, tt.wantP)
			}
		})
	}
}

func TestBodyParam(t *testing.T) {
	type args struct {
		body io.ReadCloser
		p    any
		v    func(p any) error
	}

	type x struct{}
	var y x

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "test happy path",
			args: args{
				body: io.NopCloser(strings.NewReader("{}")),
				p:    &y,
				v: func(p any) error {
					return nil
				},
			},
			wantErr: false,
		},
		{
			name: "test bad json",
			args: args{
				body: io.NopCloser(strings.NewReader("}")),
				p:    &y,
				v: func(p any) error {
					return nil
				},
			},
			wantErr: true,
		},
		{
			name: "test validation failed",
			args: args{
				body: io.NopCloser(strings.NewReader("{}")),
				p:    &y,
				v: func(p any) error {
					return errors.New("validation failed")
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BodyParam(tt.args.body, tt.args.p, tt.args.v); (err != nil) != tt.wantErr {
				t.Errorf("BodyParam() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
