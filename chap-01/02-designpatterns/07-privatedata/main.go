package main

import (
	"encoding/json"
	"fmt"
)

// AccountDetails type
type AccountDetails struct {
	id          string
	accountType string
}

// Account type
type Account struct {
	details      *AccountDetails
	CustomerName string
}

func (ac *Account) setDetails(id string, accountType string) {
	ac.details = &AccountDetails{
		id:          id,
		accountType: accountType,
	}
}

func (ac *Account) getId() string {
	return ac.details.id
}

func (ac *Account) getAccountType() string {
	return ac.details.accountType
}

func main() {
	account := &Account{
		CustomerName: "John Smith",
	}

	account.setDetails("4532", "Savings")

	jsonSliceByte, err := json.Marshal(account)
	if err != nil {
		fmt.Errorf("Error marshalling to JSON: %w", err)
		return
	}

	fmt.Println("Private hidden field:", string(jsonSliceByte))

	fmt.Println("Account id:", account.getId())
	fmt.Println("Account type:", account.getAccountType())
}
