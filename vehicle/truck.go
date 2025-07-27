package vehicle

type truck struct{
	*Vehicle
}

func NewTruck(lno string) *truck{
	return &truck{
		Vehicle: NewVehicle(lno,TRUCK),
	}
}





