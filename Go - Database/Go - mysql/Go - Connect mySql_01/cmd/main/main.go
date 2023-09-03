package main

import (
	"go/pkg/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
