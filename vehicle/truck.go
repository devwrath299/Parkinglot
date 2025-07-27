package vehicle

type truck struct{
	*vehicle
}

func NewTruck(lno string) *truck{
	return &truck{
		vehicle: NewVehicle(lno,TRUCK),
	}
}





