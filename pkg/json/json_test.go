package json

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDefaultAPI(t *testing.T) {
	data := map[string]interface{}{
		"name": "John",
		"age":  30,
	}

	bytes, err := Marshal(data)
	assert.NoError(t, err)
	assert.NotEmpty(t, bytes)

	var decoded map[string]interface{}
	err = Unmarshal(bytes, &decoded)
	assert.NoError(t, err)
	assert.Equal(t, "John", decoded["name"])
}

func TestAPIInterface(t *testing.T) {
	data := map[string]interface{}{
		"name": "Jane",
		"age":  25,
	}

	apis := []API{
		NewSonicAPI(),
		NewJsoniterAPI(),
	}

	for _, api := range apis {
		bytes, err := api.Marshal(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, bytes)

		var decoded map[string]interface{}
		err = api.Unmarshal(bytes, &decoded)
		assert.NoError(t, err)
		assert.Equal(t, "Jane", decoded["name"])

		pretty, err := api.MarshalPretty(data)
		assert.NoError(t, err)
		assert.Contains(t, string(pretty), "\n")

		jsonStr, err := api.MarshalString(data)
		assert.NoError(t, err)
		assert.Contains(t, jsonStr, "name")

		assert.True(t, api.Validate(jsonStr))
		assert.False(t, api.Validate("invalid json"))
	}
}

func TestSwitchAPI(t *testing.T) {
	data := map[string]interface{}{
		"key": "value",
	}

	UseSonic()
	bytes1, err := Marshal(data)
	assert.NoError(t, err)

	UseJsoniter()
	bytes2, err := Marshal(data)
	assert.NoError(t, err)

	assert.Equal(t, string(bytes1), string(bytes2))

	UseSonic()
}

func TestEncoderDecoder(t *testing.T) {
	data := map[string]interface{}{
		"name": "Test",
	}

	apis := []API{
		NewSonicAPI(),
		NewJsoniterAPI(),
	}

	for _, api := range apis {
		var buf []byte
		enc := api.NewEncoder(&mockWriter{buf: &buf})
		err := enc.Encode(data)
		assert.NoError(t, err)
		assert.NotEmpty(t, buf)
	}
}

type mockWriter struct {
	buf *[]byte
}

func (w *mockWriter) Write(p []byte) (n int, err error) {
	*w.buf = append(*w.buf, p...)
	return len(p), nil
}
