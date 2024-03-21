package routing

import (
	"fmt"
	"hw6/primitives"
)

type PublicTransport interface {
	InPassengers(trasport string, point primitives.Point) string
	OutPassengers(trasport string, point primitives.Point) string
}

type Route struct {
	transports map[string]PublicTransport
	route      []string
}

func NewRoute() *Route {
	r := &Route{
		transports: make(map[string]PublicTransport),
		route:      []string{},
	}

	return r
}

func (r *Route) AddTransport(transport string, pt PublicTransport, from primitives.Point, to primitives.Point) {
	r.transports[transport] = pt
	fromLine := r.transports[transport].InPassengers(transport, from)
	toLine := r.transports[transport].OutPassengers(transport, to)
	r.route = append(r.route, fromLine, toLine)
}

func (r Route) ShowTransport() {
	for _, line := range r.route {
		fmt.Print(line)
	}
}
