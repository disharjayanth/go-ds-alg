package main

import (
	"database/sql"
	"fmt"

	"github.com/disharjayanth/go-ds-alg/chap-02/secret"
	_ "github.com/go-sql-driver/mysql"
)

// Customer type has fields CustomerId int, CustomerName string, SSN string
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

// GetCustomers returns all customers
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
			err = fmt.Errorf("Error while scanning results from row:%w", err)
			fmt.Println(err)
		}
		customer.CustomerId = customerId
		customer.CustomerName = customerName
		customer.SSN = ssn
		customers = append(customers, customer)
	}

	return customers
}

// GetCustomerById returns single customer given id
func GetCustomerById(id int) Customer {
	customer := Customer{}
	row := db.QueryRow("SELECT * FROM customer WHERE customerId = ?", id)
	err := row.Scan(&customer.CustomerId, &customer.CustomerName, &customer.SSN)
	if err != nil {
		err = fmt.Errorf("Error while scanning single row from customer table given id:%w", err)
		fmt.Println(err)
	}
	return customer
}

// InsertCustomer inserts given customer to db
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

// UpdateCustomer updates given customer
func UpdateCustomer(customer Customer) sql.Result {
	stmt, err := db.Prepare("UPDATE customer SET customerName=?, SSN=? WHERE customerId=?")
	if err != nil {
		err = fmt.Errorf("Error while preparing query: %w", err)
		fmt.Println(err)
	}

	res, err := stmt.Exec(customer.CustomerName, customer.SSN, customer.CustomerId)
	if err != nil {
		err = fmt.Errorf("Error while executing prepared stmt for update query with single id: %w", err)
	}

	return res
}

func DeleteCustomer(id int) sql.Result {
	stmt, err := db.Prepare("DELETE FROM customer WHERE customerId=?")
	if err != nil {
		err = fmt.Errorf("Error while preparing delete statment: %w", err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		err = fmt.Errorf("Error while executing prepared stmt: %w", err)
		fmt.Println(err)
	}

	return res
}

func main() {
	customers := GetCustomers()
	fmt.Println(customers)

	fmt.Println("Before insert:", customers)

	// newCustomer := Customer{
	// 	CustomerName: "James",
	// 	SSN:          "73104",
	// }
	// InsertCustomer(newCustomer)

	customers = GetCustomers()
	fmt.Println("After insert:", customers)

	customer := GetCustomerById(2)
	fmt.Println(customer)

	updateCustomer := Customer{
		CustomerId:   2,
		CustomerName: "Charlie",
		SSN:          "43752391",
	}

	res := UpdateCustomer(updateCustomer)
	fmt.Println(res.RowsAffected())

	res = DeleteCustomer(3)
	fmt.Println(res.RowsAffected())
}
