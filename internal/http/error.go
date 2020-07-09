package http

import "net/http"

// Error is an error wrapper for the standard HTTP error codes
type Error int

func (e Error) Error() string {
	return http.StatusText(int(e))
}

func (e Error) Code() int {
	return int(e)
}
