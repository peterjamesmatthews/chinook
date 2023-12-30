package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	"pjm.dev/chinook/internal/db"
	"pjm.dev/chinook/internal/db/model"
)

func registerAlbumRoutes(r *mux.Router) {
	r.HandleFunc("/albums", getAlbums).
		Methods("GET")

	r.HandleFunc("/albums/{id:[0-9]+}", getAlbum).
		Methods("GET")

	r.HandleFunc("/albums", createAlbum).
		Methods("POST")

	r.HandleFunc("/albums/{id:[0-9]+}", updateAlbum).
		Methods("PUT")

	r.HandleFunc("/albums/{id:[0-9]+}", deleteAlbum).
		Methods("DELETE")
}

func getAlbums(w http.ResponseWriter, r *http.Request) {
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get db: %w", err).Error()))
		return
	}

	albums := []model.Album{}
	err = db.Find(&albums).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get albums: %w", err).Error()))
		return
	}

	content, err := json.Marshal(albums)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to marshal albums: %w", err).Error()))
		return
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(content)))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to write albums: %w", err).Error()))
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

	// get album
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get db: %w", err).Error()))
		return
	}

	album := model.Album{AlbumID: int32(id)}
	err = db.First(&album).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(fmt.Sprintf("album %d not found", id)))
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get album: %w", err).Error()))
		return
	}

	// return album
	content, err := json.Marshal(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to marshal album: %w", err).Error()))
		return
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(content)))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to write album: %w", err).Error()))
		return
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

	// connect to db
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get db: %w", err).Error()))
		return
	}

	// create album
	err = db.Create(&album).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to create album: %w", err).Error()))
		return
	}

	// return album
	content, err := json.Marshal(album)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to marshal album: %w", err).Error()))
		return
	}

	w.Header().Set("Content-Length", fmt.Sprint(len(content)))
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(content)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to write album: %w", err).Error()))
		return
	}
}

func updateAlbum(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
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

	// connect to db
	db, err := db.GetDB()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to get db: %w", err).Error()))
		return
	}

	// delete album
	album := model.Album{AlbumID: int32(id)}
	err = db.Delete(&album).Error
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Errorf("failed to delete album: %w", err).Error()))
		return
	}
}
