package data

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type AuthResult struct {
	Result bool   `json:"result"`
	Token  string `json:"token"`
}

type Result struct {
	Result bool   `json:"result"`
	Info   string `json:"info"`
}

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	RespondWithJSON(w, code, map[string]string{"error": msg})
}

func RespondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:4200")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

	w.WriteHeader(code)
	w.Write(response)
}

func GetToken(r *http.Request) string {
	if Token, ok := r.Header["Authorization"]; ok {
		d := strings.Split(Token[0], " ")
		fmt.Println("RestAPI Token: " + d[1])
		return d[1]
	}
	return ""
}
