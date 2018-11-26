package domain

import "gopkg.in/mgo.v2/bson"

type User struct {
	ID			bson.ObjectId		`bson: "_id" json: "_id"`
	Name 		string				`bson: "name" json: "name"`
	PublicKey	string				`bson: "public_key" json: "public_key"`
}
