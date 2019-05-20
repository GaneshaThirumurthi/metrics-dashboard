package clients

import (
	vsts "github.com/samkreter/vsts-goclient/client"
	"log"
)

// GetRepositoryClient retrives the respository client
func GetRepositoryClient(config *vsts.Config) (*vsts.Client, error) {

	vstsClient, err := vsts.NewClient(config)
	if err != nil {
		log.Fatal(err)
	}
	return vstsClient, nil
}
