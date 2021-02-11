package main

import (
	"fmt"
	"math"
	"sort"
)

var floorRequestButtonID int = 1
var columnID int = 1

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

func newBattery(id int, status string, amountOfFloors int, amountOfColumns int, amountOfBasements int, amountOfElevatorPerColumn int) Battery {
	b := Battery{}
	b.ID = id
	b.status = status
	b.amountOfFloors = amountOfFloors
	b.amountOfColumns = amountOfColumns
	b.amountOfBasements = amountOfBasements
	b.columnsList = []Column{}
	b.floorRequestButtonsList = []FloorRequestButton{}

	if amountOfBasements > 0 {
		b.createBasmentColumn(b.amountOfBasements, amountOfElevatorPerColumn)
		amountOfColumns--
	}
	b.createFloorRequestButtons(amountOfFloors)
	b.createColumns(amountOfColumns, amountOfFloors, amountOfElevatorPerColumn)
	// for debug
	// for column := range b.columnsList {
	// 	fmt.Println("Column: ", column)
	// 	for _, floor := range b.columnsList[0].servedFloors {
	// 		fmt.Println("	Floor: ", floor)
	// 	}
	// }

	return b
}

func (b *Battery) createBasmentColumn(amountOfBasements int, amountOfElevatorPerColumn int) {
	servedFloors := []int{}
	floor := -1

	for i := 0; i < amountOfBasements; i++ {
		servedFloors = append(servedFloors, floor)
		floor--
	}
	column := newColumn(columnID, "online", amountOfBasements, amountOfElevatorPerColumn, true, servedFloors)
	b.columnsList = append(b.columnsList, column)
	columnID++

}

func (b *Battery) createColumns(amountOfColumns int, amountOfFloors int, amountOfElevatorPerColumn int) {
	amountOfFloorsPerColumn := math.Ceil(float64(amountOfFloors / amountOfColumns))
	n := int(amountOfFloorsPerColumn)
	floor := 1

	for i := 0; i < amountOfColumns; i++ {
		servedFloors := []int{}
		for j := 0; j < n; j++ {
			if floor <= amountOfFloors {
				servedFloors = append(servedFloors, floor)
				floor++
			}
		}
		column := newColumn(columnID, "online", amountOfFloors, amountOfElevatorPerColumn, false, servedFloors)
		b.columnsList = append(b.columnsList, column)
		columnID++
	}
}

func (b *Battery) createFloorRequestButtons(amountOfFloors int) {
	buttonFloor := 1
	for i := 0; i < amountOfFloors; i++ {
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, FloorRequestButton{floorRequestButtonID, "off", buttonFloor})
		floorRequestButtonID++
		buttonFloor++
	}
}

func (b *Battery) createBasementFloorRequestButtons(amountOfBasements int) {
	buttonFloor := -1
	for i := 0; i < amountOfBasements; i++ {
		b.floorRequestButtonsList = append(b.floorRequestButtonsList, FloorRequestButton{floorRequestButtonID, "off", buttonFloor})
		buttonFloor--
		floorRequestButtonID++
	}
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

func newColumn(id int, status string, amountOfFloors int, amountOfElevators int, isBasement bool, servedFloors []int) Column {
	c := Column{}
	c.ID = id
	c.status = status
	c.amountOfFloors = amountOfFloors
	c.amountOfElevators = amountOfElevators
	c.isBasement = isBasement
	c.servedFloors = servedFloors
	c.elevatorsList = []Elevator{}
	c.callButtonsList = []CallButton{}
	c.createElevators(amountOfFloors, amountOfElevators)
	c.createCallButtons(amountOfFloors, isBasement)

	// for floor := range c.servedFloors {
	// 	fmt.Println("Floor: ", floor)
	// }

	return c
}

func (c *Column) createElevators(amountOfFloors int, amountOfElevators int) {
	elevatorID := 1
	for i := 0; i < amountOfElevators; i++ {
		elevator := newElevator(elevatorID, "idle", amountOfFloors, 1)
		c.elevatorsList = append(c.elevatorsList, elevator)
		elevatorID++
	}
}

func (c *Column) createCallButtons(amountOfFloors int, isBasement bool) {
	callButtonID := 1
	if isBasement == true {
		buttonFloor := -1
		for i := 0; i < amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, CallButton{callButtonID, "off", buttonFloor, "up"})
			buttonFloor--
			callButtonID++
		}
	} else {
		buttonFloor := 1
		for i := 0; i < amountOfFloors; i++ {
			c.callButtonsList = append(c.callButtonsList, CallButton{callButtonID, "off", buttonFloor, "down"})
			buttonFloor++
			callButtonID++
		}
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

func newElevator(id int, status string, amountOfFloors int, currentFloor int) Elevator {
	e := Elevator{}
	e.ID = id
	e.status = status
	e.amountOfFloors = amountOfFloors
	e.direction = "null"
	e.currentFloor = currentFloor
	e.door = Door{id, "closed"}
	e.floorRequestList = []int{}

	return e
}

func (e *Elevator) move() {
	i := 0
	for i != len(e.floorRequestList) {
		destination := e.floorRequestList[0]
		e.status = "moving"
		if e.currentFloor < destination {
			e.direction = "up"
			for e.currentFloor < destination {
				e.currentFloor++
			}
		} else if e.currentFloor > destination {
			e.direction = "down"
			for e.currentFloor > destination {
				e.currentFloor--
			}
		}
		e.status = "idle"
		e.floorRequestList = RemoveIndex(e.floorRequestList, 0)
	}
}

func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func (e *Elevator) sortFloorList() {
	if e.direction == "up" {
		sort.Slice(e.floorRequestList, func(i, j int) bool { return e.floorRequestList[i] < e.floorRequestList[j] })
	} else {
		sort.Slice(e.floorRequestList, func(i, j int) bool { return e.floorRequestList[i] > e.floorRequestList[j] })
	}
}

func (e *Elevator) openDoors() {
	e.door.status = "open"
	e.door.status = "closed"
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
	// goodBat := testBat.newBattery(1, "online", 60, 4, 6, 5)
	// // testCol := Column{1, "online", 66, 5, true}

	// fmt.Println("Test status: ", goodBat.status)
	// fmt.Println("Test id: ", goodBat.ID)
	// fmt.Println("Test column list: ", goodBat.columnsList)
	// fmt.Println("Test column: ", testBat.columnsList[0])
	// Battery{1, "online", 60, 4, 6, []Column{}, []FloorRequestButton{}}.newBattery(1, "online", 60, 4, 6, 5)
	testBat := newBattery(1, "online", 60, 4, 6, 5)
	fmt.Println(testBat)
}
