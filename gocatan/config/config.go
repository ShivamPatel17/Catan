package config

import (
	"encoding/json"
	"net/http"
)

const (
	HexagonImageScale  = 0.3
	HexagonImageHeight = 508.0 * HexagonImageScale // 508 is the height of the picture from the front end. Eventually, the front end should pass these values to the backend during initialization
	HexagonImageWidth  = 440.0 * HexagonImageScale
	VerticeDiamter     = 30
)

type Config struct {
	HexWidth  float32
	HexHeight float32
}

func NewConfig() Config {
	return Config{
		HexWidth:  HexagonImageWidth,
		HexHeight: HexagonImageHeight,
	}
}

func GetConfigHandler(w http.ResponseWriter, r *http.Request) {
	config := Config{
		HexWidth:  HexagonImageWidth,
		HexHeight: HexagonImageHeight,
	}
	resp, _ := json.Marshal(config)
	w.Write(resp)
}
