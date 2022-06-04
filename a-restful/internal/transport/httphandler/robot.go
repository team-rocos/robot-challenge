package httphandler

import (
	"encoding/json"
	"net/http"

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

	//token, err := h.authService.Login(service.LoginInput{
	//	Username: payload.Username,
	//	Password: payload.Password,
	//})
	//if err != nil {
	//	return NewHTTPError(err, http.StatusInternalServerError, "failed to login")
	//}
	//
	//resp, err := json.Marshal(LoginResponse{
	//	Token: token,
	//})
	//if err != nil {
	//	return NewHTTPError(err, http.StatusInternalServerError, "unable to marshall response")
	//}

	w.WriteHeader(http.StatusOK)
	//w.Write(resp)

	h.log.Info("enqueue task")
	return nil
}
