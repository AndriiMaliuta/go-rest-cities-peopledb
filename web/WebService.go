package web

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetCity(name string) []byte {
	resp, err := http.Get(fmt.Sprintf("https://restcountries.com/v3.1/capital/%s", name))
	if err != nil {
		log.Println(err)
	}
	webCity, err := ioutil.ReadAll(resp.Body)
	//return string(webCities)
	return webCity
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
