package vehicle

type bike struct{
	*vehicle
}

func NewBike(lno string)*bike{
	return &bike{
		vehicle: NewVehicle(lno,BIKE),
	}
}



