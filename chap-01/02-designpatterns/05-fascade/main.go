package main

import "fmt"

// Account type
type Account struct {
	id          string
	accountType string
}

func (account *Account) create(accountType string) *Account {
	fmt.Println("Creating new account with given type:", accountType)
	account.accountType = accountType
	return account
}

func (account *Account) getById(id string) *Account {
	fmt.Println("Getting account by id:", id)
	return account
}

func (account *Account) deleteById(id string) {
	fmt.Println("Deleting account given by id:", id)
	return
}

// Customer type
type Customer struct {
	name string
	id   int
}

func (customer *Customer) create(name string) *Customer {
	fmt.Println("Creating new customer with name:", name)
	customer.name = name
	return customer
}

// Transaction type
type Transaction struct {
	id            string
	amount        float32
	srcAccountId  string
	destAccountId string
}

func (transaction *Transaction) create(srcAccountId string, destAccountId string, amount float32) *Transaction {
	fmt.Println("Creating new transaction from src:", srcAccountId, "to dest:", destAccountId, "with amount:", amount)
	transaction.srcAccountId = srcAccountId
	transaction.destAccountId = destAccountId
	transaction.amount = amount
	return transaction
}

type BranchManagerFascade struct {
	account     *Account
	customer    *Customer
	transaction *Transaction
}

func NewBranchManagerFascade() *BranchManagerFascade {
	return &BranchManagerFascade{
		account:     &Account{},
		customer:    &Customer{},
		transaction: &Transaction{},
	}
}

func (fascade *BranchManagerFascade) createCustomerAccount(customerName string, accountType string) (*Customer, *Account) {
	customer := fascade.customer.create(customerName)
	account := fascade.account.create(accountType)
	return customer, account
}

func (fascade *BranchManagerFascade) createTransaction(srcAccountId, destAccountId string, amount float32) *Transaction {
	transaction := fascade.transaction.create(srcAccountId, destAccountId, amount)
	return transaction
}

func main() {
	fascade := NewBranchManagerFascade()

	customer, account := fascade.createCustomerAccount("Thomas Smith", "Savings")
	fmt.Println("Created customer name:", customer.name)
	fmt.Println("Created account type:", account.accountType)

	transaction := fascade.createTransaction("1", "2", 100.011)
	fmt.Println("Created transaction amount:", transaction.amount)
}
