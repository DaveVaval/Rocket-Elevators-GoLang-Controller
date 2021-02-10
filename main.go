package main

import (
	"fmt"
)

// Battery struct
type Battery struct {
	ID                      int
	status                  string
	amountOfFloors          int
	amountOfColumns         int
	amountOfBasements       int
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

// Column struct
type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []CallButton
	servedFloors      []int
}

// Elevator struct
type Elevator struct {
	ID               int
	status           string
	amountOfFloors   int
	direction        nil
	currentFloor     int
	door             Door
	floorRequestList []int
}

// CallButton struct
type CallButton struct {
	ID        int
	status    string
	floor     int
	direction string
}

// FloorRequestButton struct
type FloorRequestButton struct {
	ID     int
	status string
	floor  int
}

// Door struct
type Door struct {
	ID     int
	status string
}

func main() {
	fmt.Println("yooooooooo")
}
