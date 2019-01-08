package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
	"gopkg.in/mgo.v2/bson"
)

type DataResource struct {
	ID 					bson.ObjectId 	`bson:"_id" json:"_id"`
	Title				string			`bson: "title"`
	URL 				string 			`bson: "url"` 					// url
	Description 		string			`bson: "description"`
	Format 				string			`bson: "format"`				// 拡張子を設定する(ValueはBASE64で保持する)
	Value 				string			`bson: "value"`					// dataの中身をbase64にする
	FileSize			int 			`bson: "file_size"`				// byte
	LastModifiedDate 	date.Date 		`bson: "last_modified_date"`	// リソースの掲載日を YYYY-MM-DDののフォーマットで記入する
	License				string			`bson: "licence"`				// 択一選択
	Copyright			string			`bson: "copyright"`				// 固定文字列
	Language 			string			`bson: "language"`				// 言語(択一選択)
}

func AddDataResource (id bson.ObjectId, dataResource DataResource) (errDb error){
	selector := bson.M{"_id": id}
	update := bson.M{"$set":bson.M{"data_resources": dataResource}}
	errDb = dataset_cl.Update(selector, update)
	return
}