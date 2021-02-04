package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/disharjayanth/go-ds-alg/chap-02/07-CRUDForms/02-forms/customer"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob(`templates/*`))
}

func Home(w http.ResponseWriter, r *http.Request) {
	customers := customer.GetAllCustomers()
	// w.Header().Set("Content-Type", "text/html")
	tpl.ExecuteTemplate(w, "Home", customers)
}

func Create(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "Create", nil)
}

func SubmitCustomerValues(w http.ResponseWriter, r *http.Request) {
	name := r.FormValue("customername")
	ssn := r.FormValue("customerssn")
	customer.CreateCustomer(customer.Customer{
		CustomerName: name,
		CustomerSSN:  ssn,
	})

	customers := customer.GetAllCustomers()
	tpl.ExecuteTemplate(w, "Home", customers)
}

func UpdateForm(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("Cannot convert id string to id int: %w", err)
		fmt.Println(err)
	}

	customer := customer.GetCustomerById(idInt)
	tpl.ExecuteTemplate(w, "Update", customer)
}

func Alter(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("customerid")
	name := r.FormValue("customername")
	ssn := r.FormValue("customerssn")

	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("Cannot convert id string to id int: %w", err)
		fmt.Println(err)
	}

	customer.UpdateCustomer(customer.Customer{
		CustomerID:   idInt,
		CustomerName: name,
		CustomerSSN:  ssn,
	})

	customers := customer.GetAllCustomers()
	tpl.ExecuteTemplate(w, "Home", customers)
}

func ViewCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("Cannot convert id string to id int: %w", err)
		fmt.Println(err)
	}

	customer := customer.GetCustomerById(idInt)

	tpl.ExecuteTemplate(w, "Viewing", customer)
}

func DeleteCustomer(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		err = fmt.Errorf("Cannot convert id string to id int: %w", err)
		fmt.Println(err)
	}

	customer.DeleteCustomer(idInt)

	customers := customer.GetAllCustomers()
	tpl.ExecuteTemplate(w, "Home", customers)
}

func main() {
	log.Println("Listening @PORT:3000")

	http.HandleFunc("/", Home)
	http.HandleFunc("/create", Create)
	http.HandleFunc("/submit", SubmitCustomerValues)
	http.HandleFunc("/update", UpdateForm)
	http.HandleFunc("/alter", Alter)
	http.HandleFunc("/delete", DeleteCustomer)
	http.HandleFunc("/view", ViewCustomer)

	http.ListenAndServe(":3000", nil)
}
