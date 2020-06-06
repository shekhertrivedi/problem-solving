package services

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/model"
	"github.com/ContinuumLLC/familytreeQues/src/repository"
)

// FindRelationshipService FindRelationshipService instance
type FindRelationshipService interface {
	FindRelationship(command string) ([]string, error)
}

// FindRelationshipServiceImpl FindRelationshipService implementation
type FindRelationshipServiceImpl struct {
	repository.RelationsRepository
}

var findRelationshipService FindRelationshipService
var once2 sync.Once

// GetFindRelationshipServiceInstance get relationship service instance
func GetFindRelationshipServiceInstance() FindRelationshipService {
	once2.Do(func() {
		findRelationshipService = &FindRelationshipServiceImpl{repository.GetRepositoryInstance()}
	})
	return findRelationshipService
}

// FindRelationship FindRelationship
func (frsi *FindRelationshipServiceImpl) FindRelationship(command string) ([]string, error) {

	findRelationRequest, err := sanitizeFindRelationInput(command)
	if err != nil {
		return []string{}, fmt.Errorf("Failed to find relationship. %v", err)
	}

	var personDetails *model.Person
	if personDetails = frsi.RelationsRepository.GetMemberDetailsByName(findRelationRequest.Name); personDetails == nil {
		return []string{}, fmt.Errorf("Name given is not present. ChildName: %v", findRelationRequest.Name)
	}

	switch strings.ToUpper(findRelationRequest.RelationName) {
	case "PATERNAL-UNCLE":
		return frsi.paternalUncle(personDetails)
	case "MATERNAL-UNCLE":
		return frsi.maternalUncle(personDetails)
	case "PATERNAL-AUNT":
		return frsi.paternalAunt(personDetails)
	case "MATERNAL-AUNT":
		return frsi.maternalAunt(personDetails)
	case "SISTER-IN-LAW":
		return frsi.sisterInLaw(personDetails)
	case "BROTHER-IN-LAW":
		return frsi.brotherInLaw(personDetails)
	case "SON":
		return frsi.son(personDetails)
	case "DAUGHTER":
		return frsi.daughter(personDetails)
	case "SIBLINGS":
		return frsi.siblings(personDetails)
	default:
		return []string{}, fmt.Errorf("Relationship name given is invalid. %v", findRelationRequest.RelationName)

	}
}

func sanitizeFindRelationInput(command string) (model.FindRelationRequest, error) {

	commands := strings.Split(command, " ")
	name := strings.TrimSpace(commands[1])
	relationName := strings.TrimSpace(commands[2])

	findRelationRequest := model.FindRelationRequest{Name: name, RelationName: relationName}
	if len(name) == 0 || len(relationName) == 0 {
		return findRelationRequest, fmt.Errorf("Mandatory input values missing. Name: %v RelationName: %v", name, relationName)
	}
	return findRelationRequest, nil
}

func (frsi *FindRelationshipServiceImpl) paternalUncle(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil &&
		personDetails.Relation.Father.Relation.Mother != nil {

		fathername := personDetails.Relation.Father.Name
		uncles := personDetails.Relation.Father.Relation.Mother.Relation.Children

		for _, v := range uncles {
			if v.Name == fathername {
				continue
			}
			if strings.ToUpper(v.Gender) == "MALE" {
				names = append(names, v.Name)
			}
		}
	}

	return names, nil
}

func (frsi *FindRelationshipServiceImpl) paternalAunt(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil &&
		personDetails.Relation.Father.Relation.Mother != nil {

		aunts := personDetails.Relation.Father.Relation.Mother.Relation.Children

		for _, v := range aunts {
			if strings.ToUpper(v.Gender) == "FEMALE" {
				names = append(names, v.Name)
			}
		}
	}

	return names, nil
}

func (frsi *FindRelationshipServiceImpl) maternalUncle(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil &&
		personDetails.Relation.Mother.Relation.Mother != nil {

		uncles := personDetails.Relation.Mother.Relation.Mother.Relation.Children
		for _, v := range uncles {
			if strings.ToUpper(v.Gender) == "MALE" {
				names = append(names, v.Name)
			}
		}
	}

	return names, nil
}

func (frsi *FindRelationshipServiceImpl) maternalAunt(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil &&
		personDetails.Relation.Mother.Relation.Mother != nil {

		motherName := personDetails.Relation.Mother.Name
		aunts := personDetails.Relation.Mother.Relation.Mother.Relation.Children
		for _, v := range aunts {
			if v.Name == motherName {
				continue
			}
			if strings.ToUpper(v.Gender) == "FEMALE" {
				names = append(names, v.Name)
			}
		}

	}

	return names, nil
}

func (frsi *FindRelationshipServiceImpl) sisterInLaw(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)

	if spouse := personDetails.Relation.Spouse; spouse != nil && spouse.Relation.Mother != nil {
		spouseSiblings := spouse.Relation.Mother.Relation.Children
		for _, v := range spouseSiblings {
			if v.Name == spouse.Name {
				continue
			}
			if strings.ToUpper(v.Gender) == "FEMALE" {
				names = append(names, v.Name)
			}
		}
	}
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil {

		personSiblings := personDetails.Relation.Mother.Relation.Children
		for _, v := range personSiblings {
			if v.Name == personDetails.Name {
				continue
			}
			// Assumption MALE can be married to FEMALE only and vice versa
			// Assumption Every person will have mother and father
			if strings.ToUpper(v.Gender) == "MALE" && v.Relation.Spouse != nil {
				names = append(names, v.Relation.Spouse.Name)
			}
		}
	}
	return names, nil
}

func (frsi *FindRelationshipServiceImpl) brotherInLaw(personDetails *model.Person) ([]string, error) {

	names := make([]string, 0)

	if spouse := personDetails.Relation.Spouse; spouse != nil && spouse.Relation.Mother != nil {
		spouseSiblings := spouse.Relation.Mother.Relation.Children
		for _, v := range spouseSiblings {
			if v.Name == spouse.Name {
				continue
			}
			if strings.ToUpper(v.Gender) == "MALE" {
				names = append(names, v.Name)
			}
		}
	}
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil {
		personSiblings := personDetails.Relation.Mother.Relation.Children
		for _, v := range personSiblings {
			if v.Name == personDetails.Name {
				continue
			}
			if strings.ToUpper(v.Gender) == "FEMALE" && v.Relation.Spouse != nil {
				names = append(names, v.Relation.Spouse.Name)
			}
		}
	}
	return names, nil
}

func (frsi *FindRelationshipServiceImpl) son(personDetails *model.Person) ([]string, error) {

	// Assumption MALE too can avail this service
	// Assumption married person can only have son :D

	names := make([]string, 0)
	var children []*model.Person
	if strings.ToUpper(personDetails.Gender) == "FEMALE" {
		children = personDetails.Relation.Children
	} else if spouse := personDetails.Relation.Spouse; spouse != nil {
		children = spouse.Relation.Children
	}

	for _, v := range children {
		if strings.ToUpper(v.Gender) == "MALE" {
			names = append(names, v.Name)
		}
	}
	return names, nil
}

func (frsi *FindRelationshipServiceImpl) daughter(personDetails *model.Person) ([]string, error) {
	// Assumption MALE too can avail this service
	// Assumption married person can only have daughter :D

	names := make([]string, 0)
	var children []*model.Person
	if strings.ToUpper(personDetails.Gender) == "FEMALE" {
		children = personDetails.Relation.Children
	} else if spouse := personDetails.Relation.Spouse; spouse != nil {
		children = spouse.Relation.Children
	}

	for _, v := range children {
		if strings.ToUpper(v.Gender) == "FEMALE" {
			names = append(names, v.Name)
		}
	}
	return names, nil
}

func (frsi *FindRelationshipServiceImpl) siblings(personDetails *model.Person) ([]string, error) {
	names := make([]string, 0)
	if personDetails.Relation.Mother != nil &&
		personDetails.Relation.Father != nil {

		for _, v := range personDetails.Relation.Mother.Relation.Children {
			if v.Name == personDetails.Name {
				continue
			}
			names = append(names, v.Name)
		}

	}
	return names, nil
}
