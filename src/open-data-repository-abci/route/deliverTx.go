package route

import "open-data-repository/src/open-data-repository-abci/common/code"

func RouteDeliverTx(body map[string]interface{}, message map[string]interface{}) uint32 {
	code := code.CodeTypeOK

	switch body["type"] {
	case "createUser":
		break
	case "addDataResource":
		break
	case "addDataSet":
		break
	}

	return code
}