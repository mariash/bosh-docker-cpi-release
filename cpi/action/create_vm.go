package action

import (
	"errors"

	bwcaction "github.com/cppforlife/bosh-warden-cpi/action"
	bwcvm "github.com/cppforlife/bosh-warden-cpi/vm"

	"./container"
)

type createVM struct {
	containerCreator       *container.Creator
	settingsUpdaterFactory *container.SettingsUpdaterFactory
}

func NewCreateVM(
	containerCreator *container.Creator,
	settingsUpdaterFactory *container.SettingsUpdaterFactory,
) createVM {
	return createVM{
		containerCreator:       containerCreator,
		settingsUpdaterFactory: settingsUpdaterFactory,
	}
}

func (a createVM) Run(agentID string, stemcellCID string, cloudProperties container.Properties, networks bwcaction.Networks, _ []bwcaction.DiskCID, env bwcaction.Environment) (string, error) {
	containerID, err := a.containerCreator.Create(stemcellCID, cloudProperties)
	if err != nil {
		return "", nil
	}
	if containerID == "" {
		// todo: get stderr from container
		return "", errors.New("failed to create container for some reason, probably ports collision")
	}

	settingsUpdater := a.settingsUpdaterFactory.Create(containerID)

	err = settingsUpdater.Setup()
	if err != nil {
		return "", nil
	}

	err = settingsUpdater.Update(agentID, networks.AsVMNetworks(), bwcvm.Environment(env))
	if err != nil {
		return "", nil
	}

	return containerID, nil
}
