package json

import (
	"bytes"
	"encoding/json"
	"io"

	"github.com/jianlu8023/go-tools/v2/pkg/json/sonic"
)

type API interface {
	Marshal(v interface{}) ([]byte, error)
	MarshalIndent(v interface{}, prefix, indent string) ([]byte, error)
	NewEncoder(w io.Writer) Encoder
	Unmarshal(data []byte, v interface{}) error
	NewDecoder(r io.Reader) Decoder
	MarshalPretty(v interface{}) ([]byte, error)
	MarshalString(v interface{}) (string, error)
	UnmarshalString(str string, v interface{}) error
	ReadFromFile(filePath string, v interface{}) error
	WriteToFile(filePath string, v interface{}) error
	Validate(jsonStr string) bool
	Compact(jsonStr string) (string, error)
}

type Encoder interface {
	Encode(v interface{}) error
}

type Decoder interface {
	Decode(v interface{}) error
}

func NewSonicAPI() API {
	return &sonicAPI{Sonic: sonic.NewStandardSonic()}
}

func NewJsoniterAPI() API {
	return &jsoniterAPI{}
}

var defaultAPI = NewSonicAPI()

func Marshal(v interface{}) ([]byte, error) {
	return defaultAPI.Marshal(v)
}

func MarshalIndent(v interface{}, prefix, indent string) ([]byte, error) {
	return defaultAPI.MarshalIndent(v, prefix, indent)
}

func NewEncoder(w io.Writer) Encoder {
	return defaultAPI.NewEncoder(w)
}

func Unmarshal(data []byte, v interface{}) error {
	return defaultAPI.Unmarshal(data, v)
}

func NewDecoder(r io.Reader) Decoder {
	return defaultAPI.NewDecoder(r)
}

func MarshalPretty(v interface{}) ([]byte, error) {
	return defaultAPI.MarshalPretty(v)
}

func MarshalString(v interface{}) (string, error) {
	return defaultAPI.MarshalString(v)
}

func UnmarshalString(str string, v interface{}) error {
	return defaultAPI.UnmarshalString(str, v)
}

func ReadFromFile(filePath string, v interface{}) error {
	return defaultAPI.ReadFromFile(filePath, v)
}

func WriteToFile(filePath string, v interface{}) error {
	return defaultAPI.WriteToFile(filePath, v)
}

func Validate(jsonStr string) bool {
	return defaultAPI.Validate(jsonStr)
}

func Compact(jsonStr string) (string, error) {
	return defaultAPI.Compact(jsonStr)
}

func UseSonic() {
	defaultAPI = NewSonicAPI()
}

func UseJsoniter() {
	defaultAPI = NewJsoniterAPI()
}

func IsJsonArray(str string) bool {
	var js []interface{}
	return defaultAPI.Unmarshal([]byte(str), &js) == nil
}

func IsJsonObject(str string) bool {
	var js map[string]interface{}
	return defaultAPI.Unmarshal([]byte(str), &js) == nil
}

func GetJsonType(data json.RawMessage) string {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 {
		return "unknown"
	}
	firstChar := trimmed[0]
	switch firstChar {
	case '{':
		return "object"
	case '[':
		return "array"
	case '"':
		return "string"
	case 't', 'f':
		return "boolean"
	case 'n':
		return "null"
	default:
		return "number"
	}
}
