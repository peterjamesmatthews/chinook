package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type ChinookRouter struct {
	r http.Handler
}

// ServeHTTP implements the http.Handler interface.
func (cr *ChinookRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	cr.r.ServeHTTP(w, r)
}

func NewChinookRouter() *ChinookRouter {
	r := mux.NewRouter()
	registerAllRoutes(r)
	return &ChinookRouter{r}
}

func registerAllRoutes(r *mux.Router) {
	registerAlbumRoutes(r)
	// registerArtistRoutes(r)
	// registerCustomerRoutes(r)
	// registerEmployeeRoutes(r)
	// registerGenreRoutes(r)
	// registerInvoiceRoutes(r)
	// registerInvoiceLineRoutes(r)
	// registerMediaTypeRoutes(r)
	// registerPlaylistRoutes(r)
	// registerPlaylistTrackRoutes(r)
	// registerTrackRoutes(r)
}
