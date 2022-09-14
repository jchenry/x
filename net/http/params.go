package http

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PathParam(ctx context.Context, Param func(ctx context.Context, paramName string) string, paramName string, required bool, dt string) (p any, err error) {
	s := Param(ctx, paramName)
	if s == "" && required {
		return nil, errors.New("missing required parameter")
	}

	switch dt {
	case "int64":
		p, err = strconv.ParseInt(s, 10, 64)
	case "int32":
		var x int64
		x, err = strconv.ParseInt(s, 10, 32)
		p = int32(x)
	case "string":
		p = s
	default:
		err = errors.New("no match for type")
	}

	return
}

func BodyParam(body io.ReadCloser, p any, v func(p any) error) (err error) {
	d := json.NewDecoder(body)
	if err = d.Decode(p); err == nil {
		err = v(p)
	}
	return
}

func mappedParam(m map[string][]string, paramName string, required bool, dt string) (p any, err error) {

	var s string
	q, exists := m[paramName]
	if !exists { // intentionally left empty
	} else if len(q) > 1 {
		s = strings.Join(q, ",")
	} else {
		s = q[0]
	}

	if s == "" && required {
		return nil, errors.New("missing required parameter")
	}

	switch dt {
	case "int64":
		p, err = strconv.ParseInt(s, 10, 64)
	case "int32":
		var x int64
		x, err = strconv.ParseInt(s, 10, 32)
		p = int32(x)
	case "bool":
		var b bool
		b, err = strconv.ParseBool(s)
		p = bool(b)
	case "string":
		p = s
	case "[]int64":
		str := strings.Split(s, ",")
		ints := make([]int64, len(str))
		for i, s := range str {
			if v, err := strconv.ParseInt(s, 10, 64); err != nil {
				return nil, err
			} else {
				ints[i] = v
			}
		}
		p = ints
	case "[]int32":
		str := strings.Split(s, ",")
		ints := make([]int32, len(str))
		for i, s := range str {
			if v, err := strconv.ParseInt(s, 10, 32); err != nil {
				return nil, err
			} else {
				ints[i] = int32(v)
			}
		}
		p = ints
	case "[]string":
		p = strings.Split(s, ",")
	default:
		err = errors.New("no match for type")
	}
	return
}

func QueryParam(query url.Values, paramName string, required bool, dt string) (p any, err error) {
	return mappedParam(query, paramName, required, dt)
}

func HeaderParam(h http.Header, paramName string, required bool, dt string) (p any, err error) {
	return mappedParam(h, paramName, required, dt)
}

func FormParam(form url.Values, paramName string, required bool, dt string) (p any, err error) {
	return mappedParam(form, paramName, required, dt)
}
