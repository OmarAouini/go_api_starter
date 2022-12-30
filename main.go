package main

import (
	"github.com/OmarAouini/go_tdd/database"
)

func main() {
	Init()
	database.ConnectDb()
	PrintAppInfo()

}
