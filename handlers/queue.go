package handlers

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/kennygrant/sanitize"

	"os/exec"
	"strings"
	"net/url"

	//...
	_ "github.com/lib/pq"
	"github.com/rylio/ytdl"
	"github.com/teo-mateo/ydl/ydata"
	"errors"
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

func queueNewDL(videoUrl string, who string) error {
	url, err := url.Parse(videoUrl)
	if err != nil{
		return err
	}

	key := url.Query().Get("v")
	if key == ""{
		return errors.New(fmt.Sprintf("Could not find youtube video key in url: %s", videoUrl))
	}

	id, err := ydata.YQueueAdd(videoUrl, key, who)
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
					return
				}
			}

			fname = filepath.Join(targetdir, fname)

			//if file exists, skip download and set status to "SKIPPED"
			if _, err := os.Stat(fname); !os.IsNotExist(err) {
				fmt.Println("File exists, skipping download.")
				err = ydata.YQueueUpdate(newid, ydata.STATUSSkipped5)
				if err != nil {
					fmt.Println(err)
					return
				}

				break
			}

			//download to file
			err := downloadFormat(fname, format, vid)

			if err != nil {
				log.Println("YDL ERROR for " + fname)

				//update with status ERROR
				err = ydata.YQueueUpdate(newid, ydata.STATUSError4)
				if err != nil {
					log.Println(err)
					return
				}
			} else {
				fmt.Println("YDL Download DONE for " + fname)

				//update with status OK and filename (fname)
				err = ydata.YQueueUpdate(newid, ydata.STATUSDownloaded3, fname)
				if err != nil {
					log.Println(err)
					return
				}

				fmt.Println("Now converting to Mp3 format")

				mp3fname := strings.Replace(filepath.Base(fname), ".mp4", ".mp3", 1)
				mp3fname = filepath.Join(targetdir, mp3fname)
				if _, err := os.Stat(mp3fname); os.IsExist(err) {
					fmt.Println("Mp3 file already exists, skipping convert.")
					err = ydata.YQueueUpdate(newid, ydata.STATUSSkipped5, mp3fname)
					if err != nil {
						fmt.Println(err)
					}
				} else {
					//use ffmpeg to convert to mp3
					//ffmpegArgs := fmt.Sprintf("-i \"%s\" \"%s\"", fname, mp3fname)
					//fmt.Printf("CMD: %s %s\n", "ffmpeg", ffmpegArgs)
					cmd := exec.Command("ffmpeg", "-i", fname, mp3fname)

					fmt.Printf("ffmpeg -i \"%s\" \"%s\" \n", fname, mp3fname)

					err := cmd.Start()
					if err != nil {
						fmt.Println("Could not start ffmpeg process.")
					}

					err = cmd.Wait()
					if err != nil {
						fmt.Println(err)
						return
					}

					fmt.Println("Done converting: ", mp3fname)

					//check: mp3 file must exist now
					if _, err := os.Stat(mp3fname); os.IsNotExist(err) {
						fmt.Println("Something went wrong during conversion: Mp3 file does not exist.")
						return
					}

					//update db
					err = ydata.YQueueUpdate(newid, ydata.STATUSDownloaded3, mp3fname)
					if err != nil {
						log.Println(err)
						return
					}

					//delete mp4 file
					fmt.Println("Deleting: ", fname)
					os.Remove(fname)

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
