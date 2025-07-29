package feestrategy

import (
	parkinglot "Parkinglot"
	"Parkinglot/vehicle"
)

type VehicleBasedFee struct{}
func (v * VehicleBasedFee) Calculate(pt *parkinglot.ParkingTicket)float64{
	mp:=make(map[vehicle.VehicleType]float64)
	mp[vehicle.BIKE]=10.0
	mp[vehicle.CAR]=15.0
	mp[vehicle.TRUCK]=20.0
	var vt vehicle.VehicleType=pt.Veh.Vt
	time:=pt.Exittime.Sub(pt.Entrytime)
	totaltime:=int64(time.Hours())+1
	return float64(totaltime)*mp[vt]
}

