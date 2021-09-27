package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"payment-service/api/controllers"
)

func HandleRequests(c controllers.IController) {

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", controllers.HomePage)
	myRouter.HandleFunc("/articles", c.GetAll)
	myRouter.HandleFunc("/article", c.CreateNew).Methods("POST")
	myRouter.HandleFunc("/article/{id}", c.DeleteOne).Methods("DELETE")
	myRouter.HandleFunc("/article/ui={id}", c.GetOne)
	log.Fatal(http.ListenAndServe(":10000", myRouter))

}
