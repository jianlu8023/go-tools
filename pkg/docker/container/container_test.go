package container

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/format/json"
)

func TestGetContainersInspect(t *testing.T) {
	inspect := GetContainersInspect()
	fmt.Println(json.PrettyJSON(inspect))
}

func TestGetContainerInspect(t *testing.T) {

	inspect := GetContainerInspect("mysql")
	fmt.Println(json.PrettyJSON(inspect))
}
