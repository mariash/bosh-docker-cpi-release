package action

import (
	"github.com/fsouza/go-dockerclient"

	bosherr "github.com/cloudfoundry/bosh-agent/errors"
	boshsys "github.com/cloudfoundry/bosh-agent/system"
	bwcaction "github.com/cppforlife/bosh-warden-cpi/action"

	cfg "../config"
	"./container"
)

type factory struct {
	actions map[string]bwcaction.Action
	fs      boshsys.FileSystem
}

func NewFactory(client *docker.Client, config cfg.Config, fs boshsys.FileSystem) bwcaction.Factory {
	containerCreator := container.NewCreator(client)
	settingsUpdaterFactory := container.NewSettingsUpdaterFactory(client, config)

	return &factory{
		actions: map[string]bwcaction.Action{
			"create_stemcell": NewCreateStemcell(client),
			"create_vm":       NewCreateVM(containerCreator, settingsUpdaterFactory),
		},
	}
}

func (f *factory) Create(method string) (bwcaction.Action, error) {
	action, found := f.actions[method]
	if !found {
		return nil, bosherr.New("Could not create action with method %s", method)
	}

	return action, nil
}
