package main

import (
	"database/sql"
	"stonewall-api/app/routes"
	database "stonewall-api/config"
)

func main() {
	db := database.InitDatabaseConnection()

	conn, _ := db.DB()

	defer func(conn *sql.DB) {
		err := conn.Close()
		if err != nil {
			return
		}
	}(conn)

	routes.SetupRouter(db)
}
