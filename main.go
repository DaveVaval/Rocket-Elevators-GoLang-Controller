package main

import (
	"fmt"
	"math"
	"sort"
)

var floorRequestButtonID int = 1
var columnID int = 1

// Battery struct
type Battery struct { // this is the base blueprint of a Battery
	ID                      int
	status                  string
	amountOfFloors          int
	amountOfColumns         int
	amountOfBasements       int
	columnsList             []Column
	floorRequestButtonsList []FloorRequestButton
}

func newBattery(id int, status string, amountOfFloors int, amountOfColumns int, amountOfBasements int, amountOfElevatorPerColumn int) Battery { // This will acts as the constructor for the Battery
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

	return b
}

func (b *Battery) createBasmentColumn(amountOfBasements int, amountOfElevatorPerColumn int) { // This will create the column for the basements
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

func (b *Battery) createColumns(amountOfColumns int, amountOfFloors int, amountOfElevatorPerColumn int) { // this will create the columns with thier floors
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

func (b *Battery) findBestColumn(requestedFloor int) Column { // This will return the best column
	col := Column{}
	for _, c := range b.columnsList {
		found := find(requestedFloor, c.servedFloors)
		if found == true {
			col = c
		}
	}
	return col
}

func find(a int, list []int) bool { // function created to check if a list contains a given number
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func (b *Battery) assignElevator(requestedFloor int, direction string) { // This function will return the best elevator from the best column to the user
	column := b.findBestColumn(requestedFloor)
	elevator := column.findBestElevator(1, direction)
	elevator.floorRequestList = append(elevator.floorRequestList, requestedFloor)
	elevator.sortFloorList()
	elevator.move()
	elevator.openDoors()
	fmt.Println("............")
	fmt.Println("Elevator: ", elevator.ID, "from column: ", column.ID, "is sent to lobby")
	fmt.Println("He enters the elevator")
	fmt.Println("............")
	fmt.Println("Elevator reached floor: ", elevator.currentFloor)
	fmt.Println("He gets out...")
}

// Column struct
type Column struct { // // this is the base blueprint of a Column
	ID                int
	status            string
	amountOfFloors    int
	amountOfElevators int
	isBasement        bool
	elevatorsList     []Elevator
	callButtonsList   []CallButton
	servedFloors      []int
}

func newColumn(id int, status string, amountOfFloors int, amountOfElevators int, isBasement bool, servedFloors []int) Column { // This functions acts as a constructor for the Column
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
func (c *Column) requestElevator(userPosition int, direction string) { // This is the function that will be called when a user wants to go back to the lobby from any given floor
	elevator := c.findBestElevator(userPosition, direction)
	elevator.floorRequestList = append(elevator.floorRequestList, 1)
	elevator.sortFloorList()
	elevator.move()
	elevator.openDoors()
	fmt.Println("............")
	fmt.Println("Elevator: ", elevator.ID, " from column: ", c.ID, "is sent to floor: ", userPosition)
	fmt.Println("He enters the elevator")
	fmt.Println("............")
	fmt.Println("Elevator reached floor: ", elevator.currentFloor)
	fmt.Println("He gets out...")
}

func (c *Column) findBestElevator(requestedFloor int, requestedDirection string) Elevator { // This function in conjuction wwith checkElevator will return the best elevator
	bestElevatorInfo := map[string]interface{}{
		"bestElevator": nil,
		"bestScore":    6,
		"referenceGap": math.Inf(1),
	}
	if requestedFloor == 1 {
		for _, e := range c.elevatorsList {
			if 1 == e.currentFloor && e.status == "stopped" {
				bestElevatorInfo = c.checkElevator(1, e, requestedFloor, bestElevatorInfo)
			} else if 1 == e.currentFloor && e.status == "idle" {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if 1 > e.currentFloor && e.direction == "up" {
				bestElevatorInfo = c.checkElevator(3, e, requestedFloor, bestElevatorInfo)
			} else if 1 < e.currentFloor && e.direction == "down" {
				bestElevatorInfo = c.checkElevator(3, e, requestedFloor, bestElevatorInfo)
			} else if e.status == "idle" {
				bestElevatorInfo = c.checkElevator(4, e, requestedFloor, bestElevatorInfo)
			} else {
				bestElevatorInfo = c.checkElevator(5, e, requestedFloor, bestElevatorInfo)
			}
		}
	} else {
		for _, e := range c.elevatorsList {
			if requestedFloor == e.currentFloor && e.status == "idle" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(1, e, requestedFloor, bestElevatorInfo)
			} else if requestedFloor > e.currentFloor && e.direction == "up" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if requestedFloor < e.currentFloor && e.direction == "down" && requestedDirection == e.direction {
				bestElevatorInfo = c.checkElevator(2, e, requestedFloor, bestElevatorInfo)
			} else if e.status == "stopped" {
				bestElevatorInfo = c.checkElevator(4, e, requestedFloor, bestElevatorInfo)
			} else {
				bestElevatorInfo = c.checkElevator(5, e, requestedFloor, bestElevatorInfo)
			}
		}
	}
	return bestElevatorInfo["bestElevator"].(Elevator)
}

func (c *Column) checkElevator(baseScore int, elevator Elevator, floor int, bestElevatorInfo map[string]interface{}) map[string]interface{} {
	if baseScore < bestElevatorInfo["bestScore"].(int) {
		bestElevatorInfo["bestScore"] = baseScore
		bestElevatorInfo["bestElevator"] = elevator
		bestElevatorInfo["referenceGap"] = Abs(elevator.currentFloor - floor)
	}
	return bestElevatorInfo
}

// Abs ...
func Abs(x int) int { // Function created to return the absolute value of an int
	if x < 0 {
		return -x
	}
	return x
}

// Elevator struct
type Elevator struct { // this is the base blueprint of an elevator
	ID               int
	status           string
	amountOfFloors   int
	direction        string
	currentFloor     int
	door             Door
	floorRequestList []int
}

func newElevator(id int, status string, amountOfFloors int, currentFloor int) Elevator { // this function acts as a constructor for Elevator
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

func (e *Elevator) move() { // This function will make the elevator to any given floor
	// i := 0
	for len(e.floorRequestList) != 0 {
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

// RemoveIndex ...
func RemoveIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...) // Function created to remove the first index of a list
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

// Scenario1 ...
func Scenario1() {
	// B1
	battery.columnsList[1].elevatorsList[0].currentFloor = 20
	battery.columnsList[1].elevatorsList[0].direction = "down"
	battery.columnsList[1].elevatorsList[0].floorRequestList = append(battery.columnsList[1].elevatorsList[0].floorRequestList, 5)

	// B2
	battery.columnsList[1].elevatorsList[1].currentFloor = 3
	battery.columnsList[1].elevatorsList[1].direction = "up"
	battery.columnsList[1].elevatorsList[1].floorRequestList = append(battery.columnsList[1].elevatorsList[1].floorRequestList, 15)

	// B3
	battery.columnsList[1].elevatorsList[2].currentFloor = 13
	battery.columnsList[1].elevatorsList[2].direction = "down"
	battery.columnsList[1].elevatorsList[2].floorRequestList = append(battery.columnsList[1].elevatorsList[2].floorRequestList, 1)

	// B4
	battery.columnsList[1].elevatorsList[3].currentFloor = 15
	battery.columnsList[1].elevatorsList[3].direction = "down"
	battery.columnsList[1].elevatorsList[3].floorRequestList = append(battery.columnsList[1].elevatorsList[3].floorRequestList, 2)

	// B5
	battery.columnsList[1].elevatorsList[4].currentFloor = 6
	battery.columnsList[1].elevatorsList[4].direction = "down"
	battery.columnsList[1].elevatorsList[4].floorRequestList = append(battery.columnsList[1].elevatorsList[4].floorRequestList, 1)
	battery.columnsList[1].elevatorsList[4].move()

	fmt.Println("User is at the lobby and wants to go to floor 20")
	fmt.Println("He enters 20 on the pannel")
	battery.assignElevator(20, "up")
}

// Scenario2 ...
func Scenario2() {
	// C1
	battery.columnsList[2].elevatorsList[0].currentFloor = 1
	battery.columnsList[2].elevatorsList[0].direction = "up"
	battery.columnsList[2].elevatorsList[0].floorRequestList = append(battery.columnsList[2].elevatorsList[0].floorRequestList, 21)

	// C2
	battery.columnsList[2].elevatorsList[1].currentFloor = 23
	battery.columnsList[2].elevatorsList[1].direction = "up"
	battery.columnsList[2].elevatorsList[1].floorRequestList = append(battery.columnsList[2].elevatorsList[1].floorRequestList, 28)

	// C3
	battery.columnsList[2].elevatorsList[2].currentFloor = 33
	battery.columnsList[2].elevatorsList[2].direction = "down"
	battery.columnsList[2].elevatorsList[2].floorRequestList = append(battery.columnsList[2].elevatorsList[2].floorRequestList, 1)

	// C4
	battery.columnsList[2].elevatorsList[3].currentFloor = 40
	battery.columnsList[2].elevatorsList[3].direction = "down"
	battery.columnsList[2].elevatorsList[3].floorRequestList = append(battery.columnsList[2].elevatorsList[3].floorRequestList, 24)

	// C5
	battery.columnsList[2].elevatorsList[4].currentFloor = 39
	battery.columnsList[2].elevatorsList[4].direction = "down"
	battery.columnsList[2].elevatorsList[4].floorRequestList = append(battery.columnsList[2].elevatorsList[4].floorRequestList, 1)

	// User at lobby want's to go to floor 36, Elevator 1 should be sent
	fmt.Println("User is at the lobby and wants to go to floor 36")
	fmt.Println("He enters 36 on the pannel")
	battery.assignElevator(36, "up")
}

// Scenario3 ...
func Scenario3() {
	// D1
	battery.columnsList[3].elevatorsList[0].currentFloor = 58
	battery.columnsList[3].elevatorsList[0].direction = "down"
	battery.columnsList[3].elevatorsList[0].floorRequestList = append(battery.columnsList[3].elevatorsList[0].floorRequestList, 1)

	// D2
	battery.columnsList[3].elevatorsList[1].currentFloor = 50
	battery.columnsList[3].elevatorsList[1].direction = "up"
	battery.columnsList[3].elevatorsList[1].floorRequestList = append(battery.columnsList[3].elevatorsList[1].floorRequestList, 60)

	// D3
	battery.columnsList[3].elevatorsList[2].currentFloor = 46
	battery.columnsList[3].elevatorsList[2].direction = "up"
	battery.columnsList[3].elevatorsList[2].floorRequestList = append(battery.columnsList[3].elevatorsList[2].floorRequestList, 58)

	// D4
	battery.columnsList[3].elevatorsList[3].currentFloor = 1
	battery.columnsList[3].elevatorsList[3].direction = "up"
	battery.columnsList[3].elevatorsList[3].floorRequestList = append(battery.columnsList[3].elevatorsList[3].floorRequestList, 54)

	// D5
	battery.columnsList[3].elevatorsList[4].currentFloor = 60
	battery.columnsList[3].elevatorsList[4].direction = "down"
	battery.columnsList[3].elevatorsList[4].floorRequestList = append(battery.columnsList[3].elevatorsList[4].floorRequestList, 1)

	// User at floor 54 want's to go to floor 1, Elevator 1 should be sent
	fmt.Println("User is at floor 54 and wants to go to the lobby")
	fmt.Println("He presses on the pannel")
	battery.columnsList[3].requestElevator(54, "down")
}

// Scenario4 ...
func Scenario4() {
	// A1
	battery.columnsList[0].elevatorsList[0].currentFloor = -4

	// A2
	battery.columnsList[0].elevatorsList[1].currentFloor = 1

	//A3
	battery.columnsList[0].elevatorsList[2].currentFloor = -3
	battery.columnsList[0].elevatorsList[2].direction = "down"
	battery.columnsList[0].elevatorsList[2].floorRequestList = append(battery.columnsList[0].elevatorsList[2].floorRequestList, -5)

	// A4
	battery.columnsList[0].elevatorsList[3].currentFloor = -6
	battery.columnsList[0].elevatorsList[3].direction = "up"
	battery.columnsList[0].elevatorsList[3].floorRequestList = append(battery.columnsList[0].elevatorsList[3].floorRequestList, 1)

	// A5
	battery.columnsList[0].elevatorsList[4].currentFloor = -1
	battery.columnsList[0].elevatorsList[4].direction = "down"
	battery.columnsList[0].elevatorsList[4].floorRequestList = append(battery.columnsList[0].elevatorsList[4].floorRequestList, -6)

	// User at Basement 3 want's to go to floor 1, Elevator 4 should be sent
	fmt.Println("User is at SS3 and wants to go to the lobby")
	fmt.Println("He presses on the pannel")
	battery.columnsList[0].requestElevator(-3, "up")
}

var battery = newBattery(1, "online", 60, 4, 6, 5)

func main() {
	// fmt.Println("-------------------------------// Rocket Elevators //----------------------------------")
	// Scenario1()
	// Scenario2()
	// Scenario3()
	// Scenario4()
}
