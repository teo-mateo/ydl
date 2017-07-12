package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"strconv"

	"github.com/teo-mateo/ydl/util"
	ydata "github.com/teo-mateo/ydl/ydata"
)

// DownloadHandler ...
func DownloadHandler(w http.ResponseWriter, r *http.Request) {

	//get id from url
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		fmt.Println(err)
		return
	}

	//get file name
	result, err := ydata.YQueueGet(id)
	if err != nil {
		fmt.Println(err)
		return
	}

	//check if file exists
	if !util.FileExists(result.File.String) {
		fmt.Printf("Requested file was not found: %s", result.File.String)
		return
	}

	//serve file
	_, shortname := filepath.Split(result.File.String)
	fmt.Println("Serving file: " + shortname)
	w.Header().Set("Content-type", "application/mp3")
	w.Header().Set("Content-Disposition", "attachment; filename="+shortname)
	http.ServeFile(w, r, result.File.String)
}
