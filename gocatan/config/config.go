package config

const (
	HexagonImageScale  = 0.3
	HexagonImageHeight = 508.0 * HexagonImageScale // 508 is the height of the picture from the front end. Eventually, the front end should pass these values to the backend during initialization
	HexagonImageWidth  = 440.0 * HexagonImageScale
	VerticeDiamter     = 10
)

type Config struct {
	HexWidth  float64
	HexHeight float64
}

func NewConfig() Config {
	return Config{
		HexWidth:  HexagonImageWidth,
		HexHeight: HexagonImageHeight,
	}
}
