package query

import (
	"fmt"

	"github.com/jianlu8023/go-tools/v2/pkg/http"
	"github.com/jianlu8023/go-tools/v2/pkg/stringer"
)

var (
	clt = http.NewClient()

	queryAll  = "https://api.github.com/users/{{owner}}/repos"
	queryRepo = "https://api.github.com/repos/{{owner}}/{{repo}}"
)

// GitHubAllRepository
// @param owner:
// @param token:
// @return *resty.Response:
// @return error:
func GitHubAllRepository(owner, token string) ([]byte, int, error) {
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}
	if len(token) > 0 {
		authorization := fmt.Sprintf("Bearer %s", token)
		headers["Authorization"] = authorization
	}
	url := stringer.Replace(queryAll, map[string]string{
		"{{owner}}": owner,
	})
	return clt.SetHeaders(headers).GET(url, map[string]interface{}{})
}

// GitHubRepository
// @param owner:
// @param repo:
// @param token:
// @return *resty.Response:
// @return error:
func GitHubRepository(owner, repo, token string) ([]byte, int, error) {
	headers := map[string]string{
		"Accept":               "application/vnd.github+json",
		"X-GitHub-Api-Version": "2022-11-28",
	}
	if len(token) > 0 {
		authorization := fmt.Sprintf("Bearer %s", token)
		headers["Authorization"] = authorization
	}

	url := stringer.Replace(queryRepo, map[string]string{
		"{{owner}}": owner,
		"{{repo}}":  repo,
	})

	return clt.SetHeaders(headers).GET(url, map[string]interface{}{})

}
