package httphandler

import (
	"encoding/json"
	"net/http"

	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/service"
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
)

type RobotHandler struct {
	log          *util.Logger
	robotService service.IRobotService
}

func NewRobotHandler(robotService service.IRobotService) *RobotHandler {
	return &RobotHandler{
		log:          util.NewLogger("RobotHandler"),
		robotService: robotService,
	}
}

type EnqueueTaskRequest struct {
	Command string `json:"command"`
}

type EnqueueTaskResponse struct {
	TaskID string `json:"taskId"`
}

func (h *RobotHandler) EnqueueTaskHandler(w http.ResponseWriter, r *http.Request) error {
	var payload EnqueueTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "unable to decode request body")
	}

	err := h.robotService.ValidateMovement(payload.Command)
	if err != nil {
		return NewHTTPError(err, http.StatusBadRequest, err.Error())
	}

	taskID, err := h.robotService.EnqueueTask(payload.Command)
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "failed to enqueue task to robot")
	}

	resp, err := json.Marshal(EnqueueTaskResponse{
		TaskID: taskID,
	})
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "unable to marshall response")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}

type CancelTaskRequest struct {
	TaskID string `json:"taskId"`
}

func (h *RobotHandler) CancelTaskHandler(w http.ResponseWriter, r *http.Request) error {
	var payload CancelTaskRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		return NewHTTPError(err, http.StatusBadRequest, "unable to decode request body")
	}

	err := h.robotService.CancelTask(payload.TaskID)
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "failed to cancel task to robot")
	}

	w.WriteHeader(http.StatusOK)

	return nil
}

type QueryTaskResponse struct {
	Status thirdparty.TaskStatus `json:"status"`
}

func (h *RobotHandler) QueryTaskHandler(w http.ResponseWriter, r *http.Request) error {
	taskID := r.URL.Query().Get("taskId")

	taskStatus, err := h.robotService.QueryTask(taskID)
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "failed to query task")
	}

	resp, err := json.Marshal(QueryTaskResponse{
		Status: taskStatus,
	})
	if err != nil {
		return NewHTTPError(err, http.StatusInternalServerError, "unable to marshall response")
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resp)

	return nil
}
