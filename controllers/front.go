package controllers

import (
	"encoding/json"
	"io"
	"net/http"
)

// RegisterControllers is exported 
func RegisterControllers() {

	uc := newUserController()

	http.Handle("/users", *uc)
	http.Handle("/users/", *uc)
}

func encodeResponseAsJSON(data interface{}, w io.Writer) {

	enc := json.NewEncoder(w)
	enc.Encode(data)
}
