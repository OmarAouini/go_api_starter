package main

import "github.com/OmarAouini/golang_api_starter/database"

func main() {
	Init()
	database.ConnectDb()
	PrintAppInfo()
}
