package container

import (
	"strconv"

	"github.com/fsouza/go-dockerclient"

	boshlog "github.com/cloudfoundry/bosh-agent/logger"

	cfg "../../config"
)

type Creator struct {
	config cfg.Config
	client *docker.Client
	logger boshlog.Logger
}

func NewCreator(client *docker.Client, config cfg.Config, logger boshlog.Logger) *Creator {
	return &Creator{
		client: client,
		config: config,
		logger: logger,
	}
}

type Properties struct {
	Ports []int
}

func (c *Creator) Create(imageName string, properties Properties) (string, error) {
	portBindings := map[docker.Port][]docker.PortBinding{}
	exposedPorts := map[docker.Port]struct{}{}

	for _, port := range properties.Ports {
		portStr := strconv.Itoa(port)
		c.logger.Debug("creator", "Configuring port %s", portStr)

		portBindings[docker.Port(portStr)] = []docker.PortBinding{
			{HostIP: "0.0.0.0", HostPort: portStr},
		}
		exposedPorts[docker.Port(portStr)] = struct{}{}
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

	// clean up compilation logs (which cause permission errors for vcap)
	// fixed later in bosh, no stemcell though yet
	err = c.cleanupDataDir(container.ID)
	if err != nil {
		return "", nil
	}

	return container.ID, nil
}

func (c *Creator) cleanupDataDir(containerID string) error {
	exec, err := c.client.CreateExec(docker.CreateExecOptions{
		Container: containerID,
		Cmd:       []string{"sh", "-c", "rm -rf /var/vcap/data/sys && mkdir -p /var/vcap/data/sys"},
	})
	if err != nil {
		return err
	}

	return c.client.StartExec(exec.ID, docker.StartExecOptions{})
}
