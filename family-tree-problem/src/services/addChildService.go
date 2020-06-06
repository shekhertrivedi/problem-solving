package services

import (
	"fmt"
	"strings"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/model"
	"github.com/ContinuumLLC/familytreeQues/src/repository"
)

// InitializeRepo InitializeRepo
func InitializeRepo() {
	shan := model.Person{Name: "Shan", Gender: "Male"}
	shanRelation := model.Relationship{}

	anga := model.Person{Name: "Anga", Gender: "Female"}
	angaRelation := model.Relationship{}
	angaRelation.Spouse = &shan
	anga.Relation = angaRelation
	shanRelation.Spouse = &anga
	shan.Relation = shanRelation

	repo := repository.GetRepositoryInstance()
	repo.PutMemberDetails("Anga", &anga)
	repo.PutMemberDetails("Shan", &shan)

}

// AddChildService AddChildService instance
type AddChildService interface {
	AddChild(command string) error
}

// AddChildServiceImpl AddChildService implementation
type AddChildServiceImpl struct {
	repository.RelationsRepository
}

var addChildService AddChildService
var once1 sync.Once

// GetAddChildServiceInstance get child service instance
func GetAddChildServiceInstance() AddChildService {
	once1.Do(func() {
		addChildService = &AddChildServiceImpl{repository.GetRepositoryInstance()}
	})
	return addChildService
}

// AddChild add child to the tree
func (acsi *AddChildServiceImpl) AddChild(command string) error {

	addChildReq, err := sanitizeInput(command)
	if err != nil {
		return fmt.Errorf("Failed to add child %v", err)
	}

	if val := acsi.RelationsRepository.GetMemberDetailsByName(addChildReq.ChildName); val != nil {
		return fmt.Errorf("Child already added. ChildName: %v", addChildReq.ChildName)
	}

	var motherDetails *model.Person
	if motherDetails = acsi.RelationsRepository.GetMemberDetailsByName(addChildReq.MotherName); motherDetails != nil {

		childRelations := model.Relationship{
			Mother: motherDetails,
			Father: motherDetails.Relation.Spouse,
		}

		childDetails := model.Person{
			Name:     addChildReq.ChildName,
			Gender:   addChildReq.Gender,
			Relation: childRelations,
		}

		acsi.RelationsRepository.PutMemberDetails(addChildReq.ChildName, &childDetails)

		motherDetails.Relation.Children = append(motherDetails.Relation.Children, &childDetails)

		return nil
	}
	return fmt.Errorf("Mother is not present. MotherName %v", addChildReq.MotherName)
}

func sanitizeInput(command string) (model.AddChildRequest, error) {

	commands := strings.Split(command, " ")
	if len(commands) != 4 {
		return model.AddChildRequest{}, fmt.Errorf("mandatory parameters not provided %v", commands)
	}

	motherName := strings.TrimSpace(commands[1])
	childName := strings.TrimSpace(commands[2])
	gender := strings.TrimSpace(commands[3])

	addChildReq := model.AddChildRequest{MotherName: motherName, ChildName: childName, Gender: gender}
	if len(motherName) == 0 || len(childName) == 0 || len(gender) == 0 {
		return addChildReq, fmt.Errorf("Mandatory input values missing. MotherName: %v ChildName: %v Gender: %v", motherName, childName, gender)
	}

	if !isgenderValid(gender) {
		return addChildReq, fmt.Errorf("Invalid gender value. Gender: %v", gender)
	}
	return addChildReq, nil
}

func isgenderValid(gender string) bool {
	if strings.ToUpper(gender) == "MALE" || strings.ToUpper(gender) == "FEMALE" {
		return true
	}

	return false
}
