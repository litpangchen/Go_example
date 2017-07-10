package main

import (
	"go_web_programming/016_building_restful_api/taskmanager/common"
	"go_web_programming/016_building_restful_api/taskmanager/routers"
	"github.com/codegangsta/negroni"
	"net/http"
	"log"
)

func main() {

	common.Startup()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	server := &http.Server{
		Addr:    common.AppConfig.Server,
		Handler: n,
	}
	log.Println("Server started!")
	server.ListenAndServe()
}
