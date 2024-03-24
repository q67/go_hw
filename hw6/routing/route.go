package routing

import (
	"fmt"
)

type PublicTransport interface {
	InPassengers() string
	OutPassengers() string
}

type Route struct {
	transports []PublicTransport
}

func NewRoute() *Route {
	return &Route{}
}

func (r *Route) AddTransport(pt PublicTransport) {
	r.transports = append(r.transports, pt)
}

func (r *Route) ShowRoute() {
	for _, transport := range r.transports {
		fmt.Print(transport.InPassengers())
		fmt.Print(transport.OutPassengers())
	}
}
