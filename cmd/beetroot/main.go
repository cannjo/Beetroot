package main

import (
	"github.com/betauia/beetroot/api"
	"github.com/betauia/beetroot/config"
	"github.com/betauia/beetroot/db"
)

func main() {
	config.SetupConfig()
	config.ReadData()
	db.DBConnect()
	api.StartAPI(8080)
}
