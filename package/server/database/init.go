package database

import (
	"database/sql"
	"encoding/csv"
	"fmt"
	"github.com/YangXingHao96/DistributedSystem/package/server/model"
	"github.com/lib/pq"
	"os"
	"strconv"
)

var DbConnection *sql.DB

func Init() *sql.DB {
	const (
		host     = "localhost"
		port     = 5432
		user     = "exampleuser"
		password = "examplepassword"
		dbname   = "exampledb"
	)

	// Create the connection string
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	// Open a database connection
	DbConnection, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	// Test the connection
	err = DbConnection.Ping()
	if err != nil {
		panic(err)
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	flightSlice, err := prepTableData(wd + "/setup/init_flight.csv")
	if err != nil {
		panic(err)
	}
	err = initTables(DbConnection, flightSlice)
	if err != nil {
		panic(err)
	}
	return DbConnection
}

func isTableExistsError(err error) bool {
	pqError, ok := err.(*pq.Error)
	if !ok {
		return false
	}
	return pqError.Code.Name() == "duplicate_table"
}

func initTables(db *sql.DB, flightSlice []*model.FlightInformation) error {
	createTable := "CREATE TABLE Flight(flight_number INTEGER PRIMARY KEY,\n  source VARCHAR(256),\n    destination VARCHAR(256),\n    flight_time BIGINT,\n  air_fare FLOAT,\n    max_seat_cnt INTEGER,\n    current_seat_cnt INTEGER)"
	fmt.Printf("Executing creations: %s\n", createTable)
	_, err := db.Exec(createTable)
	if err != nil {
		if isTableExistsError(err) {
			// Handle error if table already exists
			fmt.Printf("table flight already exists\n")
		} else {
			return err
		}
	}

	for _, flight := range flightSlice {
		insertFlight := "INSERT INTO Flight (flight_number, source, destination, flight_time, air_fare, max_seat_cnt, current_seat_cnt) VALUES ($1,$2,$3,$4,$5,$6,$7)"
		_, err := db.Exec(insertFlight, flight.FlightNo, flight.Source, flight.Destination, flight.FlightTime, flight.AirFare, flight.MaxCnt, flight.CurrentCnt)
		if err != nil {
			pqErr, ok := err.(*pq.Error)
			if !ok {
				// The error is not a pq.Error, handle it as appropriate
				return err
			}
			// Check for unique key violation error
			if pqErr.Code == "23505" {
				fmt.Printf("unique key violation: %s\n", pqErr.Message)
			} else {
				return err
			}
		}
	}
	return nil
}

func prepTableData(filePath string) ([]*model.FlightInformation, error) {
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}
	var flightSlice []*model.FlightInformation
	for i, record := range records {
		if i == 0 {
			continue
		}
		flightNo, _ := strconv.Atoi(record[0])
		source := record[1]
		destionation := record[2]
		flightTime, _ := strconv.Atoi(record[3])
		airFare, _ := strconv.ParseFloat(record[4], 32)
		airFare32bits := float32(airFare)
		maxCnt, _ := strconv.Atoi(record[5])
		curCnt, _ := strconv.Atoi(record[6])
		flightInfo := &model.FlightInformation{
			FlightNo:    flightNo,
			Source:      source,
			Destination: destionation,
			FlightTime:  flightTime,
			AirFare:     airFare32bits,
			MaxCnt:      maxCnt,
			CurrentCnt:  curCnt,
		}
		flightSlice = append(flightSlice, flightInfo)
	}
	return flightSlice, nil
}
