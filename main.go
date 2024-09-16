package main

import (
	"github.com/GabrielVictorP/golang-restAPI.git/database"
	"github.com/GabrielVictorP/golang-restAPI.git/server"
)

func main() {
	database.StartDB()
	server := server.NewServer()

	server.Run()
}
