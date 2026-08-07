package json

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json/sonic"
	"io"
)

type encoderWrapper struct {
	encode func(v interface{}) error
}

func (e *encoderWrapper) Encode(v interface{}) error {
	return e.encode(v)
}

type decoderWrapper struct {
	decode func(v interface{}) error
}

func (d *decoderWrapper) Decode(v interface{}) error {
	return d.decode(v)
}

type sonicAPI struct {
	*sonic.Sonic
}

func (s *sonicAPI) NewEncoder(w io.Writer) Encoder {
	enc := s.Sonic.NewEncoder(w)
	return &encoderWrapper{encode: enc.Encode}
}

func (s *sonicAPI) NewDecoder(r io.Reader) Decoder {
	dec := s.Sonic.NewDecoder(r)
	return &decoderWrapper{decode: dec.Decode}
}
