package types

import (
	"time"
)

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

// DatabaseConfig is the configuration elements to create a database
type DatabaseConfig struct {
	Server   string
	Port     int
	User     string
	Password string
	Database string
}
