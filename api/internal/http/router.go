package http

import (
	"github.com/gorilla/mux"
)

func RegisterChinookRoutes(r *mux.Router) {
	registerAlbumRoutes(r)
	registerArtistRoutes(r)
}

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

func registerArtistRoutes(r *mux.Router) {
	r.HandleFunc("/artists", getArtists).
		Methods("GET")

	r.HandleFunc("/artists/{id:[0-9]+}", getArtist).
		Methods("GET")

	r.HandleFunc("/artists", createArtist).
		Methods("POST")

	r.HandleFunc("/artists/{id:[0-9]+}", patchArtist).
		Methods("PATCH")

	r.HandleFunc("/artists/{id:[0-9]+}", deleteArtist).
		Methods("DELETE")
}
