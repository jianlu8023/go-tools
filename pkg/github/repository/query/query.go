package query

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/jianlu8023/go-tools/pkg/replace"
)

var (
	clt = resty.New()

	queryAll  = "https://api.github.com/users/{{owner}}/repos"
	queryRepo = "https://api.github.com/repos/{{owner}}/{{repo}}"
)

// GitHubAllRepository
// @Description: GitHubAllRepository
// @author ght
// @date 2023-10-07 18:07:52
// @param owner:
// @param token:
// @return *resty.Response:
// @return error:
func GitHubAllRepository(owner, token string) (*resty.Response, error) {
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}
	if len(token) > 0 {
		authorization := fmt.Sprintf("Bearer %s", token)
		headers["Authorization"] = authorization
	}
	url := replace.Replace(queryAll, map[string]string{
		"{{owner}}": owner,
	})
	response, err := clt.R().SetHeaders(headers).Get(url)

	if err != nil {
		return nil, err
	}
	return response, nil
}

// GitHubRepository
// @Description: GitHubRepository
// @author ght
// @date 2023-10-07 18:10:19
// @param owner:
// @param repo:
// @param token:
// @return *resty.Response:
// @return error:
func GitHubRepository(owner, repo, token string) (*resty.Response, error) {
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}
	if len(token) > 0 {
		authorization := fmt.Sprintf("Bearer %s", token)
		headers["Authorization"] = authorization
	}

	url := replace.Replace(queryRepo, map[string]string{
		"{{owner}}": owner,
		"{{repo}}":  repo,
	})
	response, err := clt.R().SetHeaders(headers).Get(url)
	if err != nil {
		return nil, err
	}
	return response, nil
}
