package input

import (
	"fmt"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/services"
)

// AddChildProcessor AddChildProcessor instance
type AddChildProcessor struct {
	services.AddChildService
}

var addChildProcessor *AddChildProcessor
var once1 sync.Once

// GetAddChildProcessorInstance GetAddChildProcessorInstance
func GetAddChildProcessorInstance() *AddChildProcessor {
	once1.Do(func() {
		addChildProcessor = &AddChildProcessor{services.GetAddChildServiceInstance()}
	})
	return addChildProcessor
}

// ProcessCommand process command
func (acp *AddChildProcessor) ProcessCommand(command string) {

	// call to add child service
	fmt.Println(command)
	err := acp.AddChildService.AddChild(command)
	if err != nil {
		fmt.Println("Error occured while adding child. Error: ", err)
		return
	}

	fmt.Println("CHILD_ADDITION_SUCCEEDED")
	fmt.Println("**************************************")
}
