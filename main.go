package main

import (
	// "fmt"
	"github.com/GaneshaThirumurthi/metrics-dashboard/clients"
	"github.com/GaneshaThirumurthi/metrics-dashboard/config"
	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
	"github.com/GaneshaThirumurthi/metrics-dashboard/store"
	// "github.com/GaneshaThirumurthi/metrics-dashboard/workers"
	vsts "github.com/samkreter/vsts-goclient/client"
	// "time"
)

func main() {
	startJobs()
}

func startJobs() {
	vstsConfig := &vsts.Config{
		Token:          config.PersonalAccessToken,
		Username:       consts.Username,
		APIVersion:     consts.APIVersion,
		RepositoryName: consts.RepositoryName,
		Project:        consts.Project,
		Instance:       consts.Instance,
	}
	vstsClient, err := clients.GetRepositoryClient(vstsConfig)
	if err != nil {
		panic("Unable to get vsts client")
	}
	// workers.GetAllReleases(vstsClient)
	vstsClient.GetBranch("master")
	// prs, err := vstsClient.GetPullRequests(nil)
	// if err != nil {
	// 	panic("Unable to get prs")
	// }
	// fmt.Println(len(prs))
	// startTime := time.Now().Add(-1000 * time.Hour)
	// fmt.Println("Start time", startTime)
	// prCount, err := workers.GetNumberOfMergedPullRequestsByTime(vstsClient, startTime, time.Now())
	// if err != nil {
	// 	fmt.Println("err", err)
	// }
	// fmt.Printf("In the last week there were %d PRs", prCount)
	store.StartServer()
}
