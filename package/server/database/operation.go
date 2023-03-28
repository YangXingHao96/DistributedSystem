package database

import (
	"database/sql"
	"errors"
	"github.com/YangXingHao96/DistributedSystem/package/common/model"
	dbModel "github.com/YangXingHao96/DistributedSystem/package/server/model"
)

func GetFlights(db *sql.DB, source, destination string) ([]int, error) {
	queryStatement := "SELECT * FROM Flight WHERE source=$1 AND destination=$2"
	query, err := db.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	rows, err := query.Query(source, destination)
	if err != nil {
		return nil, err
	}
	var flightIds []int
	for rows.Next() {
		var flightInfo dbModel.FlightInformation
		if err := rows.Scan(&flightInfo.FlightNo, &flightInfo.Source, &flightInfo.Destination, &flightInfo.FlightTime, &flightInfo.AirFare, &flightInfo.MaxCnt, &flightInfo.CurrentCnt); err != nil {
			return nil, err
		}
		flightIds = append(flightIds, flightInfo.FlightNo)
	}

	return flightIds, nil
}

func GetFlightDetail(db *sql.DB, flightNo int) (*model.FlightDetail, error) {
	queryStatement := "SELECT * FROM Flight WHERE flight_number = $1"
	query, err := db.Prepare(queryStatement)
	if err != nil {
		return nil, err
	}
	defer query.Close()
	rows, err := query.Query(flightNo)
	if err != nil {
		return nil, err
	}

	var flightDetail model.FlightDetail
	for rows.Next() {
		var flightInfo dbModel.FlightInformation
		if err := rows.Scan(&flightInfo.FlightNo, &flightInfo.Source, &flightInfo.Destination, &flightInfo.FlightTime, &flightInfo.AirFare, &flightInfo.MaxCnt, &flightInfo.CurrentCnt); err != nil {
			return nil, err
		}
		flightDetail.FlightNo = flightInfo.FlightNo
		flightDetail.Source = flightInfo.Source
		flightDetail.Destination = flightInfo.Destination
		flightDetail.SeatAvailability = flightInfo.MaxCnt - flightInfo.CurrentCnt
		flightDetail.TotalSeats = flightInfo.MaxCnt
		flightDetail.CurrentSeats = flightInfo.CurrentCnt
		flightDetail.FlightTime = flightInfo.FlightTime
		flightDetail.AirFare = flightInfo.AirFare
	}
	return &flightDetail, nil
}

func MakeReservation(db *sql.DB, flightNo, seatCnt int) (int, error) {
	flightDetail, err := GetFlightDetail(db, flightNo)
	if err != nil {
		return 0, err
	}
	if flightDetail.FlightNo == 0 {
		return 0, errors.New("cannot reserve seats for non existent flight")
	}
	if seatCnt > flightDetail.SeatAvailability {
		return 0, errors.New("cannot reserve seats more than available seats")
	}
	reservationStatement := "UPDATE Flight SET current_seat_cnt = current_seat_cnt + $1 WHERE flight_number = $2"
	stmt, err := db.Prepare(reservationStatement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seatCnt, flightNo)
	if err != nil {
		return 0, err
	}
	newSeatCnt := flightDetail.SeatAvailability - seatCnt
	return newSeatCnt, nil
}

func CancelReservation(db *sql.DB, flightNo, seatCnt int) (int, error) {
	flightDetails, err := GetFlightDetail(db, flightNo)
	if err != nil {
		return 0, err
	}
	if flightDetails.FlightNo == 0 {
		return 0, errors.New("cannot cancel reservation for non existent flight")
	}
	if flightDetails.CurrentSeats < seatCnt {
		return 0, errors.New("current seats reserved are less than number of seats to be cancelled")
	}
	cancelStatement := "UPDATE Flight SET current_seat_cnt = current_seat_cnt - $1 WHERE flight_number = $2"
	stmt, err := db.Prepare(cancelStatement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seatCnt, flightNo)
	if err != nil {
		return 0, err
	}
	newSeatCnt := flightDetails.SeatAvailability + seatCnt
	return newSeatCnt, nil
}

func AddFlight(db *sql.DB, flightNo, flightTime, maxCnt, curCnt int, source, destination string, airFare float32) error {
	insertStatement := "INSERT INTO Flight (flight_number, source, destination, flight_time, air_fare, max_seat_cnt, current_seat_cnt) VALUES ($1,$2,$3,$4,$5,$6,$7)"
	_, err := db.Exec(insertStatement, flightNo, source, destination, flightTime, airFare, maxCnt, curCnt)
	if err != nil {
		return err
	}
	return nil
}
