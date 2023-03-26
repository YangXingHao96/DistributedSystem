package constant

const (
	QueryFlightsReq = iota
	QueryFlightsResp
	QueryFlightDetailReq
	QueryFlightDetailResp
	AddFlightReq
	AddFlightResp
	MakeReservationReq
	MakeReservationResp
	CancelReservationReq
	CancelReservationResp
	GetReservationForFlightReq
	GetReservationForFlightResp
	RegisterForMonitor
	MonitorUpdateResp
	MonitorBackoffResp
	GeneralErrResp
)

const (
	Broadcast   = "broadcast"
	ErrorHeader = "error"
)
