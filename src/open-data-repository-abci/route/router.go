package route

import (
	"encoding/base64"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/tendermint/tendermint/abci/types"
	"golang.org/x/crypto/ed25519"
	"open-data-repository/src/domain"
	"open-data-repository/src/infrastructure"
	"open-data-repository/src/open-data-repository-abci/common/code"
	"open-data-repository/src/open-data-repository-abci/common/util"
	"strings"
)

var _ types.Application = (*JSONStoreApplication)(nil)

// JSONStoreApplication ...
type JSONStoreApplication struct {
	types.BaseApplication
	state state
}

type state struct {
	LastBlockHeight int64
	LastBlockAppHash []byte
	//ValidatorUpdates []types.Validators
}

// NewJSONStoreApplication ...
func NewJSONStoreApplication() *JSONStoreApplication {
	return &JSONStoreApplication{}
}

// Info ...
func (app *JSONStoreApplication) Info(req types.RequestInfo) (resInfo types.ResponseInfo) {
	return types.ResponseInfo{Data: fmt.Sprintf("{\"size\":%v}", app.state.LastBlockHeight)}
}

// DeliverTx ... Update the global state
func (app *JSONStoreApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {

	var temp interface{}
	err := json.Unmarshal(tx, &temp)

	if err != nil {
		panic(err)
	}

	message := temp.(map[string]interface{})
	var bodyTemp interface{}

	errBody := json.Unmarshal([]byte(message["body"].(string)), &bodyTemp)

	if errBody != nil {
		panic(errBody)
	}

	body := bodyTemp.(map[string]interface{})

	fmt.Println("【DeliverTx】:", body)

	// 返り値用のCodeTypeを作成
	resCode := RouteDeliverTx(body, message)

	// TODO error内容によってcodeを出し分ける
	return types.ResponseDeliverTx{Code: resCode, Tags: nil}
}

// CheckTx ... Verify the transaction => transactionの有効性確認
// SimpleにtransactionのFormatを検証する => DeliverTxで担保できるならば不要
func (app *JSONStoreApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	var temp interface{}
	err := json.Unmarshal(tx, &temp)

	if err != nil {
		panic(err)
	}

	message := temp.(map[string]interface{})

	// ==== Signature Validation =======
	pubKeyBytes, err := base64.StdEncoding.DecodeString(message["publicKey"].(string))
	sigBytes, err := hex.DecodeString(message["signature"].(string))
	messageBytes := []byte(message["body"].(string))

	isCorrect := ed25519.Verify(pubKeyBytes, messageBytes, sigBytes)

	if isCorrect != true {
		return types.ResponseCheckTx{Code: code.CodeTypeBadSignature}
	}
	// ==== Signature Validation =======

	var bodyTemp interface{}

	errBody := json.Unmarshal([]byte(message["body"].(string)), &bodyTemp)

	if errBody != nil {
		panic(errBody)
	}

	body := bodyTemp.(map[string]interface{})

	fmt.Println("【CheckTx】:", body)

	// ==== Does the userDao really exist? ======
	if body["type"] != "createUser" {
		publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))

		// ここでmongoに今のstateでuserが存在しているか確認している
		count := domain.CheckExistenceOfUser(publicKey)

		if count == 0 {
			return types.ResponseCheckTx{Code: code.CodeTypeBadData}
		}
	}
	// ==== Does the userDao really exist? ======

	// ===== Data Validation =======
	codeType := RouteCheckTx(body, message)
	// ===== Data Validation =======

	return types.ResponseCheckTx{Code: codeType}
}

// Commit ...Commit the block. Calculate the appHash
func (app *JSONStoreApplication) Commit() types.ResponseCommit {
	appHash := make([]byte, 8)

	count := infrastructure.FindTotalDocuments()
	binary.PutVarint(appHash, count)

	fmt.Println("【commit】count => ", count, " appHash => ",base64.StdEncoding.EncodeToString(appHash))
	app.state.LastBlockAppHash = appHash
	return types.ResponseCommit{Data: appHash}
}

// Query ... Query the blockchain. Unimplemented as of now.
func (app *JSONStoreApplication) Query(reqQuery types.RequestQuery) (resQuery types.ResponseQuery) {


	fmt.Println("【Query】", reqQuery.Path)

	resQuery.Log = base64.StdEncoding.EncodeToString(reqQuery.Data)
	return
}

// BeginBlock implements the ABCI application interface.
func (app *JSONStoreApplication) BeginBlock(req types.RequestBeginBlock) (res types.ResponseBeginBlock) {
	// fmt.Println("【BeginBlock】Proposer Address => ", base64.StdEncoding.EncodeToString(req.Header.GetProposerAddress()))
	//fmt.Println("....BlockHeight => ", req.Header.Height)
	//app.state.LastBlockHeight = req.Header.Height
	// TODO get Validator address
	return
}

func (app *JSONStoreApplication) EndBlock(req types.RequestEndBlock) (res types.ResponseEndBlock) {

	//// TODO Add Validator logic
	//if len(app.state.ValidatorUpdates) != 0 {
	//	// 新たなValidatorを追加する
	//	res.ValidatorUpdates = app.state.ValidatorUpdates
	//	// stateを削除する
	//	app.state.ValidatorUpdates = []types.ValidatorUpdate{}
	//}
	fmt.Println("【EndBlock】", res.String())
	return
}
