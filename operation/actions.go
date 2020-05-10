package operation

import (
	"fmt"
)

//To initialize the slots
func InitializeSlots(size int) {
	slots = make([]*parkedCar, size)
}

//To park incoming car
func ParkCar(car CarInfo) {
	index := availablePosition(slots)
	if index > -1 {
		parkCar := &parkedCar{
			CarInfo: car,
			SlotNo:  index + 1,
		}

		slots[index] = parkCar

		fmt.Printf("Allocated slot number: %d\n", index+1)
	}
}

//To check empty position
func availablePosition(Slots []*parkedCar) int {
	for ind, val := range Slots {
		if val == nil {
			return ind
		}
	}

	fmt.Println("Sorry, parking lot is full")

	return -1
}

//To remove the car from given position
func VacatePosition(position int) {
	if position < 1 || position > len(slots) {
		fmt.Println("No such slot exits")
		return
	}

	if slots[position-1] == nil {
		fmt.Println("No car is present in this slot")
		return
	}

	slots[position-1] = nil
	fmt.Printf("Slot number %d is free\n", position)
}

//to get the status of occupied slots
func GetStatus() {
	fmt.Println("Slot No.\t\tRegistration No\t\tColour")

	for _, val := range slots {
		if val != nil {
			fmt.Println(val.SlotNo, "\t\t\t", val.RegistrationNos, "\t\t", val.Color)
		}
	}
}

//to get the registration number of the vehicle by color
func GetRegistrationNosByColor(color string) {
	result := make([]string, 0)

	for _, val := range slots {
		if val.Color == color {
			result = append(result, val.RegistrationNos)
		}
	}

	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			fmt.Println(result[i])
		} else {
			fmt.Printf("%s, ", result[i])
		}
	}
}

// to get the slot nos of the vehicles of a particulat color
func GetSlotNosByColor(color string) {
	result := make([]int, 0)
	for _, val := range slots {
		if val.Color == color {
			result = append(result, val.SlotNo)
		}
	}

	for i := 0; i < len(result); i++ {
		if i == len(result)-1 {
			fmt.Println(result[i])
		} else {
			fmt.Printf("%d, ", result[i])
		}
	}
}

// to get the slot no of the vehicle with a particular registration number
func GetSlotNosByRegistrationNos(registrationNos string) {
	found := false
	for _, val := range slots {
		if val.RegistrationNos == registrationNos {
			found = true
			fmt.Print(val.SlotNo)
			break
		}
	}

	if !found {
		fmt.Print("Not found")
	}
	fmt.Println()
}
