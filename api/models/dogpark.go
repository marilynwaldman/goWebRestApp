package models

type (
	// User represents the structure of our resource
	Dogpark struct {
		Name   string `json:"name"`
		Address string `json:"address"`
		Id     string `json:"id"`
	}
)