package parkingspot

import "Parkinglot/vehicle"

type SmallParkingSpot struct{
	ParkingSpot
}

func NewSmallParkingSpot(id string) *SmallParkingSpot{
	return &SmallParkingSpot{
		ParkingSpot: *NewparkingSpot(id),
	}
}

func (ps *SmallParkingSpot) Canfit(v vehicle.Vehicle) bool{
	return v.Vt==vehicle.BIKE
}



