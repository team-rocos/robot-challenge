package commanding

import (
	"RobotWarehouseSimLib/librobot"
	"errors"
	"time"
)

// moveHasCommand is used by ParseCommand to check if the command is implemented by crane.
func moveHasCommand(command string) bool {
	commands := []string{"N", "E", "S", "W", "NE", "NW", "SE", "SW", "EN", "WN", "ES", "WS"}
	for _, c := range commands {
		if c == command {
			return true
		}
	}
	return false
}

// move provides an abstraction for movement commands.
type move struct {
	command     string
	isStarted   bool
	isCompleted bool

	startTime time.Time
}

// Start begins the go routine corresponding to the command.
func (p *move) Start() error {
	if !p.isStarted {
		p.isStarted = true
		go p.movement()
		return nil
	}
	return errors.New("command was already started")
}

func (p move) HasStarted() bool {
	return p.isStarted
}

func (p move) HasCompleted() bool {
	// Conditions for command completed
	return p.isStarted && p.isCompleted
}

// FinalState returns the RobotState resulting from this command.
func (p move) FinalState(state librobot.RobotState) librobot.RobotState {
	x := state.X
	y := state.Y

	switch p.command {
	case "N":
		y += 1
		break
	case "E":
		x += 1
		break
	case "S":
		if y > 0 {
			y -= 1
		}
		break
	case "W":
		if x > 0 {
			x -= 1
		}
		break
	case "NE":
		x += 1
		y += 1
		break
	case "SE":
		x += 1
		if y > 0 {
			y -= 1
		}
		break
	case "SW":
		if x > 0 {
			x -= 1
		}
		if y > 0 {
			y -= 1
		}
		break
	case "NW":
		if x > 0 {
			x -= 1
		}
		y += 1
		break
	}

	return librobot.RobotState{X: x, Y: y, HasCrate: state.HasCrate}
}

// GetCommand returns the command string, used for debugging.
func (p move) GetCommand() string {
	return p.command
}

// movement loops until an end condition is reached.
func (p *move) movement() {
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
