package parkinglot

import (
	parkingspot "Parkinglot/ParkingSpot"
	"Parkinglot/vehicle"
	"time"
)
type ParkingTicket struct{
	Id string
	Veh vehicle.Vehicle
	Ps parkingspot.Spot
	Entrytime time.Time
	Exittime *time.Time	
}

func CreateParkingTicket (id string,v vehicle.Vehicle,ps parkingspot.Spot) *ParkingTicket{
	return &ParkingTicket{
       Id: id,
	   Veh: v,
	   Ps: ps,
	   Entrytime: time.Now(),
	   Exittime: nil,
	}
}