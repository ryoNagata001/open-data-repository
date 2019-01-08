package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/infrastructure"
)

type DataSet struct {
	ID 					bson.ObjectId 	`bson:"_id" json:"_id"`
	Title				string			`bson: "title"`
	Publisher 			string			`bson: "publisher"`
	ContactPoint		string			`bson: "contact_point"` 		// データの誤りを報告する連絡先
	Creator 			string			`bson: "creator"`				// userのIDとひもづける
	Tags 				[]string		`bson: "tags"`					// カンマ区切りのモノをsplitする
	ReleaseDate 		date.Date 		`bson: "release_date"` 			// YYYY-MM-DD
	FrequencyOfUpdate	string			`bson: "frequency_of_update"`	// dataの更新頻度
	LandingPage			string 			`bson: "landing_page"`			// URL
	Spatial 			string			`bson: "spatial"`				// データセットが対象としている都道府県名
	DataResources		[]DataResource	`bson: "data_resources"`		// データリソースの中身
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