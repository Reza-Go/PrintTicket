package main

import "fmt"

type BusTicket struct {
	Id                int
	DepartureCity     string
	ArrivalCity       string
	DepartureTime     string
	BusType           string
	PassengerName     string
	DepartureTerminal string
	ArrivalTerminal   string
	NationalCode      string
	Price             int
}

type FlightTicket struct {
	Id               int
	DepartureAirport string
	ArrivalAirport   string
	DepartureTime    string
	ArrivalTime      string
	AirplaneType     string
	PassengerName    string
	PassportId       string
	PassengerType    string
	Price            int
}

func main() {

	busTicket := BusTicket{
		Id:                1,
		DepartureCity:     "Mashhad",
		ArrivalCity:       "Tehran",
		DepartureTime:     "8:00",
		BusType:           "Normal",
		PassengerName:     "Jack",
		DepartureTerminal: "1",
		ArrivalTerminal:   "2",
		NationalCode:      "8888",
		Price:             800,
	}
	flightTicket := FlightTicket{
		Id:               2,
		DepartureAirport: "Mashhad",
		ArrivalAirport:   "Tehran",
		DepartureTime:    "8:00",
		ArrivalTime:      "23:00",
		AirplaneType:     "economy",
		PassengerName:    "Jack",
		PassportId:       "1",
		PassengerType:    "VIP",
		Price:            2500,
	}

	printer := []TicketPrinter{busTicket, flightTicket}

	for _, printer := range printer {
		printer.PrintTicket()
	}

}

// ---implement polymorphism with interface
type TicketPrinter interface {
	PrintTicket()
}

func (ticket BusTicket) PrintTicket() {
	fmt.Printf("Bus Ticket:\n ID : %d\n DepartureCity : %s\n ArrivalCity : %s\n DepartureTime : %s\n ", ticket.Id, ticket.DepartureCity, ticket.ArrivalCity, ticket.DepartureTime)
	fmt.Printf("Bus Type: %s\n Passenger Name: %s\n Departure Terminal: %s\n Arrival Terminal: %s\n ", ticket.BusType, ticket.PassengerName, ticket.DepartureTerminal, ticket.ArrivalTerminal)
	fmt.Printf("National Code: %s\n Price: %d\n", ticket.NationalCode, ticket.Price)
}

func (ticket FlightTicket) PrintTicket() {
	fmt.Printf("Flight Ticket:\n ID : %d\n DepartureAirport : %s\n ArrivalAirport : %s\n DepartureTime : %s\n ", ticket.Id, ticket.DepartureAirport, ticket.ArrivalAirport, ticket.DepartureTime)
	fmt.Printf("Arrival Time :%s\n Airplane Type: %s\n Passenger Name: %s\n Passport ID: %s\n ", ticket.ArrivalTime, ticket.AirplaneType, ticket.PassengerName, ticket.PassportId)
	fmt.Printf("Passenger Type: %s\n Price: %d\n", ticket.PassengerType, ticket.Price)
}
