package input

import (
	"fmt"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/services"
)

// RelationshipProcessor RelationshipProcessor
type RelationshipProcessor struct {
	services.FindRelationshipService
}

var relationshipProcessor *RelationshipProcessor
var once2 sync.Once

// GetRelationshipProcessorInstance get Relationship processor instance
func GetRelationshipProcessorInstance() *RelationshipProcessor {
	once2.Do(func() {
		relationshipProcessor = &RelationshipProcessor{services.GetFindRelationshipServiceInstance()}
	})
	return relationshipProcessor
}

// ProcessCommand process command
func (acp *RelationshipProcessor) ProcessCommand(command string) {

	// call to relationship handler service
	fmt.Println(command)
	names, err := acp.FindRelationshipService.FindRelationship(command)
	if err != nil {
		fmt.Println("Error occured while getting relationship. Error: ", err)
		return
	}
	for _, v := range names {
		fmt.Print(v, " ")
	}
	fmt.Println()
	fmt.Println("**************************************")
}
