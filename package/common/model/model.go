package model

type FlightDetail struct {
	FlightNo         int `json:"flightNo"`
	DepartureHour    int `json:"departureHour"`
	DepartureMinute  int `json:"departureMinute"`
	SeatAvailability int `json:"seatAvailability"`
}

type CommonRequest struct {
	RequestType string `json:"requestType"`
	RequestBody struct {
		FlightNo    int    `json:"flightNo,omitempty"`
		Source      string `json:"source,omitempty"`
		Destination string `json:"destination,omitempty"`
		SeatCount   int    `json:"seatCount,omitempty"`
	} `json:"requestBody"`
}

type CommonResponse struct {
	ResponseType string `json:"responseType"`
	ResponseBody struct {
		FlightIds       []string      `json:"flightIDs,omitempty"`
		FlightDetail    *FlightDetail `json:"flightDetail,omitempty"`
		Acknowledgement string        `json:"ack,omitempty"`
		Error           string        `json:"error,omitempty"`
	}
}