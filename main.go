package main

import (
	"fmt"
	"github.com/GaneshaThirumurthi/metrics-dashboard/clients"
	"github.com/GaneshaThirumurthi/metrics-dashboard/config"
	"github.com/GaneshaThirumurthi/metrics-dashboard/consts"
	//"github.com/GaneshaThirumurthi/metrics-dashboard/store"
	"github.com/GaneshaThirumurthi/metrics-dashboard/workers"
	vsts "github.com/samkreter/vsts-goclient/client"
)

func main() {
	startJobs()
}

func startJobs() {
	vstsConfig := &vsts.Config{
		Token:          config.PersonalAccessToken,
		Username:       config.DatabaseUsername,
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
	// db := store.Database{}
	// defer db.Instance.Close()
	// db.StartServer()

	coverage := workers.New(consts.AksRepoPath, consts.GoExecutable)
	err = coverage.GenerateCoverage("coverage.out")
	if err != nil {
		fmt.Println("Unable to generate coverage", err)
	}
	err = coverage.GenerateFuncCoverage("coverage.out", "coveragefunc.out")
	if err != nil {
		fmt.Println("Unable to generate func coverage", err)
	}
	functionCoverage, err := coverage.ParseCoverageFile("coveragefunc.out")
	if err != nil {
		fmt.Println("Unable to extract coverage", err)
	}

	fmt.Println("Coverage is: ", functionCoverage)
}
