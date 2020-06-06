package input

import (
	"errors"
	"strings"
)

// GetCommandProcessor GetCommandProcessor by command
func GetCommandProcessor(command string) (Processor, error) {

	switch strings.Split(command, " ")[0] {
	case "ADD_CHILD":
		return GetAddChildProcessorInstance(), nil
	case "ADD_SPOUSE":
		return GetAddSpouseProcessorInstance(), nil
	case "GET_RELATIONSHIP":
		return GetRelationshipProcessorInstance(), nil
	case "GET_REPO_STATE":
		return GetRepoStateInstance(), nil
	default:
		return nil, errors.New("command not supported")
	}
}
