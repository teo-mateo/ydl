package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	_ "github.com/lib/pq"
	"github.com/rylio/ytdl"

	ydlconf "github.com/teo-mateo/ydl/config"
)

func main() {

	if len(os.Args) == 1 {
		http.HandleFunc("/ydl", ydlHandler)

		fmt.Println("Starting YDL, port 8080.")

		log.Fatal(http.ListenAndServe(":8080", nil))
	} else if len(os.Args) == 2 {
		vid, err := ytdl.GetVideoInfo(os.Args[1])
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, format := range vid.Formats {
			if format.Itag == 139 ||
				format.Itag == 140 ||
				format.Itag == 141 ||
				format.Itag == 171 ||
				format.Itag == 172 {

				fname := filepath.Clean(vid.Title + "." + format.Extension)
				//download to file
				err = downloadFormat(fname, format, vid)
				fmt.Println("YDL DONE for " + fname)
				break
			}
		}
	}
}

func ydlHandler(w http.ResponseWriter, r *http.Request) {

	who := r.URL.Query().Get("who")
	if who == "" {
		fmt.Println("who query param missing")
		return
	}

	v := r.URL.Query().Get("v")
	if v == "" {
		fmt.Println("v query param missing")
		return
	}

	fmt.Println("v: " + v)
	go queueNewDL(v, who)
}

func queueNewDL(url string, who string) {

	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		ydlconf.PgHost(), ydlconf.PgPort(), ydlconf.PgUser(), ydlconf.PgPassword(), ydlconf.PgDb())

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	//query := fmt.Sprintf("INSERT INTO yqueue (yturl, status, lastupdate) VALUES('%s', 1, '%s') RETURNING id", url, time.Now().Format("2006-01-02 15:04:05.000"))
	query := "INSERT INTO yqueue (yturl, status, lastupdate) VALUES($1, 1, $2) RETURNING id"

	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
	}
	defer stmt.Close()

	var newid int
	err = stmt.QueryRow(url, time.Now().Format("2006-01-02 15:04:05.000")).Scan(&newid)
	if err != nil {
		fmt.Println(err)
	}

	for _, format := range vid.Formats {
		if format.Itag == 139 ||
			format.Itag == 140 ||
			format.Itag == 141 ||
			format.Itag == 171 ||
			format.Itag == 172 {

			fname := filepath.Clean(vid.Title + "." + format.Extension)

			//combine with "who" as dir
			wd, _ := os.Getwd()
			targetdir := filepath.Join(wd, who)
			if stat, err := os.Stat(targetdir); err == nil && stat.IsDir() {
				// path is a directory, OK
			} else {
				err = os.Mkdir(targetdir, os.ModeDir)
				if err != nil {
					fmt.Println(err)
				}
			}

			fname = filepath.Join(targetdir, fname)

			//if file exists, skip download and set status to "SKIPPED"
			if _, err := os.Stat(fname); !os.IsNotExist(err) {
				fmt.Println("File exists, skipping download.")
				query = fmt.Sprintf("UPDATE yqueue SET status = 5 WHERE id=%d", newid)
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println(err)
				}

				break
			}

			//download to file
			err = downloadFormat(fname, format, vid)

			if err != nil {
				fmt.Println("YDL ERROR for " + fname)

				//update with status ERROR
				query = fmt.Sprintf("UPDATE yqueue SET status = 4 WHERE id=%d", newid)
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println(err)
				}
			} else {
				fmt.Println("YDL DONE for " + fname)

				//update with status OK and filename
				query = fmt.Sprintf("UPDATE yqueue SET status = 3, file='%s' WHERE id=%d", fname, newid)
				_, err := db.Exec(query)
				if err != nil {
					fmt.Println(err)
				}
			}

			return
		}
	}
}

// downloads video with the specified format.
func downloadFormat(fname string, format ytdl.Format, vid *ytdl.VideoInfo) error {
	file, err := os.Create(fname)
	if err != nil {
		return err
	}

	defer file.Close()

	err = vid.Download(format, file)
	if err != nil {
		return err
	}

	return nil
}
