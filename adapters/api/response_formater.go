package api

import "encoding/json"

// responseFormater convert a string to one json response message.
func responseFormater(msg string) []byte {
	error := struct {
		Message string `json:"message"`
	}{
		msg,
	}
	r, err := json.Marshal(error)
	if err != nil {
		return []byte(err.Error())
	}
	return r
}
