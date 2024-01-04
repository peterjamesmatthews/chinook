package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pjm.dev/chinook/internal/db/model"
)

func handleGetCustomers(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get customers from database
	customers := []model.Customer{}
	err = chinook.Find(&customers).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get customers: %w", err).Error()))
		return
	}

	// respond with customers
	HandleWritingJSONToResponse(w, customers)
}

func handleGetCustomer(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}
	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable {%s} is not an integer", idVar)))
		return
	}
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}
	// get customer from database
	customer := model.Customer{CustomerID: int32(id)}
	err = chinook.First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("customer %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get customer: %w", err).Error()))
		return
	}
	// respond with customer
	HandleWritingJSONToResponse(w, customer)
}

func handleCreateCustomer(w http.ResponseWriter, r *http.Request) {
	// validate customer
	customer := model.Customer{}
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode customer: %w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// create customer
	err = chinook.Create(&customer).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create customer: %w", err).Error()))
		return
	}

	// respond with customer
	HandleWritingJSONToResponse(w, customer)
}

func handlePatchCustomer(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}
	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// validate patch
	patch := model.Customer{}
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode customer: %w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get customer
	customer := model.Customer{CustomerID: int32(id)}
	err = chinook.First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("customer %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get customer: %w", err).Error()))
		return
	}

	// patch customer
	err = chinook.Model(&customer).Updates(patch).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to patch customer: %w", err).Error()))
		return
	}

	// respond with customer
	HandleWritingJSONToResponse(w, customer)
}

func handleDeleteCustomer(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}
	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/customers/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get customer
	customer := model.Customer{CustomerID: int32(id)}
	err = chinook.First(&customer).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("customer %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get customer: %w", err).Error()))
		return
	}

	// delete customer
	err = chinook.Delete(&customer).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete customer: %w", err).Error()))
		return
	}

	// respond with deleted customer
	HandleWritingJSONToResponse(w, customer)
}
