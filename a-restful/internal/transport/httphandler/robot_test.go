package httphandler

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/mock"

	"github.com/stretchr/testify/assert"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/service"
)

func TestRobotHandler_QueryTaskHandler(t *testing.T) {

	tests := []struct {
		name           string
		robotService   service.IRobotService
		wantErr        bool
		wantHttpStatus int
		wantTaskStatus thirdparty.TaskStatus
	}{
		{
			name:           "queries the task status",
			robotService:   service.NewRobotService(mock.NewMockRobot(0, 0)),
			wantHttpStatus: http.StatusOK,
			wantTaskStatus: thirdparty.TaskStatusExecuting,
		},
		{
			name:           "queries the task status with error",
			robotService:   service.NewRobotService(mock.NewMockRobot(0, 0).WithHasQueryTaskError(true)),
			wantHttpStatus: http.StatusInternalServerError,
			wantTaskStatus: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, "/commands?taskId=task1", nil)
			w := httptest.NewRecorder()

			h := NewRobotHandler(tt.robotService)
			ResponseHandler(ErrorHandler(h.QueryTaskHandler))(w, req)

			assert.Equal(t, tt.wantHttpStatus, w.Code)

			res := w.Result()
			defer res.Body.Close()
			data, err := ioutil.ReadAll(res.Body)
			assert.Nil(t, err, "expected error to be nil")

			var result QueryTaskResponse
			json.Unmarshal(data, &result)

			assert.Equal(t, tt.wantTaskStatus, result.Status)

		})
	}
}

func TestRobotHandler_CancelTaskHandler(t *testing.T) {

	tests := []struct {
		name           string
		robotService   service.IRobotService
		wantErr        bool
		wantHttpStatus int
	}{
		{
			name:           "cancels the task",
			robotService:   service.NewRobotService(mock.NewMockRobot(0, 0)),
			wantHttpStatus: http.StatusOK,
		},
		{
			name:           "cancels the task with error",
			robotService:   service.NewRobotService(mock.NewMockRobot(0, 0).WithHasCancelTaskError(true)),
			wantHttpStatus: http.StatusInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cancelTaskReq := CancelTaskRequest{
				TaskID: "task1",
			}

			var buf bytes.Buffer
			err := json.NewEncoder(&buf).Encode(cancelTaskReq)
			assert.Nil(t, err)

			req := httptest.NewRequest(http.MethodDelete, "/commands", &buf)
			w := httptest.NewRecorder()

			h := NewRobotHandler(tt.robotService)
			ResponseHandler(ErrorHandler(h.CancelTaskHandler))(w, req)

			assert.Equal(t, tt.wantHttpStatus, w.Code)
		})
	}
}
