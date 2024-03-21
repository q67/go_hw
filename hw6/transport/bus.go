package transport

import (
	"fmt"
	"hw6/primitives"
)

type Bus struct {
}

func (b *Bus) InPassengers(transport string, busstop primitives.Point) string {
	return fmt.Sprintf("Сісти на автобус %s на зупинці %s в місті %s\n", transport, busstop.Address, busstop.City)
}

func (b *Bus) OutPassengers(transport string, busstop primitives.Point) string {
	return fmt.Sprintf("Вийти з автобусу %s на зупинці %s в місті %s\n", transport, busstop.Address, busstop.City)
}
