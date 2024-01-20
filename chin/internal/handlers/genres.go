package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"gorm.io/gorm"

	"pjm.dev/chin/internal/db/model"
)

func handleGetGenres(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get genres from database
	genres := []model.Genre{}
	if err := chinook.Find(&genres).Error; err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get genres\n%w", err).Error()))
	}

	// respond with genres
	HandleWritingJSONToResponse(w, genres)
}

func handleGetGenre(w http.ResponseWriter, r *http.Request) {
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

	// get genre from database
	genre := model.Genre{GenreID: int32(id)}
	if err := chinook.First(&genre).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("genre %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get genre %d\n%w", id, err).Error()))
		return
	}

	// respond with genre
	HandleWritingJSONToResponse(w, genre)
}

func handleCreateGenre(w http.ResponseWriter, r *http.Request) {
	// decode genre from request body
	genre := model.Genre{}
	err := json.NewDecoder(r.Body).Decode(&genre)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode genre\n%w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// create genre
	if err := chinook.Create(&genre).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create genre\n%w", err).Error()))
		return
	}

	// respond with created genre
	HandleWritingJSONToResponse(w, genre)
}

func handlePatchGenre(w http.ResponseWriter, r *http.Request) {
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

	// get genre from database
	genre := model.Genre{GenreID: int32(id)}
	if err := chinook.First(&genre).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("genre %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get genre %d\n%w", id, err).Error()))
		return
	}

	// decode patch from request body
	patch := model.Genre{}
	if err := json.NewDecoder(r.Body).Decode(&patch); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode patch\n%w", err).Error()))
		return
	}

	// patch genre
	if err := chinook.Model(&genre).Updates(patch).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to patch genre %d\n%w", id, err).Error()))
		return
	}

	// respond with patched genre
	HandleWritingJSONToResponse(w, genre)
}

func handleDeleteGenre(w http.ResponseWriter, r *http.Request) {
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

	// get genre from database
	genre := model.Genre{GenreID: int32(id)}
	if err := chinook.First(&genre).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("genre %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get genre %d\n%w", id, err).Error()))
		return
	}

	// delete genre
	if err := chinook.Delete(&genre).Error; err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete genre %d\n%w", id, err).Error()))
		return
	}

	// respond with deleted genre
	HandleWritingJSONToResponse(w, genre)
}
