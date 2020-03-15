package data

import (
	"encoding/json"
	"github.com/yoneyan/vm_mgr/ggate/client"
	"io/ioutil"
	"log"
	"net/http"
)

type UserAuth struct {
	User string `json:"user"`
	Pass string `json:"pass"`
}

func CheckToken(w http.ResponseWriter, r *http.Request) {
	log.Println("------CheckToken------")
	var result Result
	token := GetToken(r)

	ok := client.CheckTokenClient(token)
	if ok {
		result.Result = true
		result.Info = "OK"
	} else if token == "" {
		result.Result = false
	} else {
		result.Result = false
		result.Info = "Auth NG"
	}

	RespondWithJSON(w, http.StatusOK, result)
}

func GenerateToken(w http.ResponseWriter, r *http.Request) {
	log.Println("------GenerateToken------")
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		RespondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	var auth UserAuth
	if err := json.Unmarshal(body, &auth); err != nil {
		RespondWithError(w, http.StatusBadRequest, "JSON Unmarshaling failed .")
		return
	}
	var result AuthResult

	token, ok := client.GenerateTokenClient(auth.User, auth.Pass)
	if ok {
		result.Result = true
		result.Token = token
	} else {
		result.Result = false
	}

	RespondWithJSON(w, http.StatusOK, result)

}

func DeleteToken(w http.ResponseWriter, r *http.Request) {
	log.Println("------DeleteToken------")
	var result Result

	token := GetToken(r)

	ok := client.DeleteTokenClient(token)
	if ok {
		result.Result = true
		result.Info = "OK"
	} else if token == "" {
		result.Result = false
	} else {
		result.Result = false
		result.Info = "Auth NG"
	}

	RespondWithJSON(w, http.StatusOK, result)
}

func GetAllToken(w http.ResponseWriter, r *http.Request) {
	log.Println("------GetAllToken------")
	var result Result

	token := GetToken(r)

	ok := client.DeleteTokenClient(token)
	if ok {
		result.Result = true
		result.Info = "OK"
	} else if token == "" {
		result.Result = false
	} else {
		result.Result = false
		result.Info = "Auth NG"
	}

	RespondWithJSON(w, http.StatusOK, result)
}
