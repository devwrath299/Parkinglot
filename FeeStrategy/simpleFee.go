package feestrategy

import parkinglot "Parkinglot"

type SimpleFee struct{}

func(s *SimpleFee) Calculate(pt *parkinglot.ParkingTicket)float64{
	const rate=10.0
	timesinside:=pt.Exittime.Sub(pt.Entrytime)
	finaltime:= int64(timesinside.Hours())+1
	return float64(finaltime)*rate
}