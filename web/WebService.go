package web

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetCities() []byte {
	resp, err := http.Get("https://restcountries.com/v3.1/capital/Lima")
	if err != nil {
		log.Println(err)
	}
	webCities, err := ioutil.ReadAll(resp.Body)
	//return string(webCities)
	return webCities
}

func GetCountries() []byte {
	// https://restcountries.com/v3.1/all
	resp, err := http.Get("https://restcountries.com/v3.1/all")
	if err != nil {
		log.Println(err)
	}
	webCts, err := ioutil.ReadAll(resp.Body)
	//return string(webCities)
	return webCts
}
