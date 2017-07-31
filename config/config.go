package config

import (
	"os"
	"path/filepath"
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

// TempFolder ...
func TempFolder() (string, error) {
	cwd, error := os.Getwd()
	if error != nil {
		return "", error
	}

	return filepath.Join(cwd, "tmp"), nil
}

// PgConnectionString ...
func PgConnectionString() string {
	connString := os.Getenv("PG_CN")
	return connString
	//return fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", PgHost(), PgPort(), PgUser(), PgPassword(), PgDb())
}
