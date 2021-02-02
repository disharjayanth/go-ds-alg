package main

import "fmt"

type DataTransferObjectFactory struct {
	pool map[string]DataTransferObject
}

func (factory DataTransferObjectFactory) getDataTransferObject(dtoType string) DataTransferObject {
	dto := factory.pool[dtoType]
	if dto == nil {
		fmt.Println("New DTO of dtoType:", dtoType)
		switch dtoType {
		case "customer":
			factory.pool[dtoType] = Customer{id: "1"}
		case "employee":
			factory.pool[dtoType] = Employee{id: "2"}
		case "manager":
			factory.pool[dtoType] = Manager{id: "3"}
		case "address":
			factory.pool[dtoType] = Address{id: "4"}
		}
		dto = factory.pool[dtoType]
	}
	return dto
}

type DataTransferObject interface {
	getId() string
}

// Customer type
type Customer struct {
	id   string
	name string
	ssn  string
}

func (c Customer) getId() string {
	return c.id
}

// Employee type
type Employee struct {
	id   string
	name string
}

func (e Employee) getId() string {
	return e.id
}

// Manager type
type Manager struct {
	id   string
	name string
	dept string
}

func (m Manager) getId() string {
	return m.id
}

// Address type
type Address struct {
	id          string
	streetLine1 string
	streetLine2 string
	state       string
	city        string
}

func (a Address) getId() string {
	return a.id
}

func main() {
	factory := DataTransferObjectFactory{
		pool: map[string]DataTransferObject{},
	}

	customer := factory.getDataTransferObject("customer")
	fmt.Println("Customer:", customer.getId())

	employee := factory.getDataTransferObject("employee")
	fmt.Println("Employee:", employee.getId())

	manager := factory.getDataTransferObject("manager")
	fmt.Println("Manager:", manager.getId())

	address := factory.getDataTransferObject("address")
	fmt.Println("Address:", address.getId())

	fmt.Println("Factory:")
	fmt.Println(factory)
}
