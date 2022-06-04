package httphandler

import (
	"errors"
	"log"
	"net/http"
)

func ErrorHandler(next func(http.ResponseWriter, *http.Request) error) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		err := next(w, r)
		if err == nil {
			return
		}

		var httpError *HTTPError
		//clientError, ok := err.(ClientError)
		ok := errors.As(err, &httpError)
		if !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("internal server error"))
			return
		}

		body, err := httpError.ResponseBody()
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		status, headers := httpError.ResponseHeaders() // Get http status code and headers.
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(status)
		w.Write(body)
		log.Println(httpError)
	}
}
