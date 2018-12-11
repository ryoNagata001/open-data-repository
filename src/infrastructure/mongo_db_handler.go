package infrastructure

import (
	"gopkg.in/mgo.v2"
	"strconv"
	"time"
)

var db *mgo.Database

type Collection int

const (
	Users Collection = iota
	DataSets
	DataResources
)

// Enum的なやつ
func (collection  Collection) String() string {
	switch collection {
	case Users:
		return "users"
	case DataSets:
		return "datasets"
	default:
		return "done."
	}
}

// 初期設定
// TODO これらのリテラルは設定ファイルにまとめるべき
func init() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	db = session.DB("open-data-repository")
	clearDb(db)
}

// 各DAOがCollectionをSetするためのfunction
func SetCollection(Collection string) *mgo.Collection {
	return db.C(Collection)
}

// 再起動時に呼び出すmethod
func clearDb(dbCopy *mgo.Database) {
	var iota Collection = 0
	for {
		if iota.String() == "done." {
			break
		}
		dbCopy.C(iota.String()).RemoveAll(nil)
		iota++
	}
}

// データの総数を返す
func FindTotalDocuments() int64 {
	var iota Collection = 0
	var sum int64

	for {
		if iota.String() == "done." {
			break
		}
		count, _ := db.C(iota.String()).Find(nil).Count()
		sum += int64(count)
		iota++
	}
	return sum
}

// ObjectIdを時間に変換する
func FindTimeFromObjectID(id string) time.Time {
	ts, _ := strconv.ParseInt(id[0:8], 16, 64)
	return time.Unix(ts, 0)
}