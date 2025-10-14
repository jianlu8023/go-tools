package pretty

import (
	"bytes"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jianlu8023/go-tools/v2/internal/colour"
	"github.com/jianlu8023/go-tools/v2/pkg/sonic"
)

// Formatter is a struct to format JSON data.
type Formatter struct {
	// Max length of JSON string value. When the value is 1 and over, string is truncated to length of the value.
	// Default is 0 (not truncated).
	StringMaxLength int

	// Boolean to disable color. Default is false.
	DisabledColor bool

	// Indent space number. Default is 2.
	Indent int

	// Newline string. To print without new lines set it to empty string. Default is \n.
	Newline string
}

// NewFormatter returns a new formatter with following default values.
func NewFormatter() *Formatter {
	return &Formatter{
		StringMaxLength: 0,
		DisabledColor:   false,
		Indent:          2,
		Newline:         "\n",
	}
}

// Marshal marshals and formats JSON data.
func (f *Formatter) Marshal(v interface{}) ([]byte, error) {
	data, err := sonic.Marshal(v)
	if err != nil {
		return nil, err
	}
	return f.Format(data)
}

// Format formats JSON string.
func (f *Formatter) Format(data []byte) ([]byte, error) {
	var v interface{}
	decoder := sonic.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	if err := decoder.Decode(&v); err != nil {
		return nil, err
	}

	return []byte(f.pretty(v, 1)), nil
}

func (f *Formatter) pretty(v interface{}, depth int) string {
	switch val := v.(type) {
	case string:
		return f.processString(val)
	case float64:
		return f.sprintFgCyan(strconv.FormatFloat(val, 'f', -1, 64))
	// case json.Number:
	// return f.sprintColor(f.NumberColor, string(val))
	case bool:
		return f.sprintFgYellow(strconv.FormatBool(val))
	case nil:
		return f.sprintFgBlack("null")
	case map[string]interface{}:
		return f.processMap(val, depth)
	case []interface{}:
		return f.processArray(val, depth)
	}

	return ""
}

func (f *Formatter) processString(s string) string {
	r := []rune(s)
	if f.StringMaxLength != 0 && len(r) >= f.StringMaxLength {
		s = string(r[0:f.StringMaxLength]) + "..."
	}

	buf := &bytes.Buffer{}
	encoder := sonic.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	encoder.Encode(s)
	s = string(buf.Bytes())
	s = strings.TrimSuffix(s, "\n")

	return f.sprintFgGreen(s)
}

func (f *Formatter) processMap(m map[string]interface{}, depth int) string {
	if len(m) == 0 {
		return "{}"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)
	rows := []string{}
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		val := m[key]
		buf := &bytes.Buffer{}
		encoder := sonic.NewEncoder(buf)
		encoder.SetEscapeHTML(false)
		encoder.Encode(key)
		s := strings.TrimSuffix(string(buf.Bytes()), "\n")
		k := f.sprintFgBlue(s)
		v := f.pretty(val, depth+1)

		valueIndent := " "
		if f.Newline == "" {
			valueIndent = ""
		}
		row := fmt.Sprintf("%s%s:%s%s", nextIndent, k, valueIndent, v)
		rows = append(rows, row)
	}

	return fmt.Sprintf("{%s%s%s%s}", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) processArray(a []interface{}, depth int) string {
	if len(a) == 0 {
		return "[]"
	}

	currentIndent := f.generateIndent(depth - 1)
	nextIndent := f.generateIndent(depth)
	rows := []string{}

	for _, val := range a {
		c := f.pretty(val, depth+1)
		row := nextIndent + c
		rows = append(rows, row)
	}
	return fmt.Sprintf("[%s%s%s%s]", f.Newline, strings.Join(rows, ","+f.Newline), f.Newline, currentIndent)
}

func (f *Formatter) generateIndent(depth int) string {
	return strings.Repeat(" ", f.Indent*depth)
}

func (f *Formatter) sprintFgCyan(s string) string {
	if f.DisabledColor {
		return fmt.Sprint(s)
	}
	return colour.Cyan(s, colour.B)
}

func (f *Formatter) sprintFgYellow(s string) string {
	if f.DisabledColor {
		return fmt.Sprint(s)
	}
	return colour.Yellow(s, colour.B)
}

func (f *Formatter) sprintFgBlack(s string) string {
	if f.DisabledColor {
		return fmt.Sprint(s)
	}
	return colour.Black(s, colour.B)
}

func (f *Formatter) sprintFgGreen(s string) string {
	if f.DisabledColor {
		return fmt.Sprint(s)
	}
	return colour.Green(s, colour.B)
}

func (f *Formatter) sprintFgBlue(s string) string {
	if f.DisabledColor {
		return fmt.Sprint(s)
	}
	return colour.Blue(s, colour.B)
}

// Marshal JSON data with default options.
func Marshal(v interface{}) ([]byte, error) {
	return NewFormatter().Marshal(v)
}

// Format JSON string with default options.
func Format(data []byte) ([]byte, error) {
	return NewFormatter().Format(data)
}
