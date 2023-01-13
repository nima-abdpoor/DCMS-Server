package watcher

import (
	"DCMS/db/influx"
	"DCMS/parser"
	"bufio"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"path/filepath"
	"sync"
)

var watcher *fsnotify.Watcher

func StartWatching(wg *sync.WaitGroup) {
	defer wg.Done()

	// creates a new file watcher
	watcher, _ = fsnotify.NewWatcher()
	defer watcher.Close()

	// starting at the root of the project, walk each file/directory searching for directories
	if err := filepath.Walk("/home/nima/GolandProjects/DCMS-Server/LiveLogs", watchDir); err != nil {
		fmt.Println("filepath.Walk --> ERROR: ", err)
	}
	//
	done := make(chan bool)
	//
	go func() {
		for {
			select {
			// watch for events
			case event := <-watcher.Events:
				if event.Has(fsnotify.Write) {
					//log.Println("modified file:", event.Name)
					//log.Println("modified file:", event.String())
					file, err := os.Open(event.Name)
					if err != nil {
						log.Fatalf("failed to open")
					}
					scanner := bufio.NewScanner(file)
					scanner.Split(bufio.ScanLines)
					var text []string
					for scanner.Scan() {
						text = append(text, scanner.Text())
					}
					file, err = os.Create(event.Name)
					if err != nil {
						log.Fatalf("failed to open")
					}
					file.Close()
					for _, eachLn := range text {
						influx.StartInfluxDB(parser.ParsLog(eachLn), event.Name)
					}
				}

				// watch for errors
			case err := <-watcher.Errors:
				fmt.Println("ERROR", err)
			}
		}
	}()

	<-done
}

// watchDir gets run as a walk func, searching for directories to add watchers to
func watchDir(path string, fi os.FileInfo, err error) error {

	// since fsnotify can watch all the files in a directory, watchers only need
	// to be added to each nested directory
	if fi.Mode().IsDir() {
		return watcher.Add(path)
	}
	return nil
}
