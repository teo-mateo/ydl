package util

import (
	"log"
	"net/http"
	"os"
)

// FileExists tired of typing this shit all the time.
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// HTTPError ...
func SendHTTPError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, err.Error(), http.StatusInternalServerError)
}
