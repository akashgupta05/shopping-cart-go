package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	log "github.com/sirupsen/logrus"
)

type Response struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Error   string      `json:"error,omitempty"`
}

func ReadBodyBytes(r *http.Request) ([]byte, error) {
	bodyByte, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Warn("Error reading request body", err.Error())
		return []byte(``), err
	}

	return bodyByte, nil
}

func RespondWithError(rw http.ResponseWriter, statusCode int, err error) {
	responseBytes, err := json.Marshal(Response{Success: false, Error: err.Error()})
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(statusCode)
	_, err = rw.Write(responseBytes)
	if err != nil {
		log.Errorf("failed to write response")
	}
}

func RespondWithSuccess(rw http.ResponseWriter, response *Response) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusOK)
	_, err = rw.Write(responseBytes)
	if err != nil {
		log.Errorf("failed to write response")
	}
}

func RespondWithJWTSuccess(rw http.ResponseWriter, token string, expiresAt *time.Time, response *Response) {
	http.SetCookie(rw, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: *expiresAt,
	})
	RespondWithSuccess(rw, response)
}
