package crow

import (
	"net/http"

	"gorm.io/gorm"
)

func WrapInCrow(handler http.Handler, getDB func(r *http.Request) *gorm.DB) (http.Handler, *Crow) {
	c := Crow{}
	wrapped := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		db := getDB(r)
		c.seed(db)
		handler.ServeHTTP(w, r)
		c.dump()
	})
	return wrapped, &c
}
