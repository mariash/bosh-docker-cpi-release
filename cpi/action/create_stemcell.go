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

func (a createStemcell) Run(stemcellPath string, cloudProperties map[string]interface{}) (string, error) {
	name := parseStemcellName(cloudProperties)
	imageOptions := docker.ImportImageOptions{
		Repository: name,
		Source:     stemcellPath,
	}

	err := a.client.ImportImage(imageOptions)
	if err != nil {
		return "", err
	}

	return name, nil
}

func parseStemcellName(cloudProperties map[string]interface{}) string {
	name := "bosh-stemcell"

	if foundName, ok := cloudProperties["name"]; ok {
		if nameStr, ok := foundName.(string); ok {
			name = nameStr
		}
	}

	version := "latest"

	if foundVersion, ok := cloudProperties["version"]; ok {
		if versionStr, ok := foundVersion.(string); ok {
			version = versionStr
		}
	}

	return name + ":" + version
}
