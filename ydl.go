package main

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	ydlconf "github.com/teo-mateo/ydl/config"
	"github.com/teo-mateo/ydl/handlers"
)

func main() {
	fmt.Println("Starting YDL, port 8080.")
	fmt.Println("PG Conn: " + ydlconf.PgConnectionString())

	http.HandleFunc("/ydl", handlers.QueueHandler)
	http.HandleFunc("/list", handlers.ListHandler)
	http.HandleFunc("/download", handlers.DownloadHandler)
	http.HandleFunc("/delete", handlers.DeleteHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
