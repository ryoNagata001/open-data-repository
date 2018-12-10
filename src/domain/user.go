package domain

import (
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/infrastructure"
)

type User struct {
	ID			bson.ObjectId		`bson: "_id" json: "_id"`
	Name 		string				`bson: "name" json: "name"`
	PublicKey	string				`bson: "public_key" json: "public_key"`
}


// mongoのcollection
var cl = infrastructure.SetCollection(infrastructure.Users.String())

func InsertNewUser (user User) {
	dbErr := cl.Insert(user)

	if dbErr != nil {
		panic(dbErr) // TODO panic should not be used.
	}
}

// userのsig確認
func CheckExistenceOfUser(publicKey string) (int) {
	count, _ := cl.Find(bson.M{"publicKey": publicKey}).Count()
	return count
}

func GetUserByPubKey(publicKey string) (user User) {
	cl.Find(bson.M{"publicKey": publicKey}).One(&user)
	return
}

func GetUserById(id bson.ObjectId) (user User, err error) {
	err = cl.FindId(id).One(&user)
	return
}

