package route

import (
	"gopkg.in/mgo.v2/bson"
	"open-data-repository/src/open-data-repository-abci/common/code"
	"strings"
)

func RouteCheckTx(body map[string]interface{}, message map[string]interface{}) uint32 {
	code := code.CodeTypeOK

	switch body["type"] {
	case "createUser":
		code = checkTxCreateUser(body)
		break
	case "addDataSet":
		code = checkTxAddDataSet(body)
		break
	case "addDataResource":
		code = checkTxAddDataResource(body)
		break
	}

	return code
}

// formatのチェックを行う

func checkTxCreateUser(body map[string]interface{}) (codeType uint32) {
	entity := body["entity"].(map[string]interface{})

	if (entity["id"] == nil) || (bson.IsObjectIdHex(entity["id"].(string)) != true) {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["name"] == nil) || (strings.TrimSpace(entity["name"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	return code.CodeTypeOK
}

func checkTxAddDataSet(body map[string]interface{}) (codeType uint32) {
	entity := body["entity"].(map[string]interface{})

	if (entity["id"] == nil) || (bson.IsObjectIdHex(entity["id"].(string)) != true) {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["title"] == nil) || (strings.TrimSpace(entity["title"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["publisher"] == nil) || (strings.TrimSpace(entity["publisher"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["contact_point"] == nil) || (strings.TrimSpace(entity["contact_point"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	var tags = entity["tags"].([]string)
	for _, v := range tags {
		if strings.TrimSpace(v) == "" {
			codeType = code.CodeTypeBadData
			return
		}
	}

	if entity["release_date"] == nil {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["frequency_of_update"] == nil) || (strings.TrimSpace(entity["frequency_of_update"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["landing_page"] == nil) || (strings.TrimSpace(entity["landing_page"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (entity["spatial"] == nil) || (strings.TrimSpace(entity["spatial"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	rowDataResources := entity["data_resources"].([]map[string]interface{})
	for _, value := range rowDataResources {
		codeType = checkDataResource(value)
		if codeType == code.CodeTypeBadData {
			return codeType
		}
	}

	return code.CodeTypeOK
}

func checkTxAddDataResource(body map[string]interface{}) (codeType uint32) {
	entity := body["entity"].(map[string]interface{})

	if (entity["id"] == nil) || (bson.IsObjectIdHex(entity["id"].(string)) != true) {
		codeType = code.CodeTypeBadData
		return
	}

	rowDataResources := entity["data_resources"].([]map[string]interface{})
	for _, value := range rowDataResources {
		codeType = checkDataResource(value)
		if codeType == code.CodeTypeBadData {
			return codeType
		}
	}

	return code.CodeTypeOK
}

func checkDataResource(rowDataResource map[string]interface{}) (codeType uint32) {
	if (rowDataResource["id"] == nil) || (bson.IsObjectIdHex(rowDataResource["id"].(string)) != true) {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["title"] == nil) || (strings.TrimSpace(rowDataResource["title"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["url"] == nil) || (strings.TrimSpace(rowDataResource["url"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["format"] == nil) || (strings.TrimSpace(rowDataResource["format"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["value"] == nil) || (strings.TrimSpace(rowDataResource["value"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["file_size"] == nil) || (strings.TrimSpace(rowDataResource["file_size"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if rowDataResource["last_modified_date"] == nil {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["license"] == nil) || (strings.TrimSpace(rowDataResource["license"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["copy_right"] == nil) || (strings.TrimSpace(rowDataResource["copy_right"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	if (rowDataResource["language"] == nil) || (strings.TrimSpace(rowDataResource["language"].(string)) == "") {
		codeType = code.CodeTypeBadData
		return
	}

	return code.CodeTypeOK
}
