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
	t.Run("test int64 parse", func(t *testing.T) {
		var p int64
		err := PathParam(context.WithValue(context.Background(), contextKey("int64id"), "123"), Param, &p, "int64id", true)
		if (err != nil) != false {
			t.Errorf("PathParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, int64(123)) {
			t.Errorf("PathParam() = %v, want %v", p, int64(123))
		}
	})

	t.Run("test int32 parse", func(t *testing.T) {
		var p int32
		err := PathParam(context.WithValue(context.Background(), contextKey("int32id"), "123"), Param, &p, "int32id", true)
		if (err != nil) != false {
			t.Errorf("PathParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, int32(123)) {
			t.Errorf("PathParam() = %v, want %v", p, int32(123))
		}
	})

	t.Run("test string parse", func(t *testing.T) {
		var p string
		err := PathParam(
			context.WithValue(
				context.Background(),
				contextKey("stringid"),
				"foo"),
			Param,
			&p,
			"stringid",
			true)
		if (err != nil) != false {
			t.Errorf("PathParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, "foo") {
			t.Errorf("PathParam() = %v, want %v", p, "foo")
		}
	})

	t.Run("test missing required parameter", func(t *testing.T) {
		var p string
		err := PathParam(context.Background(), Param, &p, "stringid", true)
		if (err != nil) != true {
			t.Errorf("PathParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, "") {
			t.Errorf("PathParam() = %v, want %v", p, "")
		}
	})

	t.Run("test unknown type parameter", func(t *testing.T) {
		var p complex64
		err := PathParam(context.WithValue(context.Background(), contextKey("stringid"), "foo"),
			Param, p, "stringid", true)
		if (err != nil) != true {
			t.Errorf("PathParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, complex64(0)) {
			t.Errorf("PathParam() = %v, want %v", p, complex64(0))
		}
	})
}

func toValues(s string) url.Values {
	v, _ := url.ParseQuery(s)
	return v
}

func TestMappedParam(t *testing.T) {
	t.Run("test int64 parse", func(t *testing.T) {
		var p int64
		err := mappedParam(toValues("x=123"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, int64(123)) {
			t.Errorf("QueryParam() = %v, want %v", p, int64(123))
		}
	})

	t.Run("test int32 parse", func(t *testing.T) {
		var p int32
		err := mappedParam(toValues("x=123"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, int32(123)) {
			t.Errorf("QueryParam() = %v, want %v", p, int32(123))
		}
	})

	t.Run("test bool parse", func(t *testing.T) {
		var p bool
		err := mappedParam(toValues("x=true"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, true) {
			t.Errorf("QueryParam() = %v, want %v", p, true)
		}
	})

	t.Run("test string parse", func(t *testing.T) {
		var p string
		err := mappedParam(toValues("x=foobar"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, "foobar") {
			t.Errorf("QueryParam() = %v, want %v", p, "foobar")
		}
	})

	t.Run("test []int64 parse", func(t *testing.T) {
		var p []int64
		err := mappedParam(toValues("x=123&x=456"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, []int64{int64(123), int64(456)}) {
			t.Errorf("QueryParam() = %v, want %v", p, []int64{int64(123), int64(456)})
		}
	})

	t.Run("test []int64 bad parse", func(t *testing.T) {
		var p []int64
		err := mappedParam(toValues("x=123&x=4q56"), "x", &p, true)
		if (err != nil) != true {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, []int64{123}) {
			t.Errorf("QueryParam() = %v, want %v", p, []int64{})
		}
	})

	t.Run("test []int32 parse", func(t *testing.T) {
		var p []int32
		err := mappedParam(toValues("x=123&x=456"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, []int32{int32(123), int32(456)}) {
			t.Errorf("QueryParam() = %v, want %v", p, []int32{int32(123), int32(456)})
		}
	})

	t.Run("test []int32 bad parse", func(t *testing.T) {
		var p []int32
		err := mappedParam(toValues("x=123&x=4q56"), "x", &p, true)
		if (err != nil) != true {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, []int32{123}) {
			t.Errorf("QueryParam() = %v, want %v", p, []int32{123})
		}
	})

	t.Run("test []string parse", func(t *testing.T) {
		var p []string
		err := mappedParam(toValues("x=foo&x=bar"), "x", &p, true)
		if (err != nil) != false {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, false)
			return
		}
		if !reflect.DeepEqual(p, []string{"foo", "bar"}) {
			t.Errorf("QueryParam() = %v, want %v", p, []string{"foo", "bar"})
		}
	})

	t.Run("test missing required parameter", func(t *testing.T) {
		var p string
		err := mappedParam(toValues("y=hello"), "x", &p, true)
		if (err != nil) != true {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, "") {
			t.Errorf("QueryParam() = %v, want %v", p, "")
		}
	})

	t.Run("test unknown type parameter", func(t *testing.T) {
		var p complex64
		err := mappedParam(toValues("x=hello"), "x", &p, true)
		if (err != nil) != true {
			t.Errorf("QueryParam() error = %v, wantErr %v", err, true)
			return
		}
		if !reflect.DeepEqual(p, complex64(0)) {
			t.Errorf("QueryParam() = %v, want %v", p, complex64(0))
		}
	})
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

func Ptr[T any](v T) *T {
	return &v
}

func NilString() *string {
	return nil
}
