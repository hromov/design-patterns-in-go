package main

import (
	"log"

	"github.com/hromov/design-patterns-in-go/creational/factory/transport"
)

type Transport interface {
	Deliver() error
	Route() []string
}

func main() {
	var truck, ship Transport
	var err error

	truck, err = transport.New(transport.TransportTruck)
	if err != nil {
		log.Fatal("Can't' create truck error: ", err)
	}
	truck.Deliver()
	truckRoutes := truck.Route()
	log.Println("truck routes: ", truckRoutes)

	ship, err = transport.New(transport.TransportShip)
	if err != nil {
		log.Fatal("Can't' create ship error: ", err)
	}
	ship.Deliver()
	shipRoutes := ship.Route()
	log.Println("ship routes: ", shipRoutes)
}
