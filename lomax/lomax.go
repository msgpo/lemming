package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const USER string = "simar"
const PASSWORD string = "password"

var (
	deptNo    string
	deptName  string
	empNo     int
	fromDate  string
	toDate    string
	birthDate string
	firstName string
	lastName  string
	gender    string
	hireDate  string
	salary    int
	title     string
)

var operationPtr = flag.String("operator", "", "Operation: SELECT, DELETE, UPDATE, INSERT")
var countPtr = flag.Int("count", 1, "Repeat: Number of times to repeat the benchmark.")
var dbPtr = flag.String("db", "", "Database: Name of the DB to perform operations on.")
var tablePtr = flag.String("table", "", "Table: Name of the table to perform operations on.")
var conditionPtr = flag.String("condition", "", "Condition: Constraint on the transaction.")

func validateInput() {
	if *tablePtr == "" {
		log.Fatal("Please specify a MySQL table using the --table option.")
	} else if *dbPtr == "" {
		log.Fatal("Please specify a MySQL database using the --database option.")
	} else if *operationPtr == "" {
		log.Fatal("Please specify a MySQL operation using the --operator option.")
	}
}

func initializeDB() *sql.DB {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", USER, PASSWORD, *dbPtr))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
func prepareStatement(db *sql.DB, operationPtr string, tablePtr string, conditionPtr string) *sql.Rows {
	stmtOut, err := db.Prepare(fmt.Sprintf("%s FROM %s %s", operationPtr, tablePtr, conditionPtr))
	if err != nil {
		log.Fatal(err)
	}
	rows, err := stmtOut.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer stmtOut.Close()
	return rows
}

func processData(rows *sql.Rows) {
	for rows.Next() {
		err := rows.Scan(&deptNo, &deptName)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(deptNo, deptName)
	}
	err := rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	validateInput()

	db := initializeDB()
	defer db.Close()

	rows := prepareStatement(db, *operationPtr, *tablePtr, *conditionPtr)
	defer rows.Close()

	processData(rows)
}