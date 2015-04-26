package container

import (
	"bytes"
	"encoding/json"

	bwcvm "github.com/cppforlife/bosh-warden-cpi/vm"
	"github.com/fsouza/go-dockerclient"

	cfg "../../config"
)

const settingsFilePathInContainer = "/var/vcap/bosh/registry.json"
const userDataFilePathInContainer = "/var/vcap/bosh/warden-cpi-user-data.json"

type SettingsUpdaterFactory struct {
	client *docker.Client
	config cfg.Config
}

func NewSettingsUpdaterFactory(client *docker.Client, config cfg.Config) *SettingsUpdaterFactory {
	return &SettingsUpdaterFactory{
		client: client,
		config: config,
	}
}

func (f *SettingsUpdaterFactory) Create(containerID string) *settingsUpdater {
	return &settingsUpdater{
		containerID: containerID,
		client:      f.client,
		config:      f.config,
	}
}

type settingsUpdater struct {
	containerID string
	client      *docker.Client
	config      cfg.Config
}

func (s *settingsUpdater) Setup() error {
	userDataContents := bwcvm.UserDataContentsType{
		Registry: bwcvm.RegistryType{
			Endpoint: settingsFilePathInContainer,
		},
	}

	contents, err := json.Marshal(userDataContents)
	if err != nil {
		return err
	}

	return s.updateFile(userDataFilePathInContainer, contents)
}

func (s *settingsUpdater) Update(agentID string, env bwcvm.Environment) error {
	agentOptions := bwcvm.AgentOptions{
		Mbus:      s.config.Mbus,
		Blobstore: s.config.Blobstore,
	}

	agentEnv := bwcvm.NewAgentEnvForVM(agentID, s.containerID, bwcvm.Networks{}, env, agentOptions)

	contents, err := json.Marshal(agentEnv)
	if err != nil {
		return err
	}

	return s.updateFile(settingsFilePathInContainer, contents)
}

func (s *settingsUpdater) updateFile(pathInContainer string, contents []byte) error {
	exec, err := s.client.CreateExec(docker.CreateExecOptions{
		Container:   s.containerID,
		AttachStdin: true,
		Cmd:         []string{"sh", "-c", "cat $1 > " + pathInContainer},
	})
	if err != nil {
		return err
	}

	return s.client.StartExec(exec.ID, docker.StartExecOptions{
		InputStream: bytes.NewReader(contents),
	})
}
