package main

import (
	// "fmt"
	"github.com/codegangsta/negroni"
	// "github.com/nehathakur123/attendance/auth"
	"github.com/nehathakur123/attendance/routers"
	"github.com/nehathakur123/attendance/settings"
	"net/http"
)

func main() {
	settings.Init()
	router := routers.InitRoutes()
	n := negroni.Classic()
	n.UseHandler(router)
	http.ListenAndServe(":5000", n)

}
