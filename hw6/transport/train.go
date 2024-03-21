package transport

import (
	"fmt"
	"hw6/primitives"
)

type Train struct {
}

func (t *Train) InPassengers(transport string, station primitives.Point) string {
	return fmt.Sprintf("Сісти на поїзд %s на станції %s %s\n", transport, station.City, station.Country)
}

func (t *Train) OutPassengers(transport string, station primitives.Point) string {
	return fmt.Sprintf("Зійти з поїзду %s на станції %s %s\n", transport, station.City, station.Country)
}
