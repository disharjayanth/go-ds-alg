package customer

import (
	"database/sql"
	"fmt"

	"github.com/disharjayanth/go-ds-alg/chap-02/secret"
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB
var err error

func init() {
	db, err = sql.Open("mysql", secret.DbName+":"+secret.DbPassword+"@/goalg")
	if err != nil {
		err = fmt.Errorf("Cannot connect to database: %w", err)
		fmt.Println(err)
	}
}

type Customer struct {
	CustomerID   int
	CustomerName string
	CustomerSSN  string
}

func GetAllCustomers() []Customer {
	rows, err := db.Query("SELECT * FROM customer")
	defer rows.Close()
	if err != nil {
		err = fmt.Errorf("Error while querying for customers in customer table: %w", err)
		fmt.Println(err)
	}

	customer := Customer{}
	customers := []Customer{}
	for rows.Next() {
		err = rows.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerSSN)
		if err != nil {
			err = fmt.Errorf("Error while scanning for customer rows: %w", err)
			fmt.Println(err)
		}

		customers = append(customers, customer)
	}

	return customers
}

func GetCustomerById(id int) Customer {
	row := db.QueryRow("SELECT * FROM customer WHERE customerId = ?", id)
	customer := Customer{}
	err := row.Scan(&customer.CustomerID, &customer.CustomerName, &customer.CustomerSSN)
	if err != nil {
		err = fmt.Errorf("Error while scanning for GetCustomerById single customer with id: %w", err)
	}
	return customer
}

func CreateCustomer(customer Customer) {
	_, err := db.Exec("INSERT INTO customer (customerName, SSN) VALUES (?, ?)", customer.CustomerName, customer.CustomerSSN)
	if err != nil {
		err = fmt.Errorf("Error while creating new customer: %w", err)
		fmt.Println(err)
	}
}

func UpdateCustomer(customer Customer) {
	stmt, err := db.Prepare("UPDATE customer SET customerName = ?, SSN = ? WHERE customerId = ?")
	if err != nil {
		err = fmt.Errorf("Error while preparing update stmt: %w", err)
		fmt.Println(err)
	}

	_, err = stmt.Exec(customer.CustomerName, customer.CustomerSSN, customer.CustomerID)
	if err != nil {
		err = fmt.Errorf("Error while executing update prepare stmt: %w", err)
		fmt.Println(err)
	}
}

func DeleteCustomer(id int) {
	_, err := db.Exec("DELETE FROM customer WHERE customerId = ?", id)
	if err != nil {
		err = fmt.Errorf("Error while deleting customer: %w", err)
		fmt.Println(err)
	}
}
