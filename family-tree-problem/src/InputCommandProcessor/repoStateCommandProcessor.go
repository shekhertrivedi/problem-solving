package input

import (
	"fmt"
	"sync"

	"github.com/ContinuumLLC/familytreeQues/src/services"
)

// RepoStateProcessor RepoStateProcessor instance
type RepoStateProcessor struct {
	services.RepoStateService
}

var repoStateProcessor *RepoStateProcessor
var once4 sync.Once

// GetRepoStateInstance GetRepoStateInstance
func GetRepoStateInstance() *RepoStateProcessor {
	once4.Do(func() {
		repoStateProcessor = &RepoStateProcessor{services.GetRepoStateServiceInstance()}
	})
	return repoStateProcessor
}

// ProcessCommand process command
func (acp *RepoStateProcessor) ProcessCommand(command string) {
	fmt.Println(acp.GetRepoState(command))
	fmt.Println("**************************************")
}
