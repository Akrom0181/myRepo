package main

import (
 "fmt"
 "time"
)

type Airport struct {
	PlaneType  string
	FlightDate time.Time
	FromCity   string
	ToCity     string
	FlightTime string
}

func (a *Airport) SummerFlights(flights []Airport) []Airport {
 	var summerFlights []Airport
 	for _, flight := range flights {
 		month := flight.FlightDate.Month()
 		if month >= time.June && month <= time.August {
   			summerFlights = append(summerFlights, flight)
  		}
 	}
return summerFlights
}

func main() {
	flights := []Airport{
		{
			PlaneType:  "Boeing 747",
			FlightDate: time.Date(2024, time.June, 15, 9, 0, 0, 0, time.UTC),
			FromCity:   "New York",
			ToCity:     "London",
			FlightTime: "6 hours",
		},
		{
			PlaneType:  "Airbus A380",
			FlightDate: time.Date(2024, time.July, 5, 14, 30, 0, 0, time.UTC),
			FromCity:   "Paris",
			ToCity:     "Tokyo",
			FlightTime: "12 hours",
		},
		{
			PlaneType:  "Boeing 737",
			FlightDate: time.Date(2024, time.August, 20, 10, 15, 0, 0, time.UTC),
			FromCity:   "London",
			ToCity:     "Dubai",
			FlightTime: "7 hours",
		},
		{
			PlaneType:  "Boeing 747",
			FlightDate: time.Date(2024, time.January, 29, 9, 30, 0, 0, time.UTC),
			FromCity:   "Chicago",
			ToCity:     "Tashkent",
			FlightTime: "12 hours",
		},
		{
			PlaneType:  "Airbus A380",
			FlightDate: time.Date(2024, time.January, 29, 2, 30, 0, 0, time.UTC),
			FromCity:   "Lisabon",
			ToCity:     "Tashkent",
			FlightTime: "8 hours",
		},
		{
			PlaneType:  "Boeing 737",
			FlightDate: time.Date(2024, time.January, 29, 10, 30, 0, 0, time.UTC),
			FromCity:   "London",
			ToCity:     "Moskva",
			FlightTime: "6 hours",
		},
		{
			PlaneType:  "Boeing 737",
			FlightDate: time.Date(2024, time.January, 30, 11, 30, 0, 0, time.UTC),
			FromCity:   "Oslo",
			ToCity:     "London",
			FlightTime: "8 hours",
		},
		{
			PlaneType:  "Boeing 777",
			FlightDate: time.Date(2024, time.January, 31, 17, 30, 0, 0, time.UTC),
			FromCity:   "Samarkand",
			ToCity:     "Urgench",
			FlightTime: "1 hours",
		},
		{
			PlaneType:  "Boeing 737",
			FlightDate: time.Date(2024, time.January, 30, 4, 30, 0, 0, time.UTC),
			FromCity:   "London",
			ToCity:     "Abu-Dhabi",
			FlightTime: "7 hours",
		},
		{
			PlaneType:  "Boeing 777",
			FlightDate: time.Date(2024, time.January, 31, 4, 35, 0, 0, time.UTC),
			FromCity:   "Samarkand",
			ToCity:     "Abu-Dhabi",
			FlightTime: "4 hours",
		},
	}

var airport Airport
summerFlights := airport.SummerFlights(flights)

fmt.Println("Summer Flights:")
for _, flight := range summerFlights {
  	fmt.Println(flight)
}
}