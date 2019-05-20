package workers

import (
	"fmt"
	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
	"github.com/GaneshaThirumurthi/metrics-dashboard/types"
	vsts "github.com/samkreter/vsts-goclient/client"
	"time"
)

// getCompletedPullRequestsByTime returns all PRs given specific filtering params
func getCompletedPullRequestsByTime(repoClient *vsts.Client, startDate time.Time, endDate time.Time) (*types.PullRequestList, error) {
	prs, err := repoClient.GetPullRequests(nil)
	if err != nil {
		return nil, err
	}
	pullRequests := types.PullRequestList{}
	for i, pr := range prs {
		if pr.ClosedDate.After(startDate) && pr.ClosedDate.Before(endDate) {
			pullRequest := &types.PullRequest{}
			pullRequest.Status = pr.Status
			pullRequest.CreationDate = pr.CreationDate
			pullRequest.ClosedDate = pr.ClosedDate
			pullRequests.PullRequests[i] = pullRequest
		}
	}
	return &pullRequests, nil
}

// GetNumberOfMergedPullRequestsByTime returns all successfully merged PRs
func GetNumberOfMergedPullRequestsByTime(repoClient *vsts.Client, startDate time.Time, endDate time.Time) (int, error) {
	prList, err := getCompletedPullRequestsByTime(repoClient, startDate, endDate)
	if err != nil {
		return 0, err
	}
	count := 0
	for _, pr := range prList.PullRequests {
		if pr.Status == consts.PullRequestStatusCompleted {
			count++
		}
	}
	return count, nil
}

// GetPullRequestsMergeDelay given a time range, determines the delay in hours to merge a PR
func GetPullRequestsMergeDelay(repoClient *vsts.Client, startDate time.Time, endDate time.Time) (float64, float64, error) {
	prList, err := getCompletedPullRequestsByTime(repoClient, startDate, endDate)
	if err != nil {
		return 0, 0, err
	}
	delay, longest, count := 0.0, 0.0, 0.0
	for _, pr := range prList.PullRequests {
		if pr.Status == consts.PullRequestStatusCompleted {
			count++
			hours := pr.ClosedDate.Sub(pr.CreationDate).Hours()
			delay += hours
			longest = getMax(longest, hours)
		}
	}
	return delay / count, longest, nil
}

func getMax(a float64, b float64) float64 {
	if a > b {
		return a
	}
	return b
}

func GetAllReleases(repoClient *vsts.Client) {
	err := repoClient.GetReleases()
	fmt.Println(err)
}

// Unit test code coverage (fix the calculation for code coverage)
// Number of PRs merged every week
// Time from PR creation to merge
// How long a build sits in a production region
