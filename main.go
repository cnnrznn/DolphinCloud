package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"sync"

	"github.com/fsnotify/fsnotify"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Initialized")

	rootDir := os.Getenv("DOLPHIN_EMU_USERPATH")
	saveDir := fmt.Sprintf("%v/StateSaves", rootDir)
	GCDir := fmt.Sprintf("%v/GC", rootDir)

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	watcher.Add(saveDir)
	watcher.Add(GCDir)
	watcher.Add(fmt.Sprintf("%v/USA/Card A", GCDir))

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go handleWatcher(watcher, wg)

	wg.Wait()
}

func handleWatcher(watcher *fsnotify.Watcher, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		select {
		case event, ok := <-watcher.Events:
			log.Println(event, ok)
			if ok {
				if err := handleEvent(event); err != nil {
					log.Println(err)
				}
			}
		case err, ok := <-watcher.Errors:
			log.Println(err, ok)
		}
	}
}

func handleEvent(event fsnotify.Event) error {
	if event.Op != fsnotify.Create {
		return nil
	}

	if strings.HasSuffix(event.Name, ".tmp") {
		return nil
	}

	return nil
}
