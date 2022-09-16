package http

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func PathParam(ctx context.Context, Param func(ctx context.Context, paramName string) string, p interface{}, paramName string, required bool) (err error) {
	s := Param(ctx, paramName)
	if s == "" && required {
		switch v := p.(type) {
		case *string:
			p = v
			p = nil
		}
		return errors.New("missing required parameter")
	}

	switch v := p.(type) {
	case *int64:
		*v, err = strconv.ParseInt(s, 10, 64)
	case *int32:
		var x int64
		x, err = strconv.ParseInt(s, 10, 32)
		*v = int32(x)
	case *string:
		*v = s
	default:
		err = fmt.Errorf("no match for pointer type %T", v)
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

func mappedParam(m map[string][]string, paramName string, p interface{}, required bool) (err error) {
	var s string
	q, exists := m[paramName]
	if !exists { // intentionally left empty
	} else if len(q) > 1 {
		s = strings.Join(q, ",")
	} else {
		s = q[0]
	}

	if s == "" && required {
		return errors.New("missing required parameter")
	}

	switch v := p.(type) {
	case *int64:
		*v, err = strconv.ParseInt(s, 10, 64)
	case *int32:
		var x int64
		x, err = strconv.ParseInt(s, 10, 32)
		*v = int32(x)
	case *bool:
		*v, err = strconv.ParseBool(s)
	case *string:
		*v = s
	case *[]int64:
		str := strings.Split(s, ",")
		for _, s := range str {
			if e, err := strconv.ParseInt(s, 10, 64); err != nil {
				return err
			} else {
				*v = append(*v, e)
				// ints[i] = e
			}
		}
	case *[]int32:
		str := strings.Split(s, ",")
		for _, s := range str {
			if e, err := strconv.ParseInt(s, 10, 32); err != nil {
				return err
			} else {
				*v = append(*v, int32(e))
				// ints[i] = e
			}
		}
	case *[]string:
		*v = strings.Split(s, ",")
	default:
		err = fmt.Errorf("no match for pointer type %T", v)
	}

	return
}

func QueryParam(query url.Values, paramName string, p interface{}, required bool) (err error) {
	return mappedParam(query, paramName, p, required)
}

func HeaderParam(h http.Header, paramName string, p interface{}, required bool) (err error) {
	return mappedParam(h, paramName, p, required)
}

func FormParam(form url.Values, paramName string, p interface{}, required bool) (err error) {
	return mappedParam(form, paramName, p, required)
}
