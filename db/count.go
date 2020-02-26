package db

// Count is used to count the number of documents for the passed query
func Count(collection string, query interface{}) (int64, error) {
	return db.Collection(collection).CountDocuments(bg(), query)
}
