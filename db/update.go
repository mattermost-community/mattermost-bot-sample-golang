package db

import (
	"github.com/mattermost/mattermost-bot-sample-golang/types"

	"go.mongodb.org/mongo-driver/mongo/options"
)

// Update function is used to update a document with a given query
func Update(collection string, query types.JSON, in, out interface{}) error {

	after := options.After
	ret := db.Collection(collection).FindOneAndUpdate(bg(), query, in, &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
	})
	if ret.Err() != nil {
		return ret.Err()
	}

	return ret.Decode(out)
}

// UpdateAll function is used to update all the documents with a given query
func UpdateAll(collection string, query types.JSON, in interface{}) error {
	_, err := db.Collection(collection).UpdateMany(bg(), query, in)
	return err
}
