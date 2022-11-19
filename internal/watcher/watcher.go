package watcher

import (
	"fmt"
	"log"

	"github.com/fsnotify/fsnotify"
)

// Watcher watch brightness file to change and executes any function on callback
type Watcher struct {
	FilePath string
	Callback func()
}

func (w *Watcher) Watch() error {
	// Create new watcher.
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Start listening for events.
	done := make(chan error)
	defer close(done)
	go func() {
		for {
			select {
			case _, ok := <-watcher.Events:
				if !ok {
					done <- fmt.Errorf("event channel returns not ok status")
					return
				}
				w.Callback()
			case err, ok := <-watcher.Errors:
				if !ok {
					done <- fmt.Errorf("errors channel returns not ok status")
					return
				}
				done <- fmt.Errorf("got error from watcher: %w", err)
				return
			}
		}
	}()

	// Add a path.
	err = watcher.Add(w.FilePath)
	if err != nil {
		log.Fatal(err)
	}

	// Block main goroutine forever.
	return <-done
}
