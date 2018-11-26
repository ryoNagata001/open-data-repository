package domain

import (
	"google.golang.org/genproto/googleapis/type/date"
)

type DataResource struct {
	Title				string			`bson: "title"`
	URL 				string 			`bson: "url"` 					// url
	Description 		string			`bson: "description"`
	Format 				string			`bson: "format"`				// 拡張子を設定する(ValueはBASE64で保持する)
	Value 				[]byte			`bson: "value"`					// dataの中身
	FileSize			int 			`bson: "file_size"`				// byte
	LastModifiedDate 	date.Date 		`bson: "last_modified_date"`	// リソースの掲載日を YYYY-MM-DDののフォーマットで記入する
	License				string			`bson: "licence"`				// 択一選択
	Copyright			string			`bson: "copyright"`				// 固定文字列
	Language 			string			`bson: "language"`				// 言語(択一選択)
}
