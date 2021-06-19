//This example uses the ORM jet
package main

import (
	"log"
	"os"

	"github.com/eaigner/jet"
	"github.com/joho/godotenv"
	"github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	pgUrl, err := pq.ParseURL(os.Getenv("ELEPHANTSQL_URL"))
	if err != nil {
		log.Fatal(err)
	}
	db, err := jet.Open("postgres", pgUrl)
	if err != nil {
		log.Fatal(err)
	}

	var cities []*struct {
		Name     string
		Location string
	}
	if err := db.Query("SELECT * FROM cities").Rows(&cities); err != nil {
		log.Fatal(err)
	}
	for _, city := range cities {
		log.Printf("Name: %v, Point: %v\n",
			city.Name,
			city.Location)
	}
}
