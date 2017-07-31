package util

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
	"runtime/debug"

	"github.com/teo-mateo/ydl/config"
)

// FileExists tired of typing this shit all the time.
func FileExists(file string) bool {
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return false
	}
	return true
}

// SendHTTPError ...
func SendHTTPError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	debug.PrintStack()
	http.Error(w, err.Error(), http.StatusInternalServerError)
}

// Zip a bunch of files and return the path of the zip
func Zip(files []string) (string, error) {

	//check that files exist
	for _, file := range files {
		if !FileExists(file) {
			return "", errors.New("file does not exist: file")
		}
	}

	//generate zip file name

	zipFileName := fmt.Sprintf("ydl-%s.zip", time.Now().Format("20060102150405.999"))
	fmt.Println("generating temp zip file: ", zipFileName)
	tempDir, err := config.TempFolder()
	if err != nil {
		return "", err
	}

	//create temp folder if it doesn't exits
	if !FileExists(tempDir) {
		fmt.Println("creating temp dir: ", tempDir)
		err := os.Mkdir(tempDir, 0700)
		if err != nil {
			return "", err
		}
	} else {
		fmt.Println("temp dir exists: ", tempDir)
	}

	//full name of zip file (incl folder)
	zipFileName = filepath.Join(tempDir, zipFileName)
	if FileExists(zipFileName) {
		fmt.Println("zip already exists: ", zipFileName)
		return "", errors.New("zip already exists")
	} else {
		fmt.Println("full path of zip file: ", zipFileName)
	}

	//create the file
	zipFile, err := os.Create(zipFileName)
	if err != nil {
		return "", err
	}
	defer zipFile.Close()

	//create a new zip archive
	w := zip.NewWriter(zipFile)
	defer w.Close()

	for _, file := range files {

		//do this in a different func so the defers are executed at each loop
		//otherwise files remain open until all of the zip is built.
		err := func(file string) error {

			fmt.Println("zipping ", file)

			//open regular file
			f, err := os.Open(file)
			if err != nil {
				return err
			}
			defer f.Close()

			//get FileInfo
			info, err := f.Stat()
			if err != nil {
				return err
			}

			//create file info header using FileInfo
			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			//this gives us an individual writer for each file that we zip
			//so we will use io.Copy()
			writer, err := w.CreateHeader(header)
			if err != nil {
				return err
			}

			_, err = io.Copy(writer, f)
			if err != nil {
				return err
			}
			return nil
		}(file)

		if err != nil {
			return "", err
		}
	}

	fmt.Println("done.")
	return zipFileName, nil
}
