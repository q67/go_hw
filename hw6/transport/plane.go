package transport

import (
	"fmt"
	"hw6/primitives"
)

type Plane struct {
	name  string
	start primitives.Point
	stop  primitives.Point
}

func NewPlane(name string, start primitives.Point, stop primitives.Point) *Plane {
	return &Plane{name, start, stop}
}

func (p *Plane) InPassengers() string {
	return fmt.Sprintf("Сісти на літак %s в аеропорті %s %s\n", p.name, p.start.City, p.start.Country)
}

func (p *Plane) OutPassengers() string {
	return fmt.Sprintf("Зійти з літака %s в аеропорті %s %s\n", p.name, p.stop.City, p.stop.Country)
}
