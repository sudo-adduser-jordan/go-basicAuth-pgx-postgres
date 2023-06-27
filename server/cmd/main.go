package main

import (
	"fmt"
	"main/database"
	"main/routes"

	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

func main() {
	fmt.Println()
	fmt.Print(styles.BlueLabel(" Go 1.20 "))
	fmt.Print(styles.PurpleLabel(" Postgres 15 "))
	fmt.Println(styles.BlueLabel(" pgx v5 "))

	database.ConnectToDatabase()
	// database.MigrateDatabase()
	database.CreateUserTable()
	routes.SetupRoutes()
}
