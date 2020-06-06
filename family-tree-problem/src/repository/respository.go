package repository

import (
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/model"
)

// RelationsRepository RelationsRepository
type RelationsRepository interface {
	GetMemberDetailsByName(name string) *model.Person
	PutMemberDetails(name string, person *model.Person) *model.Person
	GetRepoState() map[string]*model.Person
}

// RelationsRepositoryImpl RelationsRepositoryImpl
type RelationsRepositoryImpl struct {
	Repo map[string]*model.Person
}

var relationsRepository RelationsRepository
var once sync.Once

// GetRepositoryInstance get repository instance
func GetRepositoryInstance() RelationsRepository {
	once.Do(func() {
		relationsRepository = &RelationsRepositoryImpl{make(map[string]*model.Person)}
	})
	return relationsRepository
}

// GetMemberDetailsByName GetmemberDetailsByName
func (rri *RelationsRepositoryImpl) GetMemberDetailsByName(name string) *model.Person {

	if val, ok := rri.Repo[name]; ok {
		return val
	}
	return nil
}

// PutMemberDetails PutMemberDetails
func (rri *RelationsRepositoryImpl) PutMemberDetails(name string, person *model.Person) *model.Person {
	if val, ok := rri.Repo[name]; ok {
		rri.Repo[name] = person
		return val
	}
	rri.Repo[name] = person
	return nil
}

// GetRepoState GetRepoState
func (rri *RelationsRepositoryImpl) GetRepoState() map[string]*model.Person {
	return rri.Repo
}
