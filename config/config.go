package config

import (
	"fmt"
	"os"
	"strconv"
)

// PgUser ...
func PgUser() string {
	return os.Getenv("PG_USER")
}

// PgPassword ...
func PgPassword() string {
	return os.Getenv("PG_PASSWORD")
}

// PgHost ...
func PgHost() string {
	return "localhost"
}

// PgPort ...
func PgPort() int {
	port, _ := strconv.Atoi(os.Getenv("PG_PORT"))
	return port
}

// PgDb ...
func PgDb() string {
	return "postgres"
}

// PgConnectionString ...
func PgConnectionString() string {
	return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PgHost(), PgPort(), PgUser(), PgPassword(), PgDb())
}
