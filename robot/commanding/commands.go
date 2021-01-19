package commanding

import (
	"RobotWarehouseSimLib/librobot"
)

// Command provides an abstraction of a command that can be executed by a Robot
type Command interface {
	Start() error
	HasStarted() bool
	HasCompleted() bool
	FinalState(librobot.RobotState) librobot.RobotState
	GetCommand() string
}

// ParseCommand takes the command string and returns the correct struct implementing Command (if one exists)
func ParseCommand(command string) Command {
	if craneHasCommand(command) {
		return &crane{command: command}
	}
	if moveHasCommand(command) {
		return &move{command: command}
	}
	// Add more commands here.
	return nil
}

// ALL DONE.
