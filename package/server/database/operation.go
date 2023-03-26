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
		if err := rows.Scan(&flightInfo.FlightNo, &flightInfo.Source, &flightInfo.Destination, &flightInfo.DepartureHour,
			&flightInfo.DepartureMin, &flightInfo.AirFare, &flightInfo.MaxCnt, &flightInfo.CurrentCnt); err != nil {
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
		if err := rows.Scan(&flightInfo.FlightNo, &flightInfo.Source, &flightInfo.Destination, &flightInfo.DepartureHour,
			&flightInfo.DepartureMin, &flightInfo.AirFare, &flightInfo.MaxCnt, &flightInfo.CurrentCnt); err != nil {
			return nil, err
		}
		flightDetail.FlightNo = flightInfo.FlightNo
		flightDetail.DepartureHour = flightInfo.DepartureHour
		flightDetail.DepartureMinute = flightInfo.DepartureMin
		flightDetail.SeatAvailability = flightInfo.MaxCnt - flightInfo.CurrentCnt
	}

	if flightDetail.FlightNo == 0 {
		return nil, errors.New("flight number does not exist")
	}
	return &flightDetail, nil
}

func MakeReservation(db *sql.DB, flightNo, seatCnt int) error {
	flightDetail, err := GetFlightDetail(db, flightNo)
	if err != nil {
		return err
	}
	if seatCnt > flightDetail.SeatAvailability {
		return errors.New("cannot reserve seats more than available seats")
	}
	reservationStatement := "UPDATE Flight SET current_seat_cnt = current_seat_cnt + $1 WHERE flight_number = $2"
	stmt, err := db.Prepare(reservationStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seatCnt, flightNo)
	if err != nil {
		return err
	}
	return nil
}

func CancelReservation(db *sql.DB, flightNo, seatCnt int) error {
	_, err := GetFlightDetail(db, flightNo)
	if err != nil {
		return err
	}

	cancelStatement := "UPDATE Flight SET current_seat_cnt = current_seat_cnt - $1 WHERE flight_number = $2"
	stmt, err := db.Prepare(cancelStatement)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(seatCnt, flightNo)
	if err != nil {
		return err
	}
	return nil
}

func AddFlight(db *sql.DB, flightNo, departureHour, departureMin, maxCnt, curCnt int, source, destination string, airFare float64) error {
	insertStatement := "INSERT INTO Flight (flight_number, source, destination, departure_hour, departure_min, air_fare, max_seat_cnt, current_seat_cnt) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
	_, err := db.Exec(insertStatement, flightNo, source, destination, departureHour, departureMin, airFare, maxCnt, curCnt)
	if err != nil {
		return err
	}
	return nil
}
