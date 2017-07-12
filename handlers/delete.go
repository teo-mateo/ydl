package handlers

import (
	"log"
	"net/http"
	"strconv"

	"os"

	"github.com/teo-mateo/ydl/ydata"
)

//DeleteHandler ...
func DeleteHandler(w http.ResponseWriter, r *http.Request) {

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		log.Println(err)
		return
	}

	record, err := ydata.YQueueGet(id)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("ID to delete: %d", id)

	if !record.File.Valid {
		log.Println("No file to delete.")
		return
	}

	//delete file from disk
	err = os.Remove(record.File.String)
	if err != nil {
		log.Println(err)
		return
	}

	//delete row from db
	ydata.YQueueDelete(id)

	// redirect to list
	http.Redirect(w, r, "list", 301)

}
