package data

import (
	"github.com/yoneyan/vm_mgr/ggate/client"
	"log"
	"net/http"
)

func GetUserVM(w http.ResponseWriter, r *http.Request) {
	log.Println("------GetUserVM------")

	token := GetToken(r)
	data := client.GetUserVMClient(token)

	RespondWithJSON(w, http.StatusOK, data)
}
