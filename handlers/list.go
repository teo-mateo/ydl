package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"encoding/json"

	"github.com/teo-mateo/ydl/util"
	"github.com/teo-mateo/ydl/ydata"
)

// ListHandler ...
func ListHandler(w http.ResponseWriter, r *http.Request) {

	who := r.RequestURI[6:]

	records, err := ydata.YQueueGetAll(ydata.STATUSDownloaded3, who)
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
	if err != nil {
		w.Write([]byte(err.Error()))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

// ListJSONPerUser ...
func ListJSONPerUser(w http.ResponseWriter, r *http.Request) {
	who := r.RequestURI[11:]

	records, err := ydata.YQueueGetAll(ydata.STATUSDownloaded3, who)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	jsonBytes, err := json.Marshal(records)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Write(jsonBytes)
	return
}

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	users, err := ydata.YQueueGetUsers()
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	bytes, err := json.Marshal(users)
	if err != nil {
		util.SendHTTPError(w, err)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	w.Write(bytes)
	return
}
