package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Advertisement",
		"GET",
		"/advertisement",
		ConsultAllAds,
	}, Route{
		"Advertisement",
		"GET",
		"/advertisement/{id}",
		ConsultAd,
	}, Route{
		"Advertisement",
		"POST",
		"/advertisement",
		InsertAd,
	},
}
