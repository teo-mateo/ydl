package config

import "os"

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
	return 5432
}

// PgDb ...
func PgDb() string {
	return "postgres"
}
