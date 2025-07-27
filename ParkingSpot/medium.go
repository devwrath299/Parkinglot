package parkingspot

import "Parkinglot/vehicle"

type MediumParkingSpot struct{
	ParkingSpot
}

func NewMediumParkingSpot(id string) *MediumParkingSpot{
	return &MediumParkingSpot{
		ParkingSpot: *NewparkingSpot(id),
	}
}

func (ps *MediumParkingSpot) Canfit(v vehicle.Vehicle) bool{
	return v.Vt==vehicle.CAR
}



