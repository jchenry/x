package json

import (
	"encoding/json"
	"io"
)

func Encoder(w io.Writer, e interface{}) error {
	return json.NewEncoder(w).Encode(e)
}

func Decoder(r io.Reader, e interface{}) error {
	return json.NewDecoder(r).Decode(e)
}
