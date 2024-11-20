package main

import (
	ginApi "musicboxd/api"
	"musicboxd/database"
	hlp "musicboxd/hlp"
)

func main() {
	hlp.LoadEnvKeys()

	db := database.GetDB()
	defer db.Close()

	api := ginApi.NewApi()

	api.Start()
}
