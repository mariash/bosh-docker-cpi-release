package container

import (
	"github.com/fsouza/go-dockerclient"
)

type Finder struct {
	client *docker.Client
}

func NewFinder(client *docker.Client) *Finder {
	return &Finder{
		client: client,
	}
}

func (f *Finder) Find(containerID string) (*docker.Container, bool, error) {
	container, err := f.client.InspectContainer(containerID)
	if err != nil {
		return &docker.Container{}, false, err
	}

	if container.ID == "" {
		return &docker.Container{}, false, nil
	}

	return container, true, nil
}
