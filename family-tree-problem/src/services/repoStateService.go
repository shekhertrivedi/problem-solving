package services

import (
	"fmt"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/model"
	"github.com/ContinuumLLC/familytreeQues/src/repository"
)

// RepoStateService add
type RepoStateService interface {
	GetRepoState(command string) map[string]*model.Person
}

// RepoStateServiceImpl add
type RepoStateServiceImpl struct {
	repository.RelationsRepository
}

var repoStateService RepoStateService
var once4 sync.Once

// GetRepoStateServiceInstance get instance
func GetRepoStateServiceInstance() RepoStateService {
	once4.Do(func() {
		repoStateService = &RepoStateServiceImpl{repository.GetRepositoryInstance()}
	})
	return repoStateService
}

// GetRepoState GetRepoState
func (rssi *RepoStateServiceImpl) GetRepoState(command string) map[string]*model.Person {
	fmt.Println("Get repo state")
	return rssi.RelationsRepository.GetRepoState()
}
