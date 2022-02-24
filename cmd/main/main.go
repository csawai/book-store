package main

import (
	"log"
	"net/http"
	"github.com/csawai/book-store/pkg/routes"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialetcsts/postgres"
	"github.com/gorilla/mux"

)

func main (){
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

log.Fatal(http.ListenAndServe ("Localhost:9010", r))

}