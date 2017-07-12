package util

import (
	"os"
)

// FileExists tired of typing this shit all the time.
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}
