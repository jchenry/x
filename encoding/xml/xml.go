package xml

import (
	"encoding/xml"
	"io"
)

func Encoder(w io.Writer, e interface{}) error {
	return xml.NewEncoder(w).Encode(e)
}

func Decoder(r io.Reader, e interface{}) error {
	return xml.NewDecoder(r).Decode(e)
}
