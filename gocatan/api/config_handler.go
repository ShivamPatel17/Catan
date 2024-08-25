package api

import (
	"encoding/json"
	"gocatan/config"
	"net/http"
)

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := config.NewConfig()

	resp, _ := json.Marshal(config)
	w.Write(resp)
}
