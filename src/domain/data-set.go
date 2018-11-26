package domain

import "google.golang.org/genproto/googleapis/type/date"

type DataSource struct {
	Title				string			`bson: "title"`
	Publisher 			string			`bson: "publisher"`
	ContactPoint		string			`bson: "contact_point"` 		// データの誤りを報告する連絡先
	Creator 			string			`bson: "creator"`				// userのIDとひもづける
	Tags 				[]string		`bson: "tags"`					// カンマ区切りのモノをsplitする
	ReleaseDate 		date.Date 		`bson: "release_date"` 			// YYYY-MM-DD
	FrequencyOfUpdate	string			`bson: "frequency_of_update"`	// dataの更新頻度
	LandingPage			string 			`bson: "landing_page"`			// URL
	Spatial 			string			`bson: "spatial"`				// データセットが対象としている都道府県名
}
