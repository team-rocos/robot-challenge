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

		headers := ResponseHeaders()
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		w.WriteHeader(httpError.Status)
		w.Write(body)
		log.Println(httpError)
	}
}

func ResponseHandler(next func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		headers := ResponseHeaders()
		for k, v := range headers {
			w.Header().Set(k, v)
		}
		next(w, r)
	}
}
