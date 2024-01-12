package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func handleGettingStringFromPathVariable(w http.ResponseWriter, r *http.Request, key string) (string, error) {
	variable, ok := mux.Vars(r)[key]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("missing variable %s from request path %s", key, r.URL.Path)))
		return "", errors.New("path variable missing in request path")
	}

	return variable, nil
}

func handleGettingIntFromPathVariable(w http.ResponseWriter, r *http.Request, key string) (int, error) {
	variable, err := handleGettingStringFromPathVariable(w, r, key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.Atoi(variable)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("path variable %s is not an integer", variable)))
		return 0, errors.New("path variable is not an integer")
	}

	return i, nil
}
