package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"os"

	ydlconf "github.com/teo-mateo/ydl/config"
)

//DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}

	log.Println("id to delete: " + strconv.Itoa(id))

	psqlInfo := ydlconf.PgConnectionString()
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	//get file name from db
	row := db.QueryRow("SELECT file FROM yqueue WHERE id=" + strconv.Itoa(id))
	var fname string
	err = row.Scan(&fname)
	if err != nil {
		log.Println(err)
		return
	}

	//delete file from disk
	err = os.Remove(fname)
	if err != nil {
		log.Println(err)
		return
	}

	//delete row from db
	_, err = db.Exec("DELETE FROM yqueue WHERE id=" + strconv.Itoa(id))
	if err != nil {
		log.Println(err)
		return
	}

	// redirect to list
	http.Redirect(w, r, "list", 301)

}
