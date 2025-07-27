package vehicle

type bike struct{
	*Vehicle
}

func NewBike(lno string)*bike{
	return &bike{
		Vehicle: NewVehicle(lno,BIKE),
	}
}



