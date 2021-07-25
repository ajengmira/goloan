package main

import (
	"goloan/app/controllers"
	_ "goloan/app/models"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	//"github.com/go-sql-driver/mysql"
)

func main() {

	router := mux.NewRouter()
	router.HandleFunc("/api/loans", controllers.GetLoans).Methods("GET")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.GetLoan).Methods("GET")
	router.HandleFunc("/api/loan", controllers.CreateLoan).Methods("POST")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.UpdateLoan).Methods("PUT")
	router.HandleFunc("/api/loan/approve/{id:[0-9]+}", controllers.ApproveLoan).Methods("PUT")
	router.HandleFunc("/api/loan/{id:[0-9]+}", controllers.DeleteLoan).Methods("DELETE")

	err := http.ListenAndServe(os.Getenv("app_ip")+":"+os.Getenv("app_port"), router)
	if err != nil {
		panic(err)
	}
}
