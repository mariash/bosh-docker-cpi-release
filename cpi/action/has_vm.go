package action

import (
	"./container"
)

type hasVM struct {
	containerFinder *container.Finder
}

func NewHasVM(containerFinder *container.Finder) hasVM {
	return hasVM{containerFinder: containerFinder}
}

func (a hasVM) Run(vmCID string) (bool, error) {
	_, found, err := a.containerFinder.Find(vmCID)
	if err != nil {
		return false, err
	}

	return found, nil
}
