package handlers

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	ydlconf "github.com/teo-mateo/ydl/config"
)

// ListHandler ...
func ListHandler(w http.ResponseWriter, r *http.Request) {

	psqlInfo := ydlconf.PgConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id, file FROM yqueue WHERE status = 3")
	if err != nil {
		fmt.Println(err)
		return
	}

	var id int
	var fname string
	m := make(map[int]string)

	for rows.Next() {
		err = rows.Scan(&id, &fname)
		m[id] = fname
	}

	t, _ := template.ParseFiles("ytdlist.html")
	t.Execute(w, m)
}
