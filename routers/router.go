package routers

import (
	"github.com/gorilla/mux"
	//"github.com/nehathakur123/attendance/services/models"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router = SetHelloRoutes(router)
	router = SetAuthenticationRoutes(router)
	return router
}
