package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"

	"project/vehicleParkingDesign/operation"
)

const (
	createParkingLot        = "create_parking_lot"
	parkCar                 = "park"
	leaveParking            = "leave"
	parkingStatus           = "status"
	registrationNosByColor  = "registration_numbers_for_cars_with_colour"
	slotNosByColor          = "slot_numbers_for_cars_with_colour"
	slotNosByRegistrationNo = "slot_number_for_registration_number"
)

func main() {
	var (
		reader   *bufio.Reader
		fileName string
	)

	if len(os.Args) > 1 {
		fileName = os.Args[1]
	}

	if fileName == "" {
		reader = bufio.NewReaderSize(os.Stdin, 1024*1024) //handle input from command prompt
	} else {
		file, err := os.Open(fileName)
		checkError(err)
		reader = bufio.NewReaderSize(file, 1024*1024) // handle input from file input
	}

	handleInput(reader)
}

// handle input
func handleInput(reader *bufio.Reader) {

	var (
		totalSlot int
		err       error
	)

	line := readline(reader)
	action := strings.Split(line, " ")

	if action[0] != createParkingLot {
		return // program terminates if wrong input
	}

	totalSlot, err = strconv.Atoi(action[1])
	checkError(err)

	if totalSlot < 1 {
		return // size of the parking lanes cannot be less than 1
	}

	fmt.Printf("Created a parking lot with %d slots\n", totalSlot)

	operation.InitializeSlots(totalSlot)

	for {
		line := readline(reader)
		if line == "" {
			break
		}

		action := strings.Split(line, " ")

		performAction(action)
	}

	return
}

// perform actions on the slot
func performAction(action []string) {
	switch action[0] {
	case parkCar:
		regisNos := action[1]
		color := action[2]
		car := operation.CarInfo{
			RegistrationNos: regisNos,
			Color:           color,
		}

		operation.ParkCar(car)

	case leaveParking:
		position, err := strconv.Atoi(action[1])
		checkError(err)

		operation.VacatePosition(position)

	case parkingStatus:
		operation.GetStatus()

	case registrationNosByColor:
		color := action[1]
		operation.GetRegistrationNosByColor(color)

	case slotNosByColor:
		color := action[1]
		operation.GetSlotNosByColor(color)

	case slotNosByRegistrationNo:
		registrationNos := action[1]
		operation.GetSlotNosByRegistrationNos(registrationNos)
	default:
		os.Exit(0)
	}
}

// handle error
func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

//read a single of text
func readline(reader *bufio.Reader) string {
	line, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(line), "\r\n")
}
