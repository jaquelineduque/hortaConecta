package main

type Advertisement struct {
	Id           int64  `json:"id"`
	State        int    `json:"state"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	IdAdvertiser int    `json:"idAdvertiser"`
}

type Advertisements []Advertisement

type ReturnStruct struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

type ReturnAdvertisements struct {
	Return ReturnStruct   `json:"return"`
	Ads    Advertisements `json:"ads"`
}
