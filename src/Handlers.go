package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func FuncaoAqui(w http.ResponseWriter, r *http.Request) {
	/*file, err := os.Open("config/config.json")
	if err != nil {
		panic(err)
		//respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		//return
	}
	if file != nil {
		fmt.Fprintln(w, "Conseguiu")
	}
	*/

	fmt.Fprintln(w, "Pendente")
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func formatErrorResponse(w http.ResponseWriter, statusCode int, internalState int, message string, technicalMessage string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	var ret ReturnStruct
	ret.State = internalState
	ret.Message = message
	ret.TechnicalMessage = technicalMessage

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}
}

func ConsultAllAds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(ConsultAllAdvertisements()); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func ConsultAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	//var id int64
	id, err := strconv.ParseInt(vars["id"], 10, 64) //strconv.Atoi(vars["id"])
	if err != nil {

		formatErrorResponse(w, 422, 422, "Id to be consulted not found.", err.Error())
		//panic(err)
		//respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		//return
	}

	if err := json.NewEncoder(w).Encode(ConsultAdvertisement(id)); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func InsertAd(w http.ResponseWriter, r *http.Request) {
	var ads Advertisements
	var adsReturn Advertisements
	var ret ReturnAdvertisements

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		//http.Error(w, err.Error(), 500)
		formatErrorResponse(w, 500, 500, "Json request couldn't be read.", err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		formatErrorResponse(w, 500, 500, "Request body couldn't be closed.", err.Error())
		//http.Error(w, err.Error(), 500)
		return
	}
	if err := json.Unmarshal(body, &ads); err != nil {
		formatErrorResponse(w, 422, 422, "Request should be a valid json.", err.Error())
		//http.Error(w, err.Error(), 422)
		return
	}

	for _, ad := range ads {
		adReturn := InsertAdvertisement(ad)
		adsReturn = append(adsReturn, adReturn)
	}

	ret.Return.State = 1
	ret.Return.Message = "Advertisement(s) inserted successfully"
	ret.Ads = adsReturn
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}

}
