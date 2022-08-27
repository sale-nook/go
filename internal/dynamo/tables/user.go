package tables

import "time"

type Profile struct {
	ID        string    `json:"id" dynamo:"id,hash"`
	CreatedAt time.Time `json:"createdAt" dynamo:"createdAt,range"`
	CreatedBy string    `json:"createdBy" dynamo:"createdBy"`
	Name      string    `json:"name" dynamo:"name"`
	Email     string    `json:"email" dynamo:"email"`
}

type Transaction struct {
	ID          string    `json:"id" dynamo:"id,hash"`
	AccountID   string    `json:"accountId" dynamo:"accountId,range"`
	Amount      int       `json:"amount" dynamo:"amount"`
	Category    string    `json:"category" dynamo:"category"`
	Date        time.Time `json:"date" dynamo:"date"`
	Description string    `json:"description" dynamo:"description"`
	Name        string    `json:"name" dynamo:"name"`
	PlaidID     string    `dynamo:"plaidId"`
	PlaidType   string    `dynamo:"plaidType"`
}

type User struct {
	ID           string        `json:"id" dynamodbav:"id"`
	Profiles     []Profile     `json:"profiles" dynamodbav:"profiles"`
	Transactions []Transaction `json:"transactions" dynamodbav:"transactions"`

	// Only used internally.
	PlaidAccount string    `dynamodbav:"plaidAccount"`
	PlaidToken   string    `dynamodbav:"plaidToken"`
	CreatedAt    time.Time `dynamodbav:"createdAt"`
	UpdatedAt    time.Time `dynamodbav:"updatedAt"`
}
