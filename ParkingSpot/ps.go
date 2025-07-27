package parkingspot

import (
	"Parkinglot/vehicle"
	"sync"
)

type ParkingSpot struct{
	SpotId string
	Veh *vehicle.Vehicle
	IsOccupied bool
	mu sync.Mutex
}

func NewparkingSpot(id string ) *ParkingSpot{
    return &ParkingSpot{
		SpotId: id,
		Veh:nil,
		IsOccupied: false,
	}
}

func (ps *ParkingSpot) Park(v vehicle.Vehicle) bool{
     ps.mu.Lock()
	 defer ps.mu.Unlock()
	 if ps.IsOccupied{
		return false
	 }
	 ps.IsOccupied=true
	 ps.Veh=&v
	 return true

}

func (ps * ParkingSpot) Unpark(){
	ps.mu.Lock()
	defer ps.mu.Unlock()
	ps.Veh=nil
	ps.IsOccupied=false
}

func (ps *ParkingSpot) IsSpotAvailable()bool{
	ps.mu.Lock()
	defer ps.mu.Unlock()
	return ps.IsOccupied
}



