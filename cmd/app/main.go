package main

import (
	"fmt"
	"log"
	"net/http"

	controller "github.com/getground/tech-tasks/backend/cmd/app/controllers"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)


func main() {
	router := mux.NewRouter() 
	router.HandleFunc("/tables", controller.InsertTable).Methods("POST")
	router.HandleFunc("/guest_list/{name}", controller.InsertGuest).Methods("POST")
	router.HandleFunc("/guest_list", controller.GetGuestlist).Methods("GET")
	router.HandleFunc("/guests/{name}", controller.CheckInGuest).Methods("PUT")
	router.HandleFunc("/guests", controller.GetPartyGuests).Methods("GET")
	router.HandleFunc("/seats_empty", controller.GetSeatsEmpty).Methods("GET")
	router.HandleFunc("/guests/{name}", controller.DeleteGuest).Methods("DELETE")
	http.Handle("/", router)
	fmt.Println("Connected to port 3000")
	log.Fatal(http.ListenAndServe(":3000", router))
}



