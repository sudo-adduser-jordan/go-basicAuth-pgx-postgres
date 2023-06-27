package database

import (
	"context"
	"log"
)

func CreateUser(username string, password string) bool {
	_, err := connection.Exec(context.Background(),
		INSERT_USER,
		username,
		password,
	)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func ReadUser(username string) string {
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
		if result == username {
			return result
		}
	}
	return "error reading username."
}

func ReadPassword(username string) string {
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
		return password
	}
	return "Read password error."
}

func UpdateUser(username string, new_username string, new_password string) {
	_, err := connection.Exec(context.Background(),
		UPDATE_USER,
		new_username,
		new_password,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
}

func DeleteUser(username string) {
	_, err := connection.Exec(context.Background(),
		DELETE_USER,
		username,
	)

	if err != nil {
		log.Fatal(err)
	}
}
