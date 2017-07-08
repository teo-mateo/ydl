package config

import "os"
import "strconv"

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
