package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"pjm.dev/chinook/internal/db/model"
)

func handleGetEmployees(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get employees from database
	employees := []model.Employee{}
	err = chinook.Find(&employees).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get employees\n%w", err).Error()))
		return
	}

	// respond with employees
	HandleWritingJSONToResponse(w, employees)
}

func handleGetEmployee(w http.ResponseWriter, r *http.Request) {
	// get id from path variable
	id, err := handleGettingIntFromPathVariable(w, r, "id")
	if err != nil {
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get employee from database
	employee := model.Employee{EmployeeID: int32(id)}
	err = chinook.First(&employee).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("employee %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get employee %d\n%w", id, err).Error()))
		return
	}

	// respond with employee
	HandleWritingJSONToResponse(w, employee)
}

func handleCreateEmployee(w http.ResponseWriter, r *http.Request) {
	// decode employee from request body
	employee := model.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&employee); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode employee from request body\n%w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// create employee
	if err = chinook.Create(&employee).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create employee\n%w", err).Error()))
		return
	}

	// respond with created employee
	HandleWritingJSONToResponse(w, employee)
}

func handlePatchEmployee(w http.ResponseWriter, r *http.Request) {
	// get id
	id, err := handleGettingIntFromPathVariable(w, r, "id")
	if err != nil {
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get employee
	employee := model.Employee{EmployeeID: int32(id)}
	if err = chinook.First(&employee).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("employee %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get employee\n%w", err).Error()))
		return
	}

	// decode patch
	patch := model.Employee{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode patch from request body\n%w", err).Error()))
		return
	}

	// patch employee
	if err = chinook.Model(&employee).Updates(patch).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to patch employee\n%w", err).Error()))
		return
	}

	// respond with patched employee
	HandleWritingJSONToResponse(w, employee)
}

func handleDeleteEmployee(w http.ResponseWriter, r *http.Request) {
	// get id from path variable
	id, err := handleGettingIntFromPathVariable(w, r, "id")
	if err != nil {
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get employee
	employee := model.Employee{EmployeeID: int32(id)}
	if err = chinook.First(&employee).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("employee %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get employee %d\n%w", id, err).Error()))
		return
	}

	// delete employee
	if err = chinook.Delete(&employee).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete employee\n%w", err).Error()))
		return
	}

	// respond with deleted employee
	HandleWritingJSONToResponse(w, employee)
}
