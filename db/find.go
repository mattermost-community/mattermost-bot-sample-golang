package db

// FindOne function is used to fetch data from the database for the passed query
func FindOne(collection string, query interface{}, x interface{}) error {
	return db.Collection(collection).FindOne(bg(), query).Decode(x)
}

// FindAll function is used to fetch data from the database for the passed query
// and store it in an array of elements
func FindAll(collection string, query interface{}, fn func() interface{}) error {
	i, err := db.Collection(collection).Find(bg(), query)
	if err != nil {
		return err
	}
	defer i.Close(bg())

	for i.Next(bg()) {
		if err := i.Decode(fn()); err != nil {
			return err
		}
	}

	return i.Err()
}
