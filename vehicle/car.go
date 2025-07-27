package vehicle

type car struct{
	*Vehicle
}

func NewCar(lno string)*car{
	return &car{
		Vehicle: NewVehicle(lno,CAR),
	}
}

