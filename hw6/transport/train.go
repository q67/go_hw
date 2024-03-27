package transport

import (
	"fmt"
	"hw6/primitives"
)

type Train struct {
	name  string
	start primitives.Point
	stop  primitives.Point
}

func NewTrain(name string, start primitives.Point, stop primitives.Point) *Train {
	return &Train{name, start, stop}
}

func (t *Train) InPassengers() string {
	return fmt.Sprintf("Сісти на поїзд %s на станції %s %s\n", t.name, t.start.City, t.start.Country)
}

func (t *Train) OutPassengers() string {
	return fmt.Sprintf("Зійти з поїзду %s на станції %s %s\n", t.name, t.start.City, t.start.Country)
}
