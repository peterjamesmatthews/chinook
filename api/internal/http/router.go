package http

import (
	"github.com/gorilla/mux"
)

func RegisterChinookRoutes(r *mux.Router) {
	registerAlbumRoutes(r)
}
