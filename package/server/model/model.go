package model

type FlightInformation struct {
	FlightNo      int
	Source        string
	Destination   string
	DepartureHour int
	DepartureMin  int
	AirFare       float64
	MaxCnt        int
	CurrentCnt    int
}

//type QueryFlightRequest struct {
//	MessageType
//}
