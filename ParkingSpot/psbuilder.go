package parkingspot

func CreateSpot(id string, psType ParkingSpotType) Spot{
	switch (psType){
	case SMALL :
		return NewSmallParkingSpot(id)
	case MEDIUM :
		 return NewMediumParkingSpot(id)
	case LARGE :
		return NewLargeParkingSpot(id)
	default :
		panic("No Parking spot of this type") 	
	}
}