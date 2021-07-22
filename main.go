package main

import (
	"goloan/app/controllers"
	_ "goloan/app/models"
	"net/http"

	"github.com/gorilla/mux"
	//"github.com/go-sql-driver/mysql"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/loans", controllers.GetLoans).Methods("GET")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.GetLoan).Methods("GET")
	router.HandleFunc("/api/loan", controllers.CreateLoan).Methods("POST")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.UpdateLoan).Methods("PUT")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.DeleteLoan).Methods("DELETE")

	err := http.ListenAndServe("127.0.0.1:9000", router)
	if err != nil {
		panic(err)
	}
}
