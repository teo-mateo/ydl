package cleanup

import (
	"fmt"
	"log"
	"os"
	"time"
)

var cleanupQueue = make(chan string)

func StartCleanupRoutine() {
	//start cleanup goroutine
	go func() {
		for {
			file := <-cleanupQueue
			fmt.Printf("CLEANUP: %s\n", file)
			err := os.Remove(file)
			if err != nil {
				log.Println(err)
			}
		}
	}()
}

func MarkForCleanup(file string, delay time.Duration) {
	go func() {
		fmt.Printf("will sleep for %f seconds, then delete %s\n", delay.Seconds(), file)
		time.Sleep(delay)
		cleanupQueue <- file
	}()
}
