package main

import (
	ginApi "musicboxd/api"
	hlp "musicboxd/hlp"
)

func main() {
	hlp.LoadEnvKeys()
	api := ginApi.NewApi()

	api.Start()
}
