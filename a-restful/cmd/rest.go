package main

import (
	"net/http"
	"os"

	"github.com/joho/godotenv"

	"github.com/edwardkcyu/robot-challenge/a-restful/internal/service"
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/transport/httphandler"
	"github.com/edwardkcyu/robot-challenge/a-restful/internal/util"
	"github.com/edwardkcyu/robot-challenge/a-restful/thirdparty"
	"github.com/gorilla/mux"
)

func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := util.NewServer()

	robotHandler := httphandler.NewRobotHandler(service.NewRobotService(thirdparty.NewRealRobot()))

	r := mux.NewRouter()
	r.HandleFunc("/tasks", wrapHandler(robotHandler.EnqueueTaskHandler)).Methods("POST")
	r.HandleFunc("/tasks", wrapHandler(robotHandler.CancelTaskHandler)).Methods("DELETE")
	r.HandleFunc("/tasks", wrapHandler(robotHandler.QueryTaskHandler)).Methods("GET")

	server.Start(port, r)
}

func wrapHandler(handler func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return httphandler.ResponseHandler(
		httphandler.ErrorHandler(
			handler,
		),
	)
}
