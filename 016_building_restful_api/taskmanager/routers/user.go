package routers

import (
	"github.com/gorilla/mux"
	"go_web_programming/016_building_restful_api/taskmanager/controllers"
)

func SetUserRoutes(router *mux.Router) *mux.Router {
	router.HandleFunc("/users/register", controllers.Register).Methods("POST")
	router.HandleFunc("/users/login", controllers.Login).Methods("POST")
	return router
}
