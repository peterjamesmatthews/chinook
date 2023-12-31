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

func registerAlbumRoutes(r *mux.Router) {
	r.HandleFunc("/albums", getAlbums).
		Methods("GET")

	r.HandleFunc("/albums/{id:[0-9]+}", getAlbum).
		Methods("GET")

	r.HandleFunc("/albums", createAlbum).
		Methods("POST")

	r.HandleFunc("/albums/{id:[0-9]+}", patchAlbum).
		Methods("PATCH")

	r.HandleFunc("/albums/{id:[0-9]+}", deleteAlbum).
		Methods("DELETE")
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	// get chinook from context
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database: %w", err).Error()))
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
	if err := WriteJSONToResponse(w, albums); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to write albums to response: %w", err).Error()))
		return
	}
}

func getAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("missing /albums/{id} path variable in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id %s is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database: %w", err).Error()))
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
	if err := WriteJSONToResponse(w, album); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to write album to response: %w", err).Error()))
	}
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
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database: %w", err).Error()))
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
	if err := WriteJSONToResponse(w, album); err != nil {
		w.Write([]byte(fmt.Errorf("failed to write created album to response: %w", err).Error()))
	}

	w.WriteHeader(http.StatusCreated)
}

func patchAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("missing /albums/{id} path variable in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id %s is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database: %w", err).Error()))
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
	if err := WriteJSONToResponse(w, album); err != nil {
		w.Write([]byte(fmt.Errorf("failed to write patched album to response: %w", err).Error()))
	}
}

func deleteAlbum(w http.ResponseWriter, r *http.Request) {
	// validate id
	idVar, ok := mux.Vars(r)["id"]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("missing /albums/{id} path variable in request path %s", r.URL.Path)))
		return
	}

	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("id %s is not an integer", idVar)))
		return
	}

	// get chinook from context
	chinook, err := GetChinookFromContext(r.Context())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get database: %w", err).Error()))
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
	if err := WriteJSONToResponse(w, album); err != nil {
		w.Write([]byte(fmt.Errorf("failed to write deleted album to response: %w", err).Error()))
	}
}
