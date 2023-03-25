package database

import (
	"DistributedSystem/package/server/model"
	"database/sql"
	"encoding/csv"
	"fmt"
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
	flightSlice, err := prepTableData(wd + "/init_flight.csv")
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
	createTable := "CREATE TABLE Flight(flight_number INTEGER PRIMARY KEY,\n  source VARCHAR(256),\n    destination VARCHAR(256),\n    departure_hour INTEGER,\n    departure_min INTEGER,\n    air_fare FLOAT,\n    max_cnt INTEGER,\n    current_cnt INTEGER)"
	fmt.Printf("Executing creations: %s\n", createTable)
	_, err := db.Exec(createTable)
	if err != nil {
		if isTableExistsError(err) {
			// Handle error if table already exists
			fmt.Printf("table flight already exists")
		} else {
			return err
		}
	}

	for _, flight := range flightSlice {
		insertFlight := "INSERT INTO Flight (flight_number, source, destination, departure_hour, departure_min, air_fare, max_cnt, current_cnt) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)"
		_, err := db.Exec(insertFlight, flight.FlightNo, flight.Source, flight.Destination, flight.DepartureHour, flight.DepartureMin, flight.AirFare, flight.MaxCnt, flight.CurrentCnt)
		if err != nil {
			pqErr, ok := err.(*pq.Error)
			if !ok {
				// The error is not a pq.Error, handle it as appropriate
				return err
			}
			// Check for unique key violation error
			if pqErr.Code == "23505" {
				fmt.Printf("unique key violation: %s", pqErr.Message)
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
		departureHour, _ := strconv.Atoi(record[3])
		departureMin, _ := strconv.Atoi(record[4])
		airFare, _ := strconv.ParseFloat(record[5], 64)
		maxCnt, _ := strconv.Atoi(record[6])
		curCnt, _ := strconv.Atoi(record[7])
		flightInfo := &model.FlightInformation{
			FlightNo:      flightNo,
			Source:        source,
			Destination:   destionation,
			DepartureHour: departureHour,
			DepartureMin:  departureMin,
			AirFare:       airFare,
			MaxCnt:        maxCnt,
			CurrentCnt:    curCnt,
		}
		flightSlice = append(flightSlice, flightInfo)
	}
	return flightSlice, nil
}
