package db

// Insert is used to insert a passed record in the DB
func Insert(collection string, in interface{}) error {
	_, err := db.Collection(collection).InsertOne(bg(), in)
	return err
}

// InsertMany is used to insert multiple passed records in the DB
func InsertMany(collection string, in []interface{}) error {
	_, err := db.Collection(collection).InsertMany(bg(), in)
	return err
}
