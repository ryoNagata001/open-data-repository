package route

import (
	"encoding/base64"
	"eventToken/utils"
	"google.golang.org/genproto/googleapis/type/date"
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/domain"
	"open-data-repository/src/open-data-repository-abci/common/code"
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
	publicKey := strings.ToUpper(utils.ByteToHex(pubKeyBytes))
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
	publicKey := strings.ToUpper(utils.ByteToHex(pubKeyBytes))
	dataSet.Creator = publicKey

	dataSet.Tags = entity["tags"].([]string)
	dataSet.ReleaseDate = entity["release_date"].(date.Date)
	dataSet.FrequencyOfUpdate = entity["frequency_of_update"].(string)
	dataSet.LandingPage = entity["landing_page"].(string)
	dataSet.Spatial = entity["spatial"].(string)

	var DataResoueces = entity["data_resources"].([]map[string]interface{})
	for i, value := range DataResoueces {
		dataSet.DataResources[i].ID = bson.ObjectIdHex(value["id"].(string))
		dataSet.DataResources[i].Title = value["title"].(string)
		dataSet.DataResources[i].URL = value["url"].(string)
		dataSet.DataResources[i].Description = value["description"].(string)
		dataSet.DataResources[i].Format = value["format"].(string)
		dataSet.DataResources[i].Value = value["value"].(string)
		dataSet.DataResources[i].FileSize = value["file_size"].(int)
		dataSet.DataResources[i].LastModifiedDate = value["last_modified_date"].(date.Date)
		dataSet.DataResources[i].License = value["licence"].(string)
		dataSet.DataResources[i].Copyright = value["copyright"].(string)
		dataSet.DataResources[i].Language = value["language"].(string)
	}

	err := domain.InsertNewDataSet(dataSet)
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
		DataResource.FileSize = value["FileSize"].(int)
		DataResource.LastModifiedDate = value["last_modified_date"].(date.Date)
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