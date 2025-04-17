package main

import (
	"log"

	"github.com/Alexander-s-Digital-Marketplace/core-service/internal/database"
)

func main() {
	var db database.DataBase
	db.InitDB()

	query := []string{
		`DELETE FROM profiles;`,
		`DELETE FROM products;`,
		`DELETE FROM items;`,
		`DELETE FROM carts;`,
		`DELETE FROM histories;`,
	}

	for _, stmt := range query {
		if err := db.Connection.Exec(stmt).Error; err != nil {
			log.Println("Error executing clear: ", stmt, err)
		}
	}

	log.Println("All table is clear")

	db.CloseDB()
}
