package container

import (
	"fmt"
	"testing"

	"github.com/jianlu8023/go-tools/pkg/docker/container"
	"github.com/jianlu8023/go-tools/pkg/format/json"
)

func TestGetContainersInspect(t *testing.T) {
	inspect := container.GetContainersInspect()
	fmt.Println(json.PrettyJSON(inspect))
}

func TestGetContainerInspect(t *testing.T) {
	inspect := container.GetContainerInspect("mysql")
	fmt.Println(json.PrettyJSON(inspect))
}
