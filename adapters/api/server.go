package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/wy7-source/iti-challenge/adapters/dto"
	"github.com/wy7-source/iti-challenge/application/domain"
)

type Server struct {
	Port            string
	PasswordService domain.PasswordServiceInterface
}

func MakeServer(port string, passwordService domain.PasswordServiceInterface) *Server {
	return &Server{Port: port, PasswordService: passwordService}
}

func (s Server) ValidateHandler(w http.ResponseWriter, r *http.Request) {
	// Call service, that call Pasword Type and Validate, than return response here.
	var password domain.Password
	var passwordReceived dto.PasswordDto

	err := json.NewDecoder(r.Body).Decode(&passwordReceived)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(responseFormater(err.Error()))
		return
	}
	err = passwordReceived.Hydrate(&password)
	if err != nil {
		w.WriteHeader(http.StatusOK)
		w.Write(responseFormater("false"))
		return
	}

	result := s.PasswordService.Validate(&password)
	w.Write(responseFormater(strconv.FormatBool(result)))
}
func (s Server) Serve() {
	http.HandleFunc("/validate", s.ValidateHandler)
	log.Fatal(http.ListenAndServe(s.Port, nil))
}
