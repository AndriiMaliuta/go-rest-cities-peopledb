package main

import (
	"awesomeProject/models"
	"awesomeProject/web"
	"encoding/json"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	//var city models.City
	//city = models.City{5, "Marocco", 562456457}
	//repo.InsertCity(city)
	//
	//// All
	//cities := repo.FindAll()
	//
	//for _, city := range cities {
	//	cityJson, err := json.Marshal(city)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	fmt.Println(string(cityJson))
	//}
	var ctrs []models.Country
	err := json.Unmarshal(web.GetCountries(), &ctrs)
	if err != nil {
		log.Println(err)
	}
	ctr := ctrs[0]
	fmt.Println(ctr)
}
