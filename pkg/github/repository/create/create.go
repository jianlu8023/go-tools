package create

import (
	"github.com/go-resty/resty/v2"
)

var (
	clt = resty.New()
)

type CreateRepositoryQuery struct {
	// repo name
	// The name of the repository.
	Name string `json:"name"`
	// repo description
	// A short description of the repository.
	Description string `json:"description"`
	// repo homepage
	// A URL with more information about the repository.
	Homepage string `json:"homepage"`
	// repo private
	// Whether the repository is private.
	// true this repo is private
	// false this repo is public
	// default false
	Private bool `json:"private"`
	// repo visibility
	// The visibility of the repository.
	// public or private
	Visibility string `json:"visibility"`
	// repo has issues
	// Either true to enable issues for this repository or false to disable them.
	// default true
	HasIssues bool `json:"has_issues"`

	// Either true to enable projects for this repository or
	// false to disable them. Note: If you're creating a repository
	// in an organization that has disabled repository projects,
	// the default is false, and if you pass true, the API returns an error.
	// default true
	HasProjects bool `json:"has_projects"`

	// Either true to enable the wiki for this repository or false to disable it.
	// default true
	HasWiki bool `json:"has_wiki"`

	// Whether downloads are enabled.
	// default true
	HasDownloads bool `json:"has_downloads"`

	// Either true to make this repo
	// available as a template repository or false to prevent it.
	// default false
	IsTemplate bool `json:"is_template"`

	// The id of the team that will be granted access to this repository.
	// This is only valid when creating a repository in an organization.
	TeamId int `json:"team_id"`

	// Pass true to create an initial commit with empty README.
	// default false
	AutoInit bool `json:"auto_init"`

	// Desired language or platform .gitignore template to apply.
	// Use the name of the template without the extension. For example, "Haskell".
	GitignoreTemplate string `json:"gitignore_template"`

	// Choose an open source license template that best suits your needs,
	// and then use the license keyword as the license_template string.
	// For example, "mit" or "mpl-2.0".
	LicenseTemplate string `json:"license_template"`

	// Either true to allow squash-merging pull requests,
	// or false to prevent squash-merging.
	// default true
	AllowSquashMerge bool `json:"allow_squash_merge"`

	// Either true to allow merging pull requests with a merge commit,
	// or false to prevent merging pull requests with merge commits.
	//
	// 默认: true
	AllowMergeCommit bool `json:"allow_merge_commit"`

	// Either true to allow rebase-merging pull requests,
	// or false to prevent rebase-merging.
	//
	// 默认: true
	AllowRebaseMerge bool `json:"allow_rebase_merge"`

	// Either true to allow auto-merge on pull requests,
	// or false to disallow auto-merge.
	//
	// 默认: false
	AllowAutoMerge bool `json:"allow_auto_merge"`

	// Either true to allow automatically deleting head branches
	// when pull requests are merged, or false to prevent automatic deletion.
	// The authenticated user must be an organization owner to set
	// this property to true.
	//
	// 默认: false
	DeleteBranchOnMerge bool `json:"delete_branch_on_merge"`

	// Either true to allow squash-merge commits to use pull request title,
	// or false to use commit message. **This property has been deprecated.
	// Please use squash_merge_commit_title instead.
	//
	// 默认: false
	UseSquashPrTitleAsDefault bool `json:"use_squash_pr_title_as_default"`

	// The default value for a squash merge commit title:
	//
	// PR_TITLE - default to the pull request's title.
	// COMMIT_OR_PR_TITLE - default to the commit's title (if only one commit)
	// or the pull request's title (when more than one commit).
	//
	// 可以是以下选项之一: PR_TITLE, COMMIT_OR_PR_TITLE
	SquashMergeCommitTitle string `json:"squash_merge_commit_title"`

	// The default value for a squash merge commit message:
	//
	// PR_BODY - default to the pull request's body.
	// COMMIT_MESSAGES - default to the branch's commit messages.
	// BLANK - default to a blank commit message.
	//
	// 可以是以下选项之一: PR_BODY, COMMIT_MESSAGES, BLANK
	SquashMergeCommitMessage string `json:"squash_merge_commit_message"`

	// The default value for a merge commit title.
	//
	// PR_TITLE - default to the pull request's title.
	// MERGE_MESSAGE - default to the classic title for a merge message
	// (e.g., Merge pull request #123 from branch-name).
	//
	// 可以是以下选项之一: PR_TITLE, MERGE_MESSAGE
	MergeCommitTitle string `json:"merge_commit_title"`

	// The default value for a merge commit message.
	//
	// PR_TITLE - default to the pull request's title.
	// PR_BODY - default to the pull request's body.
	// BLANK - default to a blank commit message.
	//
	// 可以是以下选项之一: PR_BODY, PR_TITLE, BLANK
	MergeCommitMessage string `json:"merge_commit_message"`
}
