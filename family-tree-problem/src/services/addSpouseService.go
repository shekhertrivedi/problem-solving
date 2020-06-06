package services

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/model"
	"github.com/ContinuumLLC/familytreeQues/src/repository"
)

// AddSpouseService AddSpouseService
type AddSpouseService interface {
	AddSpouse(command string) error
}

// AddSpouseServiceImpl add
type AddSpouseServiceImpl struct {
	repository.RelationsRepository
}

var addSpouseService AddSpouseService
var once3 sync.Once

// GetAddSpouseServiceInstance get instance
func GetAddSpouseServiceInstance() AddSpouseService {
	once3.Do(func() {
		addSpouseService = &AddSpouseServiceImpl{repository.GetRepositoryInstance()}
	})
	return addSpouseService
}

// AddSpouse GetRepoState
func (rssi *AddSpouseServiceImpl) AddSpouse(command string) error {
	fmt.Println("Adding spouse. ", command)

	addSpouseReq, err := sanitizeInputForSpouse(command)
	if err != nil {
		return fmt.Errorf("Failed to add spouse %v", err)
	}

	if val := rssi.RelationsRepository.GetMemberDetailsByName(addSpouseReq.Name); val != nil {
		return fmt.Errorf("Already added. Name: %v", addSpouseReq.Name)
	}

	var spouseDetails *model.Person
	if spouseDetails = rssi.RelationsRepository.GetMemberDetailsByName(addSpouseReq.SpouseName); spouseDetails != nil &&
		spouseDetails.Gender != addSpouseReq.Gender {

		personRelations := model.Relationship{
			Spouse: spouseDetails,
		}

		personDetails := model.Person{
			Name:     addSpouseReq.Name,
			Gender:   addSpouseReq.Gender,
			Relation: personRelations,
		}

		rssi.RelationsRepository.PutMemberDetails(addSpouseReq.Name, &personDetails)
		spouseDetails.Relation.Spouse = &personDetails

		return nil
	}
	return fmt.Errorf("Spouse is not present. %v", addSpouseReq.SpouseName)
}

func sanitizeInputForSpouse(command string) (model.AddSpouseRequest, error) {

	commands := strings.Split(command, " ")
	if len(commands) != 4 {
		return model.AddSpouseRequest{}, fmt.Errorf("mandatory parameters not provided %v", commands)
	}

	spousename := strings.TrimSpace(commands[1])
	name := strings.TrimSpace(commands[2])
	gender := strings.TrimSpace(commands[3])

	addSpouseReq := model.AddSpouseRequest{SpouseName: spousename, Name: name, Gender: gender}
	if len(spousename) == 0 || len(name) == 0 || len(gender) == 0 {
		return addSpouseReq, fmt.Errorf("Mandatory input values missing. SpouseName: %v Name: %v Gender: %v", spousename, name, gender)
	}

	if !isgenderValid(gender) {
		return addSpouseReq, fmt.Errorf("Invalid gender value. Gender: %v", gender)
	}
	return addSpouseReq, nil
}
