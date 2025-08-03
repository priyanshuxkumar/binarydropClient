package watcher

import (
	"log"
	"path/filepath"
	"strings"

	"github.com/fsnotify/fsnotify"
)

type FileEventHandler interface {
	OnCreate(path string)
	OnWrite(path string)
	OnRemove(path string)
	OnRename(path string)
}

func shouldIgnoreFile(filename string) bool {
	base := filepath.Base(filename)

	// Skip hidden files
	if strings.HasPrefix(base, ".") {
		return true
	}
	return false
}

func Watcher(dir string, handler FileEventHandler) error {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	defer watcher.Close()

	err = watcher.Add(dir)
	if err != nil {
		log.Fatal("error watching dir:", err)
	}

	log.Println("Watching directory:", dir)

	for {
		select {
		case event := <-watcher.Events:
			if shouldIgnoreFile(event.Name) {
				continue
			}

			switch {
			case event.Op&fsnotify.Create == fsnotify.Create:
				handler.OnCreate(event.Name)
			case event.Op&fsnotify.Write == fsnotify.Write:
				handler.OnWrite(event.Name)
			case event.Op&fsnotify.Remove == fsnotify.Remove:
				handler.OnRemove(event.Name)
			case event.Op&fsnotify.Rename == fsnotify.Rename:
				handler.OnRename(event.Name)
			default:
				continue
			}
		case err := <-watcher.Errors:
			log.Println("watch error:", err)
		}
	}
}
