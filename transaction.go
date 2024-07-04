package main

import (
	"fmt"
	"time"
)

// Transaction represents a basic transaction in a blockchain.
type Transaction struct {
	ID        string // Unique identifier for the transaction
	Sender    string // Address or identifier of the sender
	Recipient string // Address or identifier of the recipient
	Amount    int    // Amount transferred in the transaction
	Timestamp string // Timestamp when the transaction occurred
}

// NewTransaction creates a new Transaction instance.
func NewTransaction(sender, recipient string, amount int) *Transaction {
	t := time.Now().UTC().Format("2006-01-02 15:04:05")
	return &Transaction{
		ID:        "unique_transaction_id", // You can generate a unique ID here
		Sender:    sender,
		Recipient: recipient,
		Amount:    amount,
		Timestamp: t,
	}
}

func trans() {
	// Example usage
	sender := "sender_address_or_identifier"
	recipient := "recipient_address_or_identifier"
	amount := 10

	// Create a new transaction
	transaction := NewTransaction(sender, recipient, amount)

	// Print transaction details
	fmt.Println("Transaction ID:", transaction.ID)
	fmt.Println("Sender:", transaction.Sender)
	fmt.Println("Recipient:", transaction.Recipient)
	fmt.Println("Amount:", transaction.Amount)
	fmt.Println("Timestamp:", transaction.Timestamp)
}
