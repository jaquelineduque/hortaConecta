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
	fmt.Fprintln(w, "Pendente")
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func ConsultAllAds(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(ConsultAllAdvertisements()); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func ConsultAd(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	vars := mux.Vars(r)
	//var id int64
	id, err := strconv.ParseInt(vars["id"], 10, 64) //strconv.Atoi(vars["id"])
	if err != nil {
		panic(err)
		//respondWithError(w, http.StatusBadRequest, "Invalid product ID")
		//return
	}

	if err := json.NewEncoder(w).Encode(ConsultAdvertisement(id)); err != nil {
		panic(err)
	}
	w.WriteHeader(http.StatusOK)
}

func InsertAd(w http.ResponseWriter, r *http.Request) {
	var ads Advertisements
	var adsReturn Advertisements
	var ret ReturnAdvertisements

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := r.Body.Close(); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	if err := json.Unmarshal(body, &ads); err != nil {
		http.Error(w, err.Error(), 422)
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
		panic(err)
	}

}
