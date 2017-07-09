package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"path/filepath"

	ydlconf "github.com/teo-mateo/ydl/config"
)

// DownloadHandler ...
func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	psqlInfo := ydlconf.PgConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	row := db.QueryRow("SELECT file FROM yqueue WHERE status = 3 AND id = " + id)
	if err != nil {
		fmt.Println(err)
		return
	}

	var fname string
	row.Scan(&fname)
	_, shortname := filepath.Split(fname)
	w.Header().Set("Content-type", "application/mp3")
	w.Header().Set("Content-Disposition", "attachment; filename="+shortname)
	http.ServeFile(w, r, fname)
}
