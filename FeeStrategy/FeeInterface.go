package feestrategy

import (
	parkinglot "Parkinglot"
)
type Feestrategy interface{
	Calculate(pt *parkinglot.ParkingTicket)float64
}