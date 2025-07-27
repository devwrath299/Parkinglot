package parkingspot

import "Parkinglot/vehicle"

type LargeParkingSpot struct{
	ParkingSpot
}

func NewLargeParkingSpot(id string) *LargeParkingSpot{
	return &LargeParkingSpot{
		ParkingSpot: *NewparkingSpot(id),
	}
}

func (ps *LargeParkingSpot) Canfit(v vehicle.Vehicle) bool{
	return v.Vt==vehicle.CAR
}



