package action

import (
	"github.com/fsouza/go-dockerclient"
)

type deleteVM struct {
	client *docker.Client
}

func NewDeleteVM(client *docker.Client) deleteVM {
	return deleteVM{client: client}
}

func (a deleteVM) Run(vmCID string) (interface{}, error) {
	err := a.client.KillContainer(docker.KillContainerOptions{
		ID: vmCID,
	})
	return nil, err
}
