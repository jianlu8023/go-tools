package container

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/pkg/errors"
)

// GetContainersInspect
// @Description: GetContainersInspect 获取本地docker所有容器信息
// @author ght
// @date 2023-10-25 16:17:31
// @return map[string]map[string]interface{}:
func GetContainersInspect() map[string]*map[string]interface{} {
	containersMap := make(map[string]*map[string]interface{})
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cli, err := client.NewClientWithOpts(client.FromEnv)
	if err != nil {
		if client.IsErrConnectionFailed(err) {
			fmt.Println(fmt.Sprintf("Connect Docker Engine Error: %v", err))
			return nil
		}
		if client.IsErrUnauthorized(err) {
			fmt.Println(fmt.Sprintf("Connect Docker Engine Error: %v", err))
			return nil
		}
		if client.IsErrNotFound(err) {
			fmt.Println(fmt.Sprintf("Connect Docker Engine Error: %v", err))
			return nil
		}
		if client.IsErrNotImplemented(err) {
			fmt.Println(fmt.Sprintf("Connect Docker Engine Error: %v", err))
			return nil
		}
		fmt.Println(fmt.Sprintf("Get Docker Cli Error: %v", err))
		return nil
	}

	containerList, err := cli.ContainerList(ctx, types.ContainerListOptions{All: true})
	if err != nil {
		fmt.Println(fmt.Sprintf("Get Containers List Error: %v", err))
		return nil
	}

	for _, container := range containerList {
		containerInspect, err := cli.ContainerInspect(ctx, container.ID)
		if err != nil {
			fmt.Println(fmt.Sprintf("Get Container %v Inspect Error: %v", container.ID, err))
			return nil
		}

		containersMap[containerInspect.Name[1:]] = &map[string]interface{}{
			"ContainerId": containerInspect.ID,
			"Name":        containerInspect.Name[1:],
			"Status":      containerInspect.State.Status,
			"Image":       containerInspect.Config.Image,
			"Ports":       containerInspect.NetworkSettings.Ports,
			"Command":     containerInspect.Config.Cmd,
			"Env":         containerInspect.Config.Env,
			"Volumes":     containerInspect.Config.Volumes,
			"Mounts":      containerInspect.Mounts,
		}
	}
	return containersMap
}

// GetContainerInspect
// @Description: GetContainerInspect
// @author ght
// @date 2023-10-25 16:40:08
// @param containerName:
// @return map[string]interface{}:
func GetContainerInspect(containerName string) *map[string]interface{} {
	containersMap := GetContainersInspect()

	container := containersMap[containerName]
	if container == nil {
		fmt.Println(fmt.Sprintf("Get Container %v Inspect Error: %v", containerName, errors.Errorf("Current Env Not Exist Container %v", containerName)))
		return nil
	}
	return container
}
