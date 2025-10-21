package api

import (
	"fmt"
)

func GetReleaseAssets(owner, repo, tag string) ([]interface{}, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%v/%v/releases/tags/%v", owner, repo, tag)

	_ = url
	return nil, nil
}
