package action

import (
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

func (a createVM) Run(agentID string, stemcellCID string, _ interface{}, networks bwcaction.Networks, _ []bwcaction.DiskCID, env bwcaction.Environment) (string, error) {
	containerID, err := a.containerCreator.Create(stemcellCID)
	if err != nil {
		return "", nil
	}

	settingsUpdater := a.settingsUpdaterFactory.Create(containerID)

	err = settingsUpdater.Setup()
	if err != nil {
		return "", nil
	}

	err = settingsUpdater.Update(agentID, bwcvm.Environment(env))
	if err != nil {
		return "", nil
	}

	return containerID, nil
}
