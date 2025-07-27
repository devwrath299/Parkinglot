package main

import "fmt"

type vehicleType int

const(
	CAR vehicleType=iota
	TRUCK
	BIKE
)
func (vt vehicleType) String() string{
 return [...]string{"CAR","BIKE","TRUCk"}[vt]
}

func main(){
	vt:=TRUCK
	fmt.Println("vt",vt.String())
}