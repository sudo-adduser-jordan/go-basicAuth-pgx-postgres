package database

import (
	"context"
	"fmt"
	"log"
	"main/configs"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sudo-adduser-jordan/Toolchain/Go/styles"
)

var connection *pgxpool.Pool

func ConnectToDatabase() {

	// "postgres://username:password@localhost:5432/database_name"
	// DATABASE_URL := fmt.Sprintf(
	// 	"postgres://%s:%s@%s:%s/%s",
	// 	configs.Env("DB_USER"),
	// 	configs.Env("DB_PASSWORD"),
	// 	configs.Env("DB_HOST"),
	// 	configs.Env("DB_PORT"),
	// 	configs.Env("DB_NAME"),
	// )

	// Docker Compose
	// "postgres://username:password@database:5432/database_name"
	DATABASE_URL := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		configs.Env("DB_USER"),
		configs.Env("DB_PASSWORD"),
		configs.Env("DB_HOST"),
		configs.Env("DB_PORT"),
		configs.Env("DB_NAME"),
	)

	var err error
	connection, err = pgxpool.New(context.Background(), DATABASE_URL)
	if err != nil {
		fmt.Fprintf(
			os.Stderr,
			"Unable to connection to database: %v\n",
			err,
		)
		os.Exit(1)
	}
	fmt.Print("	-----> ")
	fmt.Println(styles.GreenText("Connected to database."))
}

func MigrateDatabase() {

	fmt.Print("	-----> ")
	fmt.Println(styles.BlueText("Database migrated."))
}

func CreateUserTable() {
	_, err := connection.Exec(context.Background(),
		CREATE_USER_TABLE,
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Print("	-----> ")
	fmt.Println(styles.BlueText("User table created."))
}

// Basic CRUD functions

func Create(username string, password string) {
	_, err := connection.Exec(context.Background(),
		INSERT_USER,
		username,
		password,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Read(username string) {
	row, err := connection.Query(context.Background(),
		SELECT_USER,
		username,
	)
	if err != nil {
		log.Fatal(err)
	}

	for row.Next() {
		var id int32
		var result string
		var password string
		err = row.Scan(&id, &result, &password)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s\n", result)
	}
}

func Update(username string, new_username string, new_password string) {
	_, err := connection.Exec(context.Background(),
		UPDATE_USER,
		username,
		new_username,
		new_password,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func Delete(username string) {
	_, err := connection.Exec(context.Background(),
		DELETE_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
}
