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

// func Decoder(get func() interface{}) func(io.Reader) (interface{}, error) {
// 	//TODO I dont like the get() function, find a better way of dealing with this
// 	return func(r io.Reader) (interface{}, error) {
// 		e := get()
// 		err := json.NewDecoder(r).Decode(e)
// 		return e, err
// 	}
// }
