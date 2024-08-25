package hexagon

import "gocatan/config"

type HexagonEngine struct {
	HexSideSize    float64
	HexTotalWidth  float64
	HexTotalHeight float64
}

func NewHexagonEngine(cfg config.Config) HexagonEngine {
	return HexagonEngine{
		HexSideSize:    cfg.HexHeight / 2,
		HexTotalWidth:  cfg.HexWidth,
		HexTotalHeight: cfg.HexHeight,
	}
}
