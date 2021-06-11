package db

import (
	"database/sql"

	_ "github.com/lib/pq" // não to usando agora mais em algum momento vou usar
)

func DBConnection() *sql.DB {
	stringConnection := "user=postgres dbname=store password=245 host=localhost sslmode=disable"

	connection, err := sql.Open("postgres", stringConnection)

	if err != nil {
		panic(err.Error())
	}

	return connection
}
