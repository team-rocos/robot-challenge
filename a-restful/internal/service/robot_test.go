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
			name: "valid move",
			robot: &mock.MockRobot{
				State: thirdparty.RobotState{
					X: 0,
					Y: 0,
				},
			},
			command: "N N N N N",
			wantErr: false,
		},
		{
			name: "move out of north boundary",
			robot: &mock.MockRobot{
				State: thirdparty.RobotState{
					X: 0,
					Y: 8,
				},
			},
			command: "N N N N N",
			wantErr: true,
		},
		{
			name: "move out of east boundary",
			robot: &mock.MockRobot{
				State: thirdparty.RobotState{
					X: 8,
					Y: 0,
				},
			},
			command: "E E E E E",
			wantErr: true,
		},
		{
			name: "move out of south boundary",
			robot: &mock.MockRobot{
				State: thirdparty.RobotState{
					X: 2,
					Y: 0,
				},
			},
			command: "S S S S S",
			wantErr: true,
		},
		{
			name: "move out of west boundary",
			robot: &mock.MockRobot{
				State: thirdparty.RobotState{
					X: 0,
					Y: 2,
				},
			},
			command: "W W W W",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &RobotService{
				robot: tt.robot,
			}
			if err := s.ValidateMovement(tt.command); (err != nil) != tt.wantErr {
				t.Errorf("validateMovement() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
