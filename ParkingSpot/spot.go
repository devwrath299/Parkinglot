package parkingspot

import "Parkinglot/vehicle"

type Spot interface{
	Park(vh vehicle.Vehicle)bool
	Unpark()
	IsSpotAvailable()bool
	Canfit(v vehicle.Vehicle) bool
}