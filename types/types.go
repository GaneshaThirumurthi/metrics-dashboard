package types

import (
	"time"
)

// type Repository struct {
// 	id string  `json:"id,omitempty"`
// 	name string `json:"name,omitempty"`
// 	url string `url:"name,omitempty"`
// }

type PullRequestList struct {
	PullRequests []*PullRequest
}

type PullRequest struct {
	PullRequestID int
	Status        string
	CreationDate  time.Time `json:"creationDate,omitempty"`
	ClosedDate    time.Time `json:"closedDate,omitempty"`
	Title         string    `json:"title,omitempty"`
	Description   string    `json:"description,omitempty"`
	MergeStatus   string    `json:"mergeStatus,omitempty"`
}

// PullRequestSearchCriteria for PRs
type PullRequestSearchCriteria struct {
	Status string `json:"status,omitempty"`
}
