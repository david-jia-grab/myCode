package main

// DataAccessLayer defines what methods we need from the database
type DataAccessLayer interface {
	Insert(collectionName string, docs ...interface{}) error
}
