package handlers

import (
	"net/http"

	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router) {
	registerAlbumsRoutes(r)
	registerArtistsRoutes(r)
	registerCustomersRoutes(r)
	registerEmployeeRoutes(r)
	registerGenreRoutes(r)
}

func registerAlbumsRoutes(r *mux.Router) {
	r.HandleFunc("/albums", handleGetAlbums).
		Methods(http.MethodGet)

	r.HandleFunc("/albums/{id:[0-9]+}", handleGetAlbum).
		Methods(http.MethodGet)

	r.HandleFunc("/albums", handleCreateAlbum).
		Methods(http.MethodPost)

	r.HandleFunc("/albums/{id:[0-9]+}", handlePatchAlbum).
		Methods(http.MethodPatch)

	r.HandleFunc("/albums/{id:[0-9]+}", handleDeleteAlbum).
		Methods(http.MethodDelete)
}

func registerArtistsRoutes(r *mux.Router) {
	r.HandleFunc("/artists", handleGetArtists).
		Methods(http.MethodGet)

	r.HandleFunc("/artists/{id:[0-9]+}", handleGetArtist).
		Methods(http.MethodGet)

	r.HandleFunc("/artists", handleCreateArtist).
		Methods(http.MethodPost)

	r.HandleFunc("/artists/{id:[0-9]+}", handlePatchArtist).
		Methods(http.MethodPatch)

	r.HandleFunc("/artists/{id:[0-9]+}", handleDeleteArtist).
		Methods(http.MethodDelete)
}

func registerCustomersRoutes(r *mux.Router) {
	r.HandleFunc("/customers", handleGetCustomers).
		Methods(http.MethodGet)

	r.HandleFunc("/customers/{id:[0-9]+}", handleGetCustomer).
		Methods(http.MethodGet)

	r.HandleFunc("/customers", handleCreateCustomer).
		Methods(http.MethodPost)

	r.HandleFunc("/customers/{id:[0-9]+}", handlePatchCustomer).
		Methods(http.MethodPatch)

	r.HandleFunc("/customers/{id:[0-9]+}", handleDeleteCustomer).
		Methods(http.MethodDelete)
}

func registerEmployeeRoutes(r *mux.Router) {
	r.HandleFunc("/employees", handleGetEmployees).
		Methods(http.MethodGet)

	r.HandleFunc("/employees/{id:[0-9]+}", handleGetEmployee).
		Methods(http.MethodGet)

	r.HandleFunc("/employees", handleCreateEmployee).
		Methods(http.MethodPost)

	r.HandleFunc("/employees/{id:[0-9]+}", handlePatchEmployee).
		Methods(http.MethodPatch)

	r.HandleFunc("/employees/{id:[0-9]+}", handleDeleteEmployee).
		Methods(http.MethodDelete)
}

func registerGenreRoutes(r *mux.Router) {
	r.HandleFunc("/genres", handleGetGenres).
		Methods(http.MethodGet)

	r.HandleFunc("/genres/{id:[0-9]+}", handleGetGenre).
		Methods(http.MethodGet)

	r.HandleFunc("/genres", handleCreateGenre).
		Methods(http.MethodPost)

	r.HandleFunc("/genres/{id:[0-9]+}", handlePatchGenre).
		Methods(http.MethodPatch)

	r.HandleFunc("/genres/{id:[0-9]+}", handleDeleteGenre).
		Methods(http.MethodDelete)
}
