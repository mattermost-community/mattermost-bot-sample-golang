package db

// Remove is used to remove all the documents matching a query
func Remove(collection string, query interface{}) error {
	_, err := db.Collection(collection).DeleteMany(bg(), query)
	return err
}
