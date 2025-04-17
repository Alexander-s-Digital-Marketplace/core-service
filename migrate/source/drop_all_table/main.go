package main

import (
	"log"

	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DROP TABLE profiles CASCADE;`,
		`DROP TABLE products CASCADE;`,
		`DROP TABLE items CASCADE;`,
		`DROP TABLE carts CASCADE;`,
		`DROP TABLE histories CASCADE;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing drop: ", stmt, err)
		}
	}

	log.Println("All table is droped")

	db.CloseDB()
}
