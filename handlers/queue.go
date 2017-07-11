package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kennygrant/sanitize"

	//I do this because I want this. fuck your motherfucking warnings.
	_ "github.com/lib/pq"
	"github.com/rylio/ytdl"
	"github.com/teo-mateo/ydl/ydata"
)

// QueueHandler ...
func QueueHandler(w http.ResponseWriter, r *http.Request) {
	who := r.URL.Query().Get("who")
	if who == "" {
		log.Println("who query param missing")
		return
	}

	v := r.URL.Query().Get("v")
	if v == "" {
		fmt.Println("v query param missing")
		return
	}

	fmt.Println("v: " + v)
	queueNewDL(v, who)
}

func queueNewDL(url string, who string) error {

	id, err := ydata.YQueueAdd(url)
	if err != nil {
		fmt.Println(err)
	}

	vid, err := ytdl.GetVideoInfo(url)
	if err != nil {
		fmt.Println(err)
		return err
	}

	//trigger download on a different routine
	go downloadVid(vid, who, id)

	return nil
}

func downloadVid(vid *ytdl.VideoInfo, who string, newid int) {

	for _, format := range vid.Formats {
		if format.Itag == 139 ||
			format.Itag == 140 ||
			format.Itag == 141 ||
			format.Itag == 171 ||
			format.Itag == 172 {

			fname := sanitize.BaseName(vid.Title) + "." + format.Extension

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
				err = ydata.YQueueUpdate(newid, 5)
				if err != nil {
					fmt.Println(err)
				}

				break
			}

			//download to file
			err := downloadFormat(fname, format, vid)

			if err != nil {
				log.Println("YDL ERROR for " + fname)

				//update with status ERROR
				err = ydata.YQueueUpdate(newid, 4)
				if err != nil {
					log.Println(err)
				}
			} else {
				fmt.Println("YDL DONE for " + fname)

				//update with status OK and filename (fname)
				err = ydata.YQueueUpdate(newid, 3, fname)
				if err != nil {
					log.Println(err)
				}
			}

			return
		}
	}
}

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
