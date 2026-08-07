package json

import (
	"github.com/jianlu8023/go-tools/v2/pkg/json/jsoniter"
	"io"
)

type jsoniterAPI struct{}

func (j *jsoniterAPI) Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}

func (j *jsoniterAPI) MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return jsoniter.MarshalIndent(v, prefix, indent)
}

func (j *jsoniterAPI) NewEncoder(w io.Writer) Encoder {
	enc := jsoniter.NewEncoder(w)
	return &encoderWrapper{encode: enc.Encode}
}

func (j *jsoniterAPI) Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

func (j *jsoniterAPI) NewDecoder(r io.Reader) Decoder {
	dec := jsoniter.NewDecoder(r)
	return &decoderWrapper{decode: dec.Decode}
}

func (j *jsoniterAPI) MarshalPretty(v interface{}) ([]byte, error) {
	return jsoniter.MarshalPretty(v)
}

func (j *jsoniterAPI) MarshalString(v interface{}) (string, error) {
	return jsoniter.MarshalString(v)
}

func (j *jsoniterAPI) UnmarshalString(str string, v interface{}) error {
	return jsoniter.UnmarshalString(str, v)
}

func (j *jsoniterAPI) ReadFromFile(filePath string, v interface{}) error {
	return jsoniter.ReadFromFile(filePath, v)
}

func (j *jsoniterAPI) WriteToFile(filePath string, v interface{}) error {
	return jsoniter.WriteToFile(filePath, v)
}

func (j *jsoniterAPI) Validate(jsonStr string) bool {
	return jsoniter.Validate(jsonStr)
}

func (j *jsoniterAPI) Compact(jsonStr string) (string, error) {
	return jsoniter.Compact(jsonStr)
}
