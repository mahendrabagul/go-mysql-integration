package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type Employee struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func main() {
	fmt.Println("Employee Management")

	// Open up our database connection.
	// I've set up a database on my local machine using phpmyadmin.
	// The database is called testDb
	db, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/test")

	// if there is an error opening the connection, handle it
	if err != nil {
		panic(err.Error())
	}

	// defer the close till after the main function has finished
	// executing
	defer db.Close()

	// perform a db.Query insert
	insert, err := db.Query("INSERT INTO employee VALUES ( 2, 'Divit', 'Pune')")

	// if there is an error inserting, handle it
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("Data inserted successfully.")
	// be careful deferring Queries if you are using transactions
	defer insert.Close()
	// Execute the query
	results, err := db.Query("SELECT id, name, city FROM employee")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	fmt.Println("Data fetched successfully.")
	defer results.Close()
	for results.Next() {
		var employee Employee
		// for each row, scan the result into our employee composite object
		err = results.Scan(&employee.ID, &employee.Name, &employee.City)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		// and then print out the employee's Name attribute
		log.Printf(employee.Name)
	}

}
