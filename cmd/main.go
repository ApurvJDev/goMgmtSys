package main

import (
	"log"
	"net/http"

	"github.com/ApurvJDev/goMgmtSys/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// we are going to create a server ie define the localhost
// we are going to tell where our routers reside

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)
	log.Fatal(http.ListenAndServe()) // helps us create the server
}
