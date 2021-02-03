package main

import (
	"database/sql"
	"fmt"

	"github.com/disharjayanth/go-ds-alg/chap-02/secret"
	_ "github.com/go-sql-driver/mysql"
)

type Customer struct {
	CustomerId   int
	CustomerName string
	SSN          string
}

var db *sql.DB
var err error

func init() {
	dbstr := fmt.Sprintf("%s:%s@/goalg", secret.DbName, secret.DbPassword)
	db, err = sql.Open("mysql", dbstr)
	if err != nil {
		panic(err)
	}
}

func GetCustomers() []Customer {
	rows, err := db.Query("SELECT * FROM customer ORDER BY customerId DESC")
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("Error while retriveing rows from customer table:%w", err)
		fmt.Println(err)
		return []Customer{}
	}

	customer := Customer{}
	customers := []Customer{}

	for rows.Next() {
		var customerId int
		var customerName string
		var ssn string
		err := rows.Scan(&customerId, &customerName, &ssn)
		if err != nil {
			fmt.Errorf("Error while scanning results from row:%w", err)
			fmt.Println(err)
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	return customers
}

func InsertCustomer(customer Customer) {
	stmt, err := db.Prepare("INSERT INTO customer (customerName, SSN) VALUES (?, ?)")
	defer stmt.Close()
	if err != nil {
		err = fmt.Errorf("Error while preparing insert statment for customer:%w", err)
		fmt.Println(err)
		return
	}

	_, err = stmt.Exec(customer.CustomerName, customer.SSN)
	if err != nil {
		err = fmt.Errorf("Error while executing prepared stmt for insertion:%w", err)
		return
	}
}

func main() {
	customers := GetCustomers()
	fmt.Println(customers)

	fmt.Println("Before insert:", customers)

	newCustomer := Customer{
		CustomerName: "James",
		SSN:          "73104",
	}
	InsertCustomer(newCustomer)

	customers = GetCustomers()
	fmt.Println("After insert:", customers)
}
