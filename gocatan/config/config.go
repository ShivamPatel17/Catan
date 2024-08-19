package config

import (
	"encoding/json"
	"net/http"
)

const (
	HexagonImageScale  = 0.3
	HexagonImageHeight = 508.0 * HexagonImageScale
	HexagonImageWidth  = 440.0 * HexagonImageScale
)

type Config struct {
	HexWidth  float32
	HexHeight float32
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {

	config := Config{
		HexWidth:  HexagonImageWidth,
		HexHeight: HexagonImageHeight,
	}
	resp, _ := json.Marshal(config)
	w.Write(resp)
}
