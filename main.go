package main

import (
	"database/sql"
	"log"
	"project_2/config"
	"project_2/routes"
)

var err error

func init() {
	config.DB, err = sql.Open("mysql", config.DbURL())

	if err != nil {
		log.Fatal("Status:", err)
	}

}

func main() {

	r := routes.SetupRouter()
	r.Run()
}
