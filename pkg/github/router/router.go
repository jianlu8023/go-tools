package router

var (
	Router = map[string]map[string]METHOD{
		"ListOrganizationRepositories":   {"https://api.github.com/orgs/{{org}}/repos": GET},
		"CreateAnOrganizationRepository": {"https://api.github.com/orgs/{{org}}/repos": POST},
		"GetARepository":                 {"https://api.github.com/repos/{{owner}}/{{repo}}": GET},
		"UpdateARepository":              {" https://api.github.com/repos/{{owner}}/{{repo}}": PATCH},
		"DeleteARepository":              {"https://api.github.com/repos/{{owner}}/{{repo}}": DELETE},
		"ListRepositoryActivities":       {"https://api.github.com/repos/{{owner}}/{{repo}}/activity": GET},

		//
		"repository_query_all": {"https://api.github.com/users/{{owner}}/repos": GET},
		"repository_query_one": {"https://api.github.com/repos/{{owner}}/{{repo}}": POST},
	}
)

type METHOD int

const (
	GET METHOD = iota
	POST
	DELETE
	PUT
	PATCH
)
