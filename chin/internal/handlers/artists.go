package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"pjm.dev/chin/internal/nook/model"
)

func handleGetArtists(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get artists from database
	artists := []model.Artist{}
	err = chinook.Find(&artists).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get artists: %w", err).Error()))
		return
	}

	// respond with artists
	HandleWritingJSONToResponse(w, artists)
}

func handleGetArtist(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get artist from database
	artist := model.Artist{ArtistID: int32(id)}
	err = chinook.First(&artist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf("artist with id %d not found", id).Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get artist: %w", err).Error()))
		return
	}

	// respond with artist
	HandleWritingJSONToResponse(w, artist)
}

func handleCreateArtist(w http.ResponseWriter, r *http.Request) {
	// validate album
	artist := model.Artist{}
	err := json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode artist: %w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// create artist
	err = chinook.Create(&artist).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create artist: %w", err).Error()))
		return
	}

	// respond with artist
	if err = HandleWritingJSONToResponse(w, artist); err != nil {
		return
	}
}

func handlePatchArtist(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get artist
	artist := model.Artist{ArtistID: int32(id)}
	err = chinook.First(&artist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf("artist with id %d not found", id).Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get artist: %w", err).Error()))
		return
	}

	// decode patch
	err = json.NewDecoder(r.Body).Decode(&artist)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode artist: %w", err).Error()))
		return
	}

	// update artist
	err = chinook.Save(&artist).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to update artist: %w", err).Error()))
		return
	}

	// respond with artist
	HandleWritingJSONToResponse(w, artist)
}

func handleDeleteArtist(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/artists/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get artist
	artist := model.Artist{ArtistID: int32(id)}
	err = chinook.First(&artist).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Errorf("artist with id %d not found", id).Error()))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get artist: %w", err).Error()))
		return
	}

	// delete artist
	err = chinook.Delete(&artist).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete artist: %w", err).Error()))
		return
	}

	// respond with artist
	HandleWritingJSONToResponse(w, artist)
}
