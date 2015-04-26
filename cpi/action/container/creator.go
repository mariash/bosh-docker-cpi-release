package container

import (
	"net"
	"net/url"
	"strings"

	"github.com/fsouza/go-dockerclient"

	cfg "../../config"
)

type Creator struct {
	config cfg.Config
	client *docker.Client
}

func NewCreator(client *docker.Client, config cfg.Config) *Creator {
	return &Creator{
		client: client,
		config: config,
	}
}

func (c *Creator) Create(imageName string) (string, error) {
	portBindings := map[docker.Port][]docker.PortBinding{}
	exposedPorts := map[docker.Port]struct{}{}

	if strings.HasPrefix(c.config.Mbus, "http") {
		mbusURL, err := url.Parse(c.config.Mbus)
		if err != nil {
			return "", err
		}

		_, port, err := net.SplitHostPort(mbusURL.Host)
		if err != nil {
			return "", err
		}

		portBindings[docker.Port(port)] = []docker.PortBinding{
			{HostIP: "0.0.0.0", HostPort: port},
		}
		exposedPorts[docker.Port(port)] = struct{}{}
	}

	hostConfig := &docker.HostConfig{
		Privileged:      true,
		PublishAllPorts: true,
		PortBindings:    portBindings,
	}
	containerOptions := docker.CreateContainerOptions{
		Config: &docker.Config{
			ExposedPorts: exposedPorts,
			Image:        imageName,
			Cmd:          []string{"/usr/sbin/runsvdir-start"},
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
