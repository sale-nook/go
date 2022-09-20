package tables

import "time"

type Profile struct {
	ID        string    `json:"id" dynamo:"id,hash"`
	CreatedAt time.Time `json:"createdAt" dynamo:"createdAt,range"`
	CreatedBy string    `json:"createdBy" dynamo:"createdBy"`
	Name      string    `json:"name" dynamo:"name"`
	Email     string    `json:"email" dynamo:"email"`
}

type User struct {
	ID           string        `json:"id" dynamodbav:"id"`
	Profiles     []Profile     `json:"profiles" dynamodbav:"profiles"`

	// Only used internally.
	CreatedAt    time.Time `dynamodbav:"createdAt"`
	UpdatedAt    time.Time `dynamodbav:"updatedAt"`
}
