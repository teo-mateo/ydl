package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	ydlconf "github.com/teo-mateo/ydl/config"
	"github.com/teo-mateo/ydl/handlers"
	"github.com/teo-mateo/ydl/util/cleanup"
)

func main() {
	fmt.Println("Starting YDL, port 8080.")
	fmt.Println("PG Conn: " + ydlconf.PgConnectionString())
	fmt.Println("Will convert to mp3")

	cleanup.StartCleanupRoutine()

	http.HandleFunc("/", handlers.RootHandler)
	http.HandleFunc("/users", handlers.UsersHandler)
	http.HandleFunc("/ydl", handlers.QueueHandler)
	http.HandleFunc("/list/", handlers.ListHandler)
	http.HandleFunc("/list/json/", handlers.ListJSONPerUser)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.HandleFunc("/multidownload", handlers.MultiDownloadHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)
	http.HandleFunc("/static/", handlers.StaticFilesHandler)
	http.HandleFunc("/app", handlers.RootHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
