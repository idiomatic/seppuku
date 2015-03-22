package seppuku

import (
	"log"
	"os"
	"path/filepath"

	"github.com/go-fsnotify/fsnotify"
)

func Seppuku(watching []string) {
	if len(watching) == 0 {
		return
	}
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	for _, pattern := range watching {
		matches, err := filepath.Glob(pattern)
		if err != nil {
			log.Fatal(err)
		}
		for _, globbed := range matches {
			watcher.Add(globbed)
		}
	}

	for {
		select {
		case event := <-watcher.Events:
			if event.Op&fsnotify.Write == fsnotify.Write {
				os.Exit(0)
			} else if event.Op&fsnotify.Create == fsnotify.Create {
				os.Exit(0)
			} else if event.Op&fsnotify.Remove == fsnotify.Remove {
				os.Exit(0)
			} else {
				//log.Printf("event %v\n", event)
			}
		}
	}
}
