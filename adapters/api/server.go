// Package api provides one way to interact with application.
// More Specificly, one simple HTTP Server/API.
package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/wy7-source/iti-challenge/adapters/dto"
	"github.com/wy7-source/iti-challenge/application/domain"
)

// Server is our HTTP server.
type Server struct {
	Port            string
	PasswordService domain.PasswordServiceInterface
}

// MakeServer is like a constructor of our HTTP server.
func MakeServer(port string, passwordService domain.PasswordServiceInterface) *Server {
	return &Server{Port: port, PasswordService: passwordService}
}

// ValidateHandler is the main handler to interact with the application.
// Here, we receive a Dto and hydrate the entity, to validate.
// It write the result of the validation service to client.
func (s Server) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Call service, that call Pasword Type and Validate, than return response here.
	var password domain.Password
	var passwordReceived dto.PasswordDto

	err := json.NewDecoder(r.Body).Decode(&passwordReceived)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		_, err = w.Write(responseFormater(err.Error()))
		if err != nil {
			panic(err.Error())
		}
		return
	}
	err = passwordReceived.Hydrate(&password)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(responseFormater("false"))
		if err != nil {
			panic(err.Error())
		}
		return
	}

	result := s.PasswordService.Validate(&password)
	_, err = w.Write(responseFormater(strconv.FormatBool(result)))
	if err != nil {
		panic(err.Error())
	}
}

// Serve start the server to listen on port provided.
func (s Server) Serve() {
	http.HandleFunc("/validate", s.ValidateHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
