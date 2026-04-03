// Package routes sets up the routes
package routes

import (
	"net/http"

	"micahasowata.com/restapi/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux, handler *handlers.Handler) {
	SetupHealthRoute(mux, handler)
}
