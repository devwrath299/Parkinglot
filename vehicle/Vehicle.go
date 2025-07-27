package vehicle

type vehicle struct{
	LicenseNumber string
	Vt VehicleType
}

func NewVehicle(lno string,vt VehicleType) *vehicle{
	return &vehicle{
		LicenseNumber: lno,Vt: vt,
	}
}






