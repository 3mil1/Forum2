package main

import (
	"flag"
	"forum/internal/app/api"
	"log"
	"strconv"
)

var (
	port        int
	db          string
	//createNewDb bool
)

func init() {
	flag.IntVar(&port, "port", 8080, "Specify the port to listen to.")
	flag.StringVar(&db, "db", "dataBase.db", "Specify path to database")
	//flag.BoolVar(&createNewDb, "createDB", false, "Specify whether to create a new database")
}

func main() {
	flag.Parse()
	log.Println("It works")

	//server instance initialization
	config := api.NewConfig(strconv.Itoa(port), db)
	server := api.New(config)

	//start server
	log.Fatal(server.Start())
}
