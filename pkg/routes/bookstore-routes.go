
package routes

import (
	"github.com/gorilla/mux"
	"gihub.com/csawai/book-store/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.router){
	router.HandlFunc("/book", controllers.CreateBook).Methods(POST)
	router.HandlFunc("/book", controllers.GetBook).Methods(GET)
	router.HandlFunc("/book"{bookId}, controllers.GetBookByIdeateBook).Methods(GET)
	router.HandlFunc("/book"{bookId}, controllers.UpdateBook).Methods(PUT)
	router.HandlFunc("/book"{bookId}, controllers.DeleteBook).Methods(POST)






}


