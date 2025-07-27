package vehicle

type Vehicle struct{
	LicenseNumber string
	Vt VehicleType
}

func NewVehicle(lno string,vt VehicleType) *Vehicle{
	return &Vehicle{
		LicenseNumber: lno,Vt: vt,
	}
}






