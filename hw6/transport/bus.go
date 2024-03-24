package transport

import (
	"fmt"
	"hw6/primitives"
)

type Bus struct {
	name  string
	start primitives.Point
	stop  primitives.Point
}

func NewBus(name string, start primitives.Point, stop primitives.Point) *Bus {
	return &Bus{name, start, stop}
}

func (b *Bus) InPassengers() string {
	return fmt.Sprintf("Сісти на автобус %s на зупинці %s в місті %s\n", b.name, b.start.Address, b.start.City)
}

func (b *Bus) OutPassengers() string {
	return fmt.Sprintf("Вийти з автобусу %s на зупинці %s в місті %s\n", b.name, b.stop.Address, b.stop.City)
}
