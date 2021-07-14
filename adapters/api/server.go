package api

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/wy7-source/iti-challenge/adapters/dto"
	"github.com/wy7-source/iti-challenge/application"
)

type Server struct {
	Port string
}

func MakeServer(port string) *Server {
	return &Server{Port: port}
}

func (s Server) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Call service, that call Pasword Type and Validate, than return response here.
	var password application.Password
	var passwordReceived dto.PasswordDto
	err := json.NewDecoder(r.Body).Decode(&passwordReceived)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responseFormater(err.Error()))
		return
	}
	err = passwordReceived.Hydrate(&password)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("false"))
		return
	} else {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("true"))
	}
}

func (s Server) Serve() {
	http.HandleFunc("/validate", s.ValidateHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
