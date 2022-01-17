package repo

import (
	"awesomeProject/models"
	"database/sql"
	"fmt"
	"log"
	"os"
)

func initDb() *sql.DB {
	pswd, _ := os.LookupEnv("PSQL_PSWD")
	db, err := sql.Open("postgres", fmt.Sprintf("postgres://dev2:%s@localhost/cities?sslmode=disable", pswd))
	if err != nil {
		log.Panicln(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("You connected to your database.")
	return db
}

func InsertCity(c models.City) {
	query := fmt.Sprintf("insert into cities (id, name, population) values (%d, '%s', %d);", c.Id, c.Name, c.Population)
	db := initDb()
	log.Println(query)
	res, err := db.Query(query)
	if err != nil {
		log.Panicln(res, err)
	}
	log.Printf("City created with id = %d", c.Id)
	defer db.Close()
}

func FindAll() []models.City {
	db := initDb()
	rows, err := db.Query("select * from cities")
	if err != nil {
		log.Panicln(rows, err)
	}
	cities := make([]models.City, 0)

	for rows.Next() {
		city := models.City{}
		err := rows.Scan(&city.Id, &city.Name, &city.Population)
		if err != nil {
			log.Panicln(err)
		}
		cities = append(cities, city)
	}

	defer db.Close()

	return cities
}
