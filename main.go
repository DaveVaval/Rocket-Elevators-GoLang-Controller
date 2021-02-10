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

func (b Battery) batteryInit(id int, status string, amountOfFloors int, amountOfColumns int, amountOfBasements int, amountOfElevatorPerColumn int) *Battery {
	b.ID = id
	b.status = status
	b.amountOfFloors = amountOfFloors
	b.amountOfColumns = amountOfColumns
	b.amountOfBasements = amountOfBasements
	b.columnsList = []Column{}
	b.floorRequestButtonsList = []FloorRequestButton{}

	if amountOfBasements > 0 {
		b.createBasmentColumn(b.amountOfBasements, amountOfElevatorPerColumn)
	}

	// for debug
	for column := range b.columnsList {
		fmt.Println("Column: ", column)
		for _, floor := range b.columnsList[0].servedFloors {
			fmt.Println("	Floor: ", floor)
		}
	}

	return &Battery{
		ID:                      id,
		status:                  status,
		amountOfFloors:          amountOfFloors,
		amountOfColumns:         amountOfColumns,
		amountOfBasements:       amountOfBasements,
		columnsList:             b.columnsList,
		floorRequestButtonsList: []FloorRequestButton{},
	}
}

func (b *Battery) createBasmentColumn(amountOfBasements int, amountOfElevatorPerColumn int) {
	servedFloors := []int{}
	floor := -1
	columnID := 1

	for i := 0; i < amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	b.columnsList = append(b.columnsList, Column{columnID, "online", amountOfBasements, amountOfElevatorPerColumn, true, []Elevator{}, []CallButton{}, servedFloors})
	columnID++

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

func (c Column) columnInit(id int, status string, amountOfFloors int, amountOfElevators int, isBasement bool, servedFloors []int) *Column {
	c.ID = id
	c.status = status
	c.amountOfFloors = amountOfFloors
	c.amountOfElevators = amountOfElevators
	c.isBasement = isBasement
	c.servedFloors = servedFloors
	c.elevatorsList = []Elevator{}
	c.callButtonsList = []CallButton{}

	for floor := range c.servedFloors {
		fmt.Println("Floor: ", floor)
	}

	return &Column{
		ID:                id,
		status:            status,
		amountOfFloors:    amountOfFloors,
		amountOfElevators: amountOfElevators,
		isBasement:        isBasement,
		elevatorsList:     []Elevator{},
		callButtonsList:   []CallButton{},
		servedFloors:      c.servedFloors,
	}
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
	// testBat := Battery{1, "online", 66, 4, 6, []Column{}, []FloorRequestButton{}}
	// goodBat := testBat.batteryInit(1, "online", 60, 4, 6, 5)
	// // testCol := Column{1, "online", 66, 5, true}

	// fmt.Println("Test status: ", goodBat.status)
	// fmt.Println("Test id: ", goodBat.ID)
	// fmt.Println("Test column list: ", goodBat.columnsList)
	// fmt.Println("Test column: ", testBat.columnsList[0])
	Battery{1, "online", 60, 4, 6, []Column{}, []FloorRequestButton{}}.batteryInit(1, "online", 60, 4, 6, 5)

}
