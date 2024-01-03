package http

import (
	"github.com/gorilla/mux"
)

func RegisterChinookRoutes(r *mux.Router) {
	registerAlbumsRoutes(r)
	registerArtistsRoutes(r)
	registerCustomersRoutes(r)
}

func registerAlbumsRoutes(r *mux.Router) {
	r.HandleFunc("/albums", handleGetAlbums).
		Methods("GET")

	r.HandleFunc("/albums/{id:[0-9]+}", handleGetAlbum).
		Methods("GET")

	r.HandleFunc("/albums", handleCreateAlbum).
		Methods("POST")

	r.HandleFunc("/albums/{id:[0-9]+}", handlePatchAlbum).
		Methods("PATCH")

	r.HandleFunc("/albums/{id:[0-9]+}", handleDeleteAlbum).
		Methods("DELETE")
}

func registerArtistsRoutes(r *mux.Router) {
	r.HandleFunc("/artists", handleGetArtists).
		Methods("GET")

	r.HandleFunc("/artists/{id:[0-9]+}", handleGetArtist).
		Methods("GET")

	r.HandleFunc("/artists", handleCreateArtist).
		Methods("POST")

	r.HandleFunc("/artists/{id:[0-9]+}", handlePatchArtist).
		Methods("PATCH")

	r.HandleFunc("/artists/{id:[0-9]+}", handleDeleteArtist).
		Methods("DELETE")
}

func registerCustomersRoutes(r *mux.Router) {
	r.HandleFunc("/customers", handleGetCustomers).
		Methods("GET")

	r.HandleFunc("/customers/{id:[0-9]+}", handleGetCustomer).
		Methods("GET")

	r.HandleFunc("/customers", handleCreateCustomer).
		Methods("POST")

	r.HandleFunc("/customers/{id:[0-9]+}", handlePatchCustomer).
		Methods("PATCH")

	r.HandleFunc("/customers/{id:[0-9]+}", handleDeleteCustomer).
		Methods("DELETE")
}
