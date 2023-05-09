package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	_ "github.com/lib/pq"
)

func main() {
	// Open a connection to the database
	db, err := sql.Open("postgres", "postgres://postgres:ibuchi596643@localhost:5432/gotask?sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// Drop the Users table
	_, err = db.Exec("DROP TABLE IF EXISTS users")
	if err != nil {
		panic(err)
	}

	// Create Users table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			user_id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			age NUMERIC NOT NULL,
			phone TEXT
		)
	`)

	if err != nil {
		panic(err)
	}

	// Insert records into the Users table
	_, err = db.Exec("INSERT INTO users (name, age, phone) VALUES ($1, $2, $3), ($4, $5, $6), ($7, $8, $9)",
		"John", "28", "",
		"Tom", "29", "1-800-123-1234",
		"Jenny", "30", "",
	)
	if err != nil {
		panic(err)
	}
	// Create a Response Structure
	type Response struct {
		Status  int                      `json:"status_code"`
		Results []map[string]interface{} `json:"data"`
	}

	// Query the database
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	// Create a slice of maps to store the results
	results := []map[string]interface{}{}

	// Iterate over the results
	for rows.Next() {
		// Create a map to store each row
		row := map[string]interface{}{}

		// Create temporary variables for the scan
		var id int
		var name string
		var age int
		var phone string

		// Scan the row into the temporary variables
		err := rows.Scan(&id, &name, &age, &phone)
		if err != nil {
			panic(err)
		}

		// Assign the temporary variables to the map
		row["user_id"] = id
		row["name"] = name
		row["age"] = age
		row["phone"] = phone

		// Append the map to the results slice
		results = append(results, row)
	}

	// Convert the results slice to JSON
	jsonResult, err := json.Marshal(Response{http.StatusOK, results})
	if err != nil {
		panic(err)
	}

	// Print the JSON result
	fmt.Println(string(jsonResult))

	// Check for errors after iterating over all rows
	if err := rows.Err(); err != nil {
		panic(err)
	}
}

