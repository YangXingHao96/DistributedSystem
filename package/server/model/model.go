package model

type FlightInformation struct {
	FlightNo    int
	Source      string
	Destination string
	FlightTime  int
	AirFare     float32
	MaxCnt      int
	CurrentCnt  int
}

//type QueryFlightRequest struct {
//	MessageType
//}
