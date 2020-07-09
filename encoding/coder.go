package encoding

import "io"

type Encoder func(io.Writer, interface{}) error
type Decoder func(io.Reader, interface{}) error
