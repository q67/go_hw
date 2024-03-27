package forecast

import (
	"fmt"
	"math/rand"
)

type WRandom struct {
}

func (r *WRandom) GetTomorrowWeather() string {
	randTemp := rand.Intn(10)
	return fmt.Sprintf("%dC\n", randTemp)
}
