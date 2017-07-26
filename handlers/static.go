package handlers

import (
	"net/http"
	"os"
	"path/filepath"
)

//StaticFilesHandler ...
func StaticFilesHandler(rw http.ResponseWriter, req *http.Request) {
	//get the url of the file to serve
	src := req.RequestURI[1:]

	//file should be under current dir
	cwd, err := os.Getwd()
	if err != nil {
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	//build file name
	fname := filepath.Join(cwd, src)
	if _, err := os.Stat(fname); os.IsNotExist(err) {
		rw.WriteHeader(http.StatusNotFound)
		return
	}

	//serve file
	http.ServeFile(rw, req, fname)
}
