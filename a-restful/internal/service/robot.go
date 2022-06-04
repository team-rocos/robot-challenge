package service

import (
	"fmt"
	"strings"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
)

type IRobotService interface {
	EnqueueTask(command string) error
}

type RobotService struct {
	log   *util.Logger
	robot thirdparty.Robot
}

func NewRobotService(robot thirdparty.Robot) *RobotService {
	return &RobotService{
		log:   util.NewLogger("RobotService"),
		robot: robot,
	}
}

func (s *RobotService) EnqueueTask(command string) error {
	s.log.Info("run command")
	return nil
}

func (s *RobotService) validateMovement(command string) error {
	position := s.robot.CurrentState()
	moves := strings.Split(command, " ")

	x := position.X
	y := position.Y
	for _, move := range moves {
		switch move {
		case North:
			y = y + 1
		case East:
			x = x + 1
		case South:
			y = y - 1
		case West:
			x = x - 1
		}
	}

	if y > MaxNorth {
		return fmt.Errorf("the move is out of the north boundary: (%d, %d)", x, y)
	}

	if y < MinNorth {
		return fmt.Errorf("the move is out of the south boundary: (%d, %d)", x, y)
	}

	if x > MaxEast {
		return fmt.Errorf("the move is out of the east bound: (%d, %d)", x, y)
	}

	if x < MinEast {
		return fmt.Errorf("the move is out of the west bound: (%d, %d)", x, y)
	}

	return nil
}

const (
	MaxEast  = 9
	MinEast  = 0
	MaxNorth = 9
	MinNorth = 0
)

const (
	North string = "N"
	East         = "E"
	South        = "S"
	West         = "W"
)
