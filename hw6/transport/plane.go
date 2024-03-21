package transport

import (
	"fmt"
	"hw6/primitives"
)

type Plane struct {
}

func (p *Plane) InPassengers(transport string, airport primitives.Point) string {
	return fmt.Sprintf("Сісти на літак %s в аеропорті %s %s\n", transport, airport.City, airport.Country)
}

func (p *Plane) OutPassengers(transport string, airport primitives.Point) string {
	return fmt.Sprintf("Зійти з літака %s в аеропорті %s %s\n", transport, airport.City, airport.Country)
}
