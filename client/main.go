package main

import (
	"bytes"
	"log"
	"net/http"
	"os"
	"sync"

	"github.com/cnnrznn/DolphinCloud/types"
	"github.com/fsnotify/fsnotify"
)

func main() {
	log.SetFlags(log.Lshortfile)
	log.Println("Initialized")

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	cardDir := os.Getenv("DOLPHINCLOUD_MEMCARD")
	watcher.Add(cardDir)
	log.Printf("Monitoring %v\n", cardDir)

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
	cloudAddr := os.Getenv("DOLPHINCLOUD_SERVER")

	file, err := os.Open(event.Name)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := os.ReadFile(event.Name)
	if err != nil {
		return err
	}

	gci := types.GCI{
		Name: event.Name,
		Data: data,
	}
	bs, err := gci.Marshal()
	if err != nil {
		return err
	}

	http.Post(cloudAddr, "application/json", bytes.NewReader(bs))

	return nil
}
