package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

const (
	host = "DB_HOST"
	port = "DB_PORT"
	user = "DB_USER"
	pass = "DB_PASS"
	name = "DB_NAME"
)

func GetDatabase() (db *sql.DB, err error) {

	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		checkIfEnvVarExists(host),
		checkIfEnvVarExists(port),
		checkIfEnvVarExists(user),
		checkIfEnvVarExists(pass),
		checkIfEnvVarExists(name),
	)

	db, err = sql.Open("postgres", connStr)
	return
}

func checkIfEnvVarExists(envVar string) string {
	env := os.Getenv(envVar)
	if env == "" {
		fmt.Printf("Using value for %s default\n", envVar)
		switch envVar {
		case host:
			env = "localhost"
		case port:
			env = "5432"
		case user:
			env = "postgres"
		case pass:
			env = "root"
		case name:
			env = "bossabox"
		}
	}

	return env
}
