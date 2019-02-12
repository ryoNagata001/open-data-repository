package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"log"
	"net/http"
	"open-data-repository/src/domain"
	"open-data-repository/src/open-data-repository-abci/common/util"
	"strconv"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

)

func GetDataSetByObjectId(w http.ResponseWriter, r *http.Request) {
	objectId := bson.ObjectIdHex(r.FormValue("id"))
	dataset, err := domain.GetDataSetById(objectId)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, dataset)
	// fmt.Fprintln(w, domain.GetDataSetAll())
}

func FindMyDataSet(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	publickey := r.FormValue("publicKey")
	pubKeyBytes, err := base64.StdEncoding.DecodeString(publickey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	dataSet, errDb := domain.GetMyDataSet(publicKey)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, dataSet)
	// fmt.Fprintln(w, "not implemented yet !")
}

func SearchDataSet(w http.ResponseWriter, r *http.Request) {
	title := r.FormValue("title")
	publisher := r.FormValue("publisher")
	tags := r.FormValue("tags")
	spatial := r.FormValue("spatial")

	fmt.Println(spatial)

	dataSet, errDb := domain.SearchDataSet(title, publisher, tags, spatial)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, errDb.Error())
		return
	}
	respondWithJson(w, http.StatusOK, dataSet)
}

func FindUserName(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	publickey := r.FormValue("publicKey")
	pubKeyBytes, err := base64.StdEncoding.DecodeString(publickey)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	publicKey := strings.ToUpper(util.ByteToHex(pubKeyBytes))
	dataSet, errDb := domain.GetUserByPubKey(publicKey)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, dataSet)
}

func GetDataSetList(w http.ResponseWriter, r *http.Request) {
	// params := mux.Vars(r)
	page, _ := strconv.Atoi(r.FormValue("page"))
	perPage, _ := strconv.Atoi(r.FormValue("perPage"))
	dataSet, errDb := domain.GetDataSetList(page, perPage)
	if errDb != nil {
		respondWithError(w, http.StatusInternalServerError, errDb.Error())
		return
	}
	count, errCollection := domain.GetCollectionCount()
	if errCollection != nil {
		respondWithError(w, http.StatusInternalServerError, errCollection.Error())
		return
	}
	result := map[string]interface{}{
		"dataset": dataSet,
		"count": count,
	}
	respondWithJson(w, http.StatusOK, result)
	// fmt.Fprintln(w, "not implemented yet !")
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/datasets", GetDataSetByObjectId).Methods("GET")
	r.HandleFunc("/datasets/search", SearchDataSet).Methods("GET")
	r.HandleFunc("/datasets/user", FindMyDataSet).Methods("GET")
	r.HandleFunc("/user", FindUserName).Methods("GET")
	r.HandleFunc("/datasets/list", GetDataSetList).Methods("GET")
	if err := http.ListenAndServe(":3000", handlers.CORS()(r)); err != nil {
		log.Fatal(err)
	}
}