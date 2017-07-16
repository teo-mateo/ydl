package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/teo-mateo/ydl/ydata"
)

// ListHandler ...
func ListHandler(w http.ResponseWriter, r *http.Request) {

	records, err := ydata.YQueueGetAll(ydata.STATUSDownloaded3)
	if err != nil {
		fmt.Println(err)
		return
	}

	m := make(map[int]string)

	for _, record := range records {
		m[record.ID] = record.File.String
	}

	t, _ := template.ParseFiles("ytdlist.html")
	err = t.Execute(w, m)
	if err != nil{
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
