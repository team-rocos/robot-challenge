package service

import (
	"testing"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/mock"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
)

func TestRobotService_ValidateMovement(t *testing.T) {
	tests := []struct {
		name    string
		robot   thirdparty.Robot
		command string
		wantErr bool
	}{
		{
			name:    "valid move",
			robot:   mock.NewMockRobot(0, 0),
			command: "N N N N N",
			wantErr: false,
		},
		{
			name:    "move out of north boundary",
			robot:   mock.NewMockRobot(0, 8),
			command: "N N N N N",
			wantErr: true,
		},
		{
			name:    "move out of east boundary",
			robot:   mock.NewMockRobot(8, 0),
			command: "E E E E E",
			wantErr: true,
		},
		{
			name:    "move out of south boundary",
			robot:   mock.NewMockRobot(2, 0),
			command: "S S S S S",
			wantErr: true,
		},
		{
			name:    "move out of west boundary",
			robot:   mock.NewMockRobot(0, 2),
			command: "W W W W",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRobotService(tt.robot)

			if err := s.ValidateMovement(tt.command); (err != nil) != tt.wantErr {
				t.Errorf("validateMovement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRobotService_EnqueueTask(t *testing.T) {
	tests := []struct {
		name    string
		robot   thirdparty.Robot
		wantErr bool
	}{
		{
			name:  "enqueues successfully",
			robot: mock.NewMockRobot(1, 2),
		},
		{
			name:    "enqueues with error",
			robot:   mock.NewMockRobot(1, 2).WithHasEnqueueTaskError(true),
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewRobotService(tt.robot)
			dummyCommand := "E"
			if _, err := s.EnqueueTask(dummyCommand); (err != nil) != tt.wantErr {
				t.Errorf("EnqueueTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
