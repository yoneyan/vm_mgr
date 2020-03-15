package main

import (
	"github.com/gorilla/mux"
	"github.com/yoneyan/vm_mgr/ggate/data"
	"log"
	"net/http"
	"strconv"
)

// GetConnection DBとのコネクションを張る
func GetID(r *http.Request) (id int, err error) {
	vars := mux.Vars(r)
	return strconv.Atoi(vars["id"])
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/api/v1/token", data.GenerateToken).Methods("POST")
	router.HandleFunc("/api/v1/token", data.DeleteToken).Methods("DELETE")
	router.HandleFunc("/api/v1/token/check", data.CheckToken).Methods("GET")

	//router.HandleFunc("/users", findAllUsers)
	//router.HandleFunc("/users/{id}", findByID)
	log.Fatal(http.ListenAndServe(":8080", router))
}
