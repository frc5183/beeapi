package models

import "beeapi/response"

type Model interface {
	//todo implement this without causing a circular dependency
	//// Create creates the model in the database.
	//Create() error
	//
	//// Update updates the model in the database.
	//Update() error
	//
	//// Delete deletes the model from the database.
	//Delete() error

	// Verify checks if the model is valid.
	Verify() *response.Error

	// ResponseCopy creates a copy of the object which has modified fields which are ready to be put into a response.
	ResponseCopy() (interface{}, error)
}
