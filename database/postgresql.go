
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // PostgreSQL driver
)

var DB *sql.DB

func ConnectPostgreSql() error {
	config := struct {
		Host     string
		Port     int
		User     string
		Password string
		Database string
	}{"localhost", 8080, "admin", "admin@123", "db1"}

	source := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", config.Host, config.Port, config.User, config.Password, config.Database)

	database, err := sql.Open("postgres", source)
	if err != nil {
		log.Fatal("Failed to open DB:", err)
	}

	if err := database.Ping(); err != nil {
		log.Fatal("Failed to ping DB:", err)
	}

	database.SetMaxOpenConns(20)
	DB = database
	return nil
}
