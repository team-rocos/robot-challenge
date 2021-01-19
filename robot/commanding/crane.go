package commanding

import (
	"RobotWarehouseSimLib/librobot"
	"errors"
	"time"
)

// craneHasCommand is used by ParseCommand to check if the command is implemented by crane.
func craneHasCommand(command string) bool {
	commands := []string{"G", "D"}
	for _, c := range commands {
		if c == command {
			return true
		}
	}
	return false
}

// crane provides an abstraction for grab and drop commands.
type crane struct {
	command     string
	isStarted   bool
	isCompleted bool

	startTime time.Time
}

// Start begins the go routine corresponding to the command.
func (p *crane) Start() error {
	if !p.isStarted {
		p.isStarted = true
		switch p.command {
		case "G":
			go p.grab()
		case "D":
			go p.drop()
		}
		return nil
	}
	return errors.New("command was already started")
}

func (p crane) HasStarted() bool {
	return p.isStarted
}

func (p crane) HasCompleted() bool {
	// Conditions for command completed
	return p.isStarted && p.isCompleted
}

// FinalState returns the RobotState resulting from this command.
func (p crane) FinalState(state librobot.RobotState) librobot.RobotState {
	hasCrate := state.HasCrate
	switch p.command {
	case "G":
		hasCrate = true
		break
	case "D":
		hasCrate = false
		break
	}
	return librobot.RobotState{X: state.X, Y: state.Y, HasCrate: hasCrate}
}

// GetCommand returns the command string, used for debugging.
func (p crane) GetCommand() string {
	return p.command
}

// grab loops until an end condition is reached.
func (p *crane) grab() {
	// loop while command unfinished
	p.startTime = time.Now()
	for true {
		// Assuming millisecond resolution is adequate
		if time.Now().Sub(p.startTime).Milliseconds() > 1000 {
			break
		}
	}
	p.isCompleted = true
}

// drop loops until an end condition is reached.
func (p *crane) drop() {
	// loop while command unfinished
	p.startTime = time.Now()
	for true {
		// Assuming millisecond resolution is adequate
		if time.Now().Sub(p.startTime).Milliseconds() > 1000 {
			break
		}
	}
	p.isCompleted = true
}

// ALL DONE.
