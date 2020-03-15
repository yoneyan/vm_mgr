package main

import (
	"github.com/gorilla/mux"
	"github.com/yoneyan/vm_mgr/ggate/data"
	"log"
	"net/http"
)

func main() {

	router := mux.NewRouter().StrictSlash(true)
	//token
	router.HandleFunc("/api/v1/token", data.GenerateToken).Methods("POST")
	router.HandleFunc("/api/v1/token", data.DeleteToken).Methods("DELETE")
	router.HandleFunc("/api/v1/token/check", data.CheckToken).Methods("GET")
	//router.HandleFunc("/api/v1/token",data.GetUserVM).Methods("GET")
	//vm
	router.HandleFunc("/api/v1/vm", data.GetUserVM).Methods("GET")
	//router.HandleFunc("/api/v1/vm/group", data.GetUserVM).Methods("GET")
	//router.HandleFunc("/api/v1/vm", data.GetUserVM).Methods("GET")

	//router.HandleFunc("/users", findAllUsers)
	//router.HandleFunc("/users/{id}", findByID)
	log.Fatal(http.ListenAndServe(":8080", router))
}
