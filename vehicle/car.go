package vehicle

type car struct{
	*vehicle
}

func NewCar(lno string)*car{
	return &car{
		vehicle: NewVehicle(lno,CAR),
	}
}

