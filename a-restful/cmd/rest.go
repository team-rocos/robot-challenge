package main

import (
	"net/http"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/service"
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/transport/httphandler"
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
	"github.com/gorilla/mux"
)

func main() {
	server := util.NewServer()

	robotHandler := httphandler.NewRobotHandler(service.NewRobotService(thirdparty.NewRealRobot()))

	r := mux.NewRouter()
	r.HandleFunc("/commands", wrapHandler(robotHandler.EnqueueTaskHandler)).Methods("POST")
	//s.router.HandleFunc("/commands", httphandler.ErrorHandler(robotHandler.CommandHandler)).Methods("GET")
	//s.router.HandleFunc("/commands", httphandler.ErrorHandler(robotHandler.CommandHandler)).Methods("PUT")

	server.Start("8080", r)
}

func wrapHandler(handler func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return httphandler.ResponseHandler(
		httphandler.ErrorHandler(
			handler,
		),
	)
}
