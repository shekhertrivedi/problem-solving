package input

import (
	"fmt"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/services"
)

// AddSpouseProcessor AddSpouseProcessor instance
type AddSpouseProcessor struct {
	services.AddSpouseService
}

var addSpouseProcessor *AddSpouseProcessor
var once3 sync.Once

// GetAddSpouseProcessorInstance GetAddSpouseProcessorInstance
func GetAddSpouseProcessorInstance() *AddSpouseProcessor {
	once3.Do(func() {
		addSpouseProcessor = &AddSpouseProcessor{services.GetAddSpouseServiceInstance()}
	})
	return addSpouseProcessor
}

// ProcessCommand process command
func (acp *AddSpouseProcessor) ProcessCommand(command string) {

	// call to add spouse service
	fmt.Println(command)
	err := acp.AddSpouseService.AddSpouse(command)
	if err != nil {
		fmt.Println("Error occured while adding spouse. Error: ", err)
		return
	}

	fmt.Println("SPOUSE_ADDITION_SUCCEEDED")

}
