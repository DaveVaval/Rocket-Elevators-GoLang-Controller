package main

import (
	"fmt"
)

// Battery struct
type Battery struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfColumns   int
	amountOfBasements int
	columnsList       []Column
	// floorRequestButtonsList []FloorRequestButton
}

func New(id int, status string, amountOfFloors int, amountOfColumns int, amountOfBasements int) *Battery {
	return &Battery{
		ID:                id,
		status:            status,
		amountOfFloors:    amountOfFloors,
		amountOfColumns:   amountOfColumns,
		amountOfBasements: amountOfBasements,
		columnsList:       []Column{},
	}
}

// Column struct
type Column struct {
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	isBasement        bool
	// elevatorsList     []Elevator
	// callButtonsList   []CallButton
	// servedFloors []int
}

// Elevator struct
type Elevator struct {
	ID               int
	status           string
	amountOfFloors   int
	direction        string
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
	fmt.Println("-------------------------------// TESTING //----------------------------------")
	testBat := New(1, "online", 60, 4, 6)
	testCol := Column{1, "online", 66, 5, true}
	testBat.columnsList = append(testBat.columnsList, testCol)

	fmt.Println("Test battery: ", testBat.status)
	fmt.Println("Test column: ", testBat.columnsList[0])
}
