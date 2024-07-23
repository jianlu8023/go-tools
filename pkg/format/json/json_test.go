package json

import (
	"fmt"
	"testing"
)

var s = struct {
	Name string
	Age  int
}{
	Name: "John",
	Age:  30,
}

func TestPrettyJSON(t *testing.T) {

	json, err := PrettyJSON(s)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Log(json)
}

func TestToJSON(t *testing.T) {
	json, err := ToJSON(s)
	if err != nil {
		t.Errorf("Error: %v", err)
	}
	t.Log(json)
}

type Animal struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func TestPrettyJSON2(t *testing.T) {
	animals := []Animal{
		{Name: "兔子", Age: 10},
		{Name: "猫", Age: 11},
	}
	prettyJSON, _ := PrettyJSON(animals)
	fmt.Println(prettyJSON)
}
