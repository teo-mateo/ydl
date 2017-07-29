package handlers

import (
	"fmt"
	"net/http"
	"path/filepath"

	"strconv"

	"github.com/teo-mateo/ydl/util"
	"github.com/teo-mateo/ydl/util/cleanup"
	ydata "github.com/teo-mateo/ydl/ydata"
	"strings"
	"time"
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

func MultiDownloadHandler(w http.ResponseWriter, r *http.Request) {
	//get ids from url
	idsParam := r.URL.Query().Get("ids")
	ids := strings.Split(idsParam, ",")

	files := make([]string, 0)

	for _, x := range ids {
		id, err := strconv.Atoi(x)
		if err != nil {
			util.SendHTTPError(w, err)
			return
		}
		song, err := ydata.YQueueGet(id)
		if err != nil {
			util.SendHTTPError(w, err)
		}

		files = append(files, song.File.String)
	}

	zipFile, err := util.Zip(files)
	if err != nil {
		util.SendHTTPError(w, err)
		return
	}

	_, shortname := filepath.Split(zipFile)
	fmt.Println("Serving file: " + shortname)
	w.Header().Set("Content-type", "application/zip")
	w.Header().Set("Content-Disposition", "attachment; filename="+shortname)
	http.ServeFile(w, r, zipFile)

	//will be deleted after N seconds
	cleanup.MarkForCleanup(zipFile, time.Second*60*10)
}
