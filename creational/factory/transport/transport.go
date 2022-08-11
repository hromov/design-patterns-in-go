package transport

import (
	"errors"

	"github.com/hromov/design-patterns-in-go/creational/factory/routing"
)

type TransportType int

const (
	TransportTruck TransportType = iota
	TransportShip
)

type RoutingService interface {
	ShipsRoute() []string
	TrucksRoute() []string
}

type transport struct {
	route map[string]bool
}

func New(t TransportType) (transport, error) {
	rs := routing.New()
	transport := transport{}
	switch t {
	case TransportTruck:
		transport.new_route(rs.TrucksRoute())
		return transport, nil
	case TransportShip:
		transport.new_route(rs.ShipsRoute())
		return transport, nil
	default:
		return transport, errors.New("Unknown Transport Type")
	}
}
func (t *transport) new_route(route []string) {
	t.route = map[string]bool{}
	for _, r := range route {
		t.route[r] = false
	}
}

func (t transport) Deliver() error {
	if len(t.route) == 0 {
		return errors.New("Empty delivery list")
	}
	for point := range t.route {
		t.route[point] = true
	}
	return nil
}

func (t transport) Route() []string {
	route := []string{}
	for point := range t.route {
		route = append(route, point)
	}
	return route
}
