package route

import (
	"encoding/base64"
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/domain"
	"open-data-repository/src/open-data-repository-abci/common/code"
	"open-data-repository/src/open-data-repository-abci/common/util"
	"strings"
)

func RouteDeliverTx(body map[string]interface{}, message map[string]interface{}) uint32 {
	code := code.CodeTypeOK

	switch body["type"] {
	case "createUser":
		code = createUser(body, message)
		break
	case "addDataSet":
		code = addDataSet(body, message)
		break
	case "addDataResource":
		code = addDataResource(body, message)
		break
	case "deleteDataSet":
		code = deleteDataSet(body, message)
		break
	case "editDataSet":
		code = editDataSet(body, message)

	}

	return code
}

// dataのinsertを行う
func createUser(body map[string]interface{}, message map[string]interface{}) uint32 {
	entity := body["entity"].(map[string]interface{})

	var user domain.User
	user.ID = bson.ObjectIdHex(entity["id"].(string))
	user.Name = entity["name"].(string)

	pubKeyBytes, errDecode := base64.StdEncoding.DecodeString(message["publicKey"].(string))

	if errDecode != nil {
		return code.CodeTypeBadData
	}

	// public keyを取得
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	user.PublicKey = publicKey

	// dataのinsert
	err := domain.InsertNewUser(user)
	if err != nil {
		return code.CodeTypeBadData
	}
	return code.CodeTypeOK
}

func addDataSet(body map[string]interface{}, message map[string]interface{}) uint32 {
	entity := body["entity"].(map[string]interface{})

	// data set instanceを作る
	var dataSet domain.DataSet
	dataSet.ID = bson.ObjectIdHex(entity["id"].(string))
	dataSet.Title = entity["title"].(string)
	dataSet.Publisher = entity["publisher"].(string)
	dataSet.ContactPoint = entity["contact_point"].(string)
	// public keyを取得
	pubKeyBytes, errDecode := base64.StdEncoding.DecodeString(message["publicKey"].(string))
	if errDecode != nil {
		return code.CodeTypeBadData
	}
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	dataSet.Creator = publicKey

	dataSet.Tags = entity["tags"].(string)
	dataSet.ReleaseDate = entity["release_date"].(string)
	dataSet.FrequencyOfUpdate = entity["frequency_of_update"].(string)
	dataSet.LandingPage = entity["landing_page"].(string)
	dataSet.Spatial = entity["spatial"].(string)

	var DataResoueces = entity["data_resources"].([]interface{})
	for _, value := range DataResoueces {
		temp := value.(map[string]interface{})
		var dataResource domain.DataResource
		dataResource.ID = bson.ObjectIdHex(temp["id"].(string))
		dataResource.Title = temp["title"].(string)
		dataResource.URL = temp["url"].(string)
		dataResource.Description = temp["description"].(string)
		dataResource.Format = temp["format"].(string)
		dataResource.Value = temp["value"].(string)
		dataResource.FileSize = temp["file_size"].(float64)
		dataResource.LastModifiedDate = temp["last_modified_date"].(string)
		dataResource.License = temp["license"].(string)
		dataResource.Copyright = temp["copyright"].(string)
		dataResource.Language = temp["language"].(string)
		dataSet.DataResources = append(dataSet.DataResources, dataResource)
	}

	err := domain.InsertNewDataSet(dataSet)
	if err != nil {
		return code.CodeTypeBadData
	}

	return code.CodeTypeOK
}

func editDataSet(body map[string]interface{}, message map[string]interface{}) uint32 {
	entity := body["entity"].(map[string]interface{})

	// data set instanceを作る
	var dataSet domain.DataSet
	dataSet.ID = bson.ObjectIdHex(entity["id"].(string))

	// data set instanceを取得
	var dataSetTemp, err = domain.GetDataSetById(bson.ObjectIdHex(entity["id"].(string)))

	// データセットが存在しない場合
	if err != nil {
		return code.CodeTypeBadData
	}

	pubKeyBytes, errDecode := base64.StdEncoding.DecodeString(message["publicKey"].(string))

	if errDecode != nil {
		return code.CodeTypeBadData
	}

	// public keyを取得
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))

	// リクエストユーザーに権限がが存在しなかった場合
	if publicKey != dataSetTemp.Creator {
		return code.CodeTypeUnauthorized
	}

	dataSet.Title = entity["title"].(string)
	dataSet.Publisher = entity["publisher"].(string)
	dataSet.ContactPoint = entity["contact_point"].(string)
	dataSet.Creator = publicKey

	dataSet.Tags = entity["tags"].(string)
	dataSet.ReleaseDate = entity["release_date"].(string)
	dataSet.FrequencyOfUpdate = entity["frequency_of_update"].(string)
	dataSet.LandingPage = entity["landing_page"].(string)
	dataSet.Spatial = entity["spatial"].(string)

	var DataResoueces = entity["data_resources"].([]interface{})
	for _, value := range DataResoueces {
		temp := value.(map[string]interface{})
		var dataResource domain.DataResource
		dataResource.ID = bson.ObjectIdHex(temp["id"].(string))
		dataResource.Title = temp["title"].(string)
		dataResource.URL = temp["url"].(string)
		dataResource.Description = temp["description"].(string)
		dataResource.Format = temp["format"].(string)
		dataResource.Value = temp["value"].(string)
		dataResource.FileSize = temp["file_size"].(float64)
		dataResource.LastModifiedDate = temp["last_modified_date"].(string)
		dataResource.License = temp["license"].(string)
		dataResource.Copyright = temp["copyright"].(string)
		dataResource.Language = temp["language"].(string)
		dataSet.DataResources = append(dataSet.DataResources, dataResource)
	}

	err = domain.DeleteDataSet(bson.ObjectIdHex(entity["id"].(string)))
	if err != nil {
		return code.CodeTypeBadData
	}
	err = domain.InsertNewDataSet(dataSet)
	if err != nil {
		return code.CodeTypeBadData
	}

	return code.CodeTypeOK
}

func addDataResource(body map[string]interface{}, message map[string]interface{}) uint32 {
	entity := body["entity"].(map[string]interface{})

	// data set instanceを取得
	var dataSet, err = domain.GetDataSetById(bson.ObjectIdHex(entity["id"].(string)))

	if err != nil {
		return code.CodeTypeBadData
	}

	var DataResoueces = entity["data_resources"].([]map[string]interface{})
	for _, value := range DataResoueces {
		var DataResource domain.DataResource
		DataResource.ID = bson.ObjectIdHex(value["id"].(string))
		DataResource.Title = value["title"].(string)
		DataResource.URL = value["url"].(string)
		DataResource.Description = value["description"].(string)
		DataResource.Format = value["format"].(string)
		DataResource.Value = value["value"].(string)
		DataResource.FileSize = value["FileSize"].(float64)
		DataResource.LastModifiedDate = value["last_modified_date"].(string)
		DataResource.License = value["licence"].(string)
		DataResource.Copyright = value["copyright"].(string)
		DataResource.Language = value["language"].(string)

		dataSet.DataResources = append(dataSet.DataResources, DataResource)
	}

	errDb := domain.InsertNewDataSet(dataSet)
	if errDb != nil {
		return code.CodeTypeBadData
	}

	return code.CodeTypeOK

}

func deleteDataSet(body map[string]interface{}, message map[string]interface{}) uint32 {
	entity := body["entity"].(map[string]interface{})

	// data set instanceを取得
	var dataSet, err = domain.GetDataSetById(bson.ObjectIdHex(entity["id"].(string)))

	// データセットが存在しない場合
	if err != nil {
		return code.CodeTypeBadData
	}

	pubKeyBytes, errDecode := base64.StdEncoding.DecodeString(message["publicKey"].(string))

	if errDecode != nil {
		return code.CodeTypeBadData
	}

	// public keyを取得
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))

	// リクエストユーザーに権限がが存在しなかった場合
	if publicKey != dataSet.Creator {
		return code.CodeTypeUnauthorized
	}

	// データセットのMongoからの削除
	err = domain.DeleteDataSet(bson.ObjectIdHex(entity["id"].(string)))
	if err != nil {
		return code.CodeTypeBadData
	}

	return code.CodeTypeOK

}