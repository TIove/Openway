package models

import "time"

type DbRecord struct {
	Name 		string
	LastName 	string
	Email       string
	Birthday 	time.Time // TODO add other fields
}