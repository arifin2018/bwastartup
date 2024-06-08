package main

import (
	"bwastartup/config"
	"bwastartup/routes"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var db = config.DB
	routes.Router()
	fmt.Println(db)
}
