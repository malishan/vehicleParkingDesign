package operation

// struct to hold the value of each car
type CarInfo struct {
	RegistrationNos string
	Color           string
}

// struct to hold the car information and the slot nos for each parked vehicle
type parkedCar struct {
	CarInfo
	SlotNo int
}

// slots present in the parking area
var slots []*parkedCar
