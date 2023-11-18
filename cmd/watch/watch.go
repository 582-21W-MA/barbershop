package watch

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"github.com/582-21W-MA/barbershop/cmd/barbershop"
	"github.com/582-21W-MA/barbershop/cmd/serve"
	"github.com/fsnotify/fsnotify"
)

func Run(rootDir string) {
	// barbershop needs to run once to set OutputDir.
	// TODO: Remove OutputDir from barbershop.
	barbershop.Quiet = true
	if err := barbershop.Run(rootDir); err != nil {
		log.Fatalf("Error %v", err)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	// Walk root directory recursively and add each directory to the
	// monitoring list, including the root directory.
	err = filepath.WalkDir(rootDir, func(entryPath string, entry fs.DirEntry, err error) error {
		if !entry.IsDir() {
			return nil
		}
		if err := watcher.Add(entryPath); err != nil {
			log.Fatalf("Error adding watch entry: %v", err)
		}
		return nil
	})

	go watchLoop(watcher, rootDir)
	fmt.Printf("Watching %q\n", rootDir)

	serve.Run(barbershop.OutputDir)

	// Block main goroutine forever.
	<-make(chan struct{})
}

// watchLoop monitors directories in the watcher's monitoring list.
func watchLoop(watcher *fsnotify.Watcher, rootDir string) {
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				return
			}
			if err := triageEvent(event, watcher, rootDir); err != nil {
				log.Fatalf("Error triaging event: %v", err)
			}
		case err, ok := <-watcher.Errors:
			if !ok {
				return
			}
			log.Fatalf("Error: %v", err)
		}
	}
}

// triageEvent triages event based on their type (chmod, create,
// delete, rename, modify). barbershop is called when a file is either
// created, written to, removed or renamed. When a folder is created,
// it is added to the monitoring list.
func triageEvent(event fsnotify.Event, watcher *fsnotify.Watcher, rootDir string) error {
	// See fsnotify documentation. Chmod is often fired.
	if event.Has(fsnotify.Chmod) {
		return nil
	}

	// We check for create events in order to add new directories to
	// the watchlist. os.Stat can't be called on remove or rename
	// events because event.Name then points to a file that does not
	// exist anymore.
	if event.Has(fsnotify.Create) {
		if fi, err := os.Stat(event.Name); err != nil {
			return fmt.Errorf("checking if %q is a directory: %w",
				event.Name, err)
		} else if fi.IsDir() {
			if err := watcher.Add(event.Name); err != nil {
				return fmt.Errorf("adding file to watcher: %w", err)
			}
			return nil
		}
		if err := barbershop.Run(rootDir); err != nil {
			return fmt.Errorf("executing barbershop: %w", err)
		}
	}

	// barbershop is executed for rename, modify and delete events.
	if err := barbershop.Run(rootDir); err != nil {
		return fmt.Errorf("executing barbershop: %w", err)
	}
	return nil
}
