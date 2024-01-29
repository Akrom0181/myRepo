package main

import (
 "fmt"
 "time"
)

type Aeroport struct {
 PlaneType  string
 FlightDate time.Time
 FromCity   string
 ToCity     string
 FlightTime string
}

func (a *Aeroport) Flights(flights []Aeroport) []Aeroport {
    var filteredFlights []Aeroport
    for _, flight := range flights {
     if flight.FlightDate.Hour() >= 2 && flight.FlightDate.Hour() <= 3 && flight.ToCity == "Tashkent" {
      filteredFlights = append(filteredFlights, flight)
     }
    }
    return filteredFlights
   }
   

func main() {
    flights := []Aeroport{
    {
        PlaneType:  "Boeing 747",
        FlightDate: time.Date(2024, time.January, 29, 1, 30, 0, 0, time.UTC),
        FromCity:   "New York",
        ToCity:     "Tashkent",
        FlightTime: "10 hours",
    },
    {
        PlaneType:  "Airbus A380",
        FlightDate: time.Date(2024, time.January, 29, 2, 30, 0, 0, time.UTC),
        FromCity:   "Paris",
        ToCity:     "Tashkent",
        FlightTime: "8 hours",
    },
    {
        PlaneType:  "Boeing 737",
        FlightDate: time.Date(2024, time.January, 29, 3, 30, 0, 0, time.UTC),
        FromCity:   "London",
        ToCity:     "Dubai",
        FlightTime: "6 hours",
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
    }

    var aeroport Aeroport
    filteredFlights := aeroport.Flights(flights)

    fmt.Println("Filtered Flights:")
    for _, flight := range filteredFlights {
        fmt.Println(flight)
    }
}