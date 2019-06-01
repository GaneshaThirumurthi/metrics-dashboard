package main

import (
	"github.com/GaneshaThirumurthi/metrics-dashboard/clients"
	"github.com/GaneshaThirumurthi/metrics-dashboard/config"
	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
	"github.com/GaneshaThirumurthi/metrics-dashboard/store"
	vsts "github.com/samkreter/vsts-goclient/client"
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
	vstsClient.GetBranch("master")
	db := store.Database{}
	defer db.Instance.Close()
	db.StartServer()
}
