package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"

	"github.com/yousefzinsazk78/fiber_post_second_version/internal/routes"
)

const (
	host     = "localhost"
	portt    = 5432
	user     = "postgres"
	password = "13781378"
	dbname   = "simpletestpostdb"
)

func main() {
	///set port in flag of program
	port := flag.String("default port", ":8000", "you can set your custom font")
	flag.Parse()

	pgqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, portt, user, password, dbname)
	db, err := sql.Open("postgres", pgqlInfo)
	if err != nil {
		log.Fatal(err)
	}

	routes.Run(*port, db)
}
