package main

func main() {
	mongoDAL, _ := NewMongoDAL("localhost", "database")

	Bar(mongoDAL)
}

// Bar ...
func Bar(data DataAccessLayer) Thing {
	doc := Thing{}

	data.Insert("foo", doc)
	return doc
}

// Thing ...
type Thing struct {
	id string
}
