package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	UserRoutes(r)
	AuthRoutes(r)
	ProfileRoutes(r)
	ProductRoutes(r)
	TopingRoutes(r)
	CartRoutes(r)
	Transaction(r)
}
