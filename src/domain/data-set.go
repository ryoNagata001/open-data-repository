package domain

import (
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/infrastructure"
)

type DataSet struct {
	ID 					bson.ObjectId 	`bson:"_id" json:"_id"`
	Title				string			`bson: "title" json:"title"`
	Publisher 			string			`bson: "publisher"`
	ContactPoint		string			`bson: "contactpoint"` 		// データの誤りを報告する連絡先
	Creator 			string			`bson: "creator"`				// userのIDとひもづける
	Tags 				string			`bson: "tags"`					// カンマ区切りのモノをsplitする
	ReleaseDate 		string 			`bson: "releasedate"` 			// YYYY-MM-DD
	FrequencyOfUpdate	string			`bson: "frequencyofupdate"`	// dataの更新頻度
	LandingPage			string 			`bson: "landingpage"`			// URL
	Spatial 			string			`bson: "spatial"`				// データセットが対象としている都道府県名
	DataResources		[]DataResource	`bson: "dataresources"`		// データリソースの中身
}

var dataset_cl = infrastructure.SetCollection(infrastructure.DataSets.String())

// insert
func InsertNewDataSet(dataSet DataSet) (dbErr error) {
	dbErr = dataset_cl.Insert(dataSet)
	return
}

// idで取得
func GetDataSetById(id bson.ObjectId) (dataSet DataSet, err error) {
	err = dataset_cl.FindId(id).One(&dataSet)
	return
}

// idで取得
func GetDataSetAll() (dataSet []DataSet, err error) {
	err = dataset_cl.Find(nil).All(&dataSet)
	return
}

func GetMyDataSet(publicKey string) (dataSet []DataSet, err error) {
	err = dataset_cl.Find(bson.M{"creator": publicKey}).All(&dataSet)
	return
}

func SearchDataSet(title string, publisher string, tags string, spatial string) (dataSet []DataSet, err error) {
	err = dataset_cl.Find(bson.M{
		"title": bson.RegEx{title, ""},
		"publisher": bson.RegEx{ publisher, ""},
		"tags": bson.RegEx{tags, ""},
		"spatial": bson.RegEx{spatial, ""},
	}).All(&dataSet)
	return
}

func GetDataSetList(page int, perPage int) (dataSet []DataSet, err error) {
	err = dataset_cl.Find(nil).Skip((page-1) * perPage).Limit(perPage).All(&dataSet)
	return
}

func GetCollectionCount() (count int, err error) {
	count, err = dataset_cl.Count()
	return
}

func DeleteDataSet(id bson.ObjectId) (err error) {
	err = dataset_cl.RemoveId(id)
	return
}