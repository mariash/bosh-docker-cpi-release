package container

import (
	"github.com/fsouza/go-dockerclient"
)

type Creator struct {
	client *docker.Client
}

func NewCreator(client *docker.Client) *Creator {
	return &Creator{
		client: client,
	}
}

func (c *Creator) Create(imageName string) (string, error) {
	hostConfig := &docker.HostConfig{
		Privileged:      true,
		PublishAllPorts: true,
		PortBindings: map[docker.Port][]docker.PortBinding{
			docker.Port("6868"): {{HostIP: "0.0.0.0", HostPort: "6868"}},
		},
	}
	containerOptions := docker.CreateContainerOptions{
		Config: &docker.Config{
			ExposedPorts: map[docker.Port]struct{}{
				docker.Port("6868"): {},
			},
			Image: imageName,
			Cmd:   []string{"/usr/sbin/runsvdir-start"},
		},
		HostConfig: hostConfig,
	}

	container, err := c.client.CreateContainer(containerOptions)
	if err != nil {
		return "", nil
	}

	err = c.client.StartContainer(container.ID, hostConfig)
	if err != nil {
		return "", nil
	}

	return container.ID, nil
}
