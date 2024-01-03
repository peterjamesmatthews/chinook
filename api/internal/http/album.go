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

func getAlbums(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get albums
	albums := []model.Album{}
	err = chinook.Find(&albums).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get albums: %w", err).Error()))
		return
	}

	// respond with albums
	handleWritingJSONToResponse(w, albums)
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/albums/{id} path variable missing in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/albums/{id} path variable {%s} is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get album
	album := model.Album{AlbumID: int32(id)}
	err = chinook.First(&album).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("album %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get album: %w", err).Error()))
		return
	}

	// respond with album
	handleWritingJSONToResponse(w, album)
}

func createAlbum(w http.ResponseWriter, r *http.Request) {
	// validate album
	album := model.Album{}
	err := json.NewDecoder(r.Body).Decode(&album)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode album: %w", err).Error()))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// create album
	err = chinook.Create(&album).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create album: %w", err).Error()))
		return
	}

	// respond with created album
	if err = handleWritingJSONToResponse(w, album); err != nil {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func patchAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/albums/{id} path variable missing from request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id %s is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get album
	album := model.Album{AlbumID: int32(id)}
	err = chinook.First(&album).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("album %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get album: %w", err).Error()))
		return
	}

	// decode patch
	patch := model.Album{}
	err = json.NewDecoder(r.Body).Decode(&patch)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Errorf("failed to decode album: %w", err).Error()))
		return
	}

	// patch album
	err = chinook.Model(&album).Updates(patch).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to patch album: %w", err).Error()))
		return
	}

	// respond with patched album
	handleWritingJSONToResponse(w, album)
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("/albums/{id} path variable missing from request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id %s is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := handleGettingChinookFromContext(w, r)
	if err != nil {
		return
	}

	// get album
	album := model.Album{AlbumID: int32(id)}
	if err := chinook.First(&album).Error; err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("album %d not found", id)))
		return
	}

	// delete album
	err = chinook.Delete(&album).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete album: %w", err).Error()))
		return
	}

	// respond with deleted album
	handleWritingJSONToResponse(w, album)
}
