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
	columnsList             []string
	floorRequestButtonsList []string
}

// Column struct
type Column struct {
}

// Elevator struct
type Elevator struct {
}

// CallButton struct
type CallButton struct {
}

// FloorRequestButton struct
type FloorRequestButton struct {
}

// Door struct
type Door struct {
}

func main() {
	fmt.Println("yooooooooo")
}
