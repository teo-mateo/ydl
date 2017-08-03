package handlers

import (
	"log"
	"net/http"
	"strconv"
	"encoding/json"
	"os"

	"github.com/teo-mateo/ydl/ydata"
	"github.com/ydl/util"
	"io/ioutil"
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

	//check if the file is still there to delete
	if _, err := os.Stat(record.File.String); !os.IsNotExist(err) {
		//delete file from disk
		err = os.Remove(record.File.String)
		if err != nil {
			log.Println(err)
			return
		}
	}

	//delete row from db
	ydata.YQueueDelete(id)

	// redirect to list
	http.Redirect(w, r, "list", 301)
}

func MultiDeleteHandler(w http.ResponseWriter, r *http.Request){
	buf, err := ioutil.ReadAll(r.Body)

	ids := make([]int, 0)
	err = json.Unmarshal(buf, &ids)
	if err != nil{
		util.SendHTTPError(w, err);
		return
	}

	for _, id := range ids{
		record, err := ydata.YQueueGet(id)
		if err != nil {
			util.SendHTTPError(w, err)
			return
		}

		log.Printf("ID to delete: %d", id)

		if !record.File.Valid {
			util.SendHTTPError(w, err)
			return
		}

		//check if the file is still there to delete
		if _, err := os.Stat(record.File.String); !os.IsNotExist(err) {
			//delete file from disk
			err = os.Remove(record.File.String)
			if err != nil {
				util.SendHTTPError(w, err)
				return
			}
		}

		//delete row from db
		err = ydata.YQueueDelete(id)
		if err != nil{
			util.SendHTTPError(w, err)
			return
		}
	}


}
