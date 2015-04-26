package action

import (
	"github.com/fsouza/go-dockerclient"
)

type createStemcell struct {
	client *docker.Client
}

func NewCreateStemcell(client *docker.Client) createStemcell {
	return createStemcell{
		client: client,
	}
}

func (a createStemcell) Run(stemcellPath string, _ interface{}) (string, error) {
	imageOptions := docker.ImportImageOptions{
		Repository: "bosh:stemcell",
		Source:     stemcellPath,
	}

	err := a.client.ImportImage(imageOptions)
	if err != nil {
		return "", err
	}

	return "bosh:stemcell", nil
}
