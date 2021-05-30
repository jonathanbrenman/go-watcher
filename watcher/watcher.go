package watcher

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
	"path/filepath"
)

type Watcher interface {
	Add(root string)
}

type watcher struct {
	root string
	files []string
	delay time.Duration
	debug bool
	filters []string
}

// Add new watcher
func (w *watcher) Add(root string) {
	w.root = root
	for {
		w.files = []string{}
		w.files = w.listRecursive(root)
		if w.debug {
			log.Println(fmt.Sprintf(
				"Start watching folder %s with configs: debug = %v, delay = %s, only files with extensions %s",
				root,
				w.debug,
				w.delay,
				w.filters,
			))
			log.Printf("Files -> %v\n", w.files)
		}
		time.Sleep(w.delay)
	}
}

// List nested files
func (w *watcher) listRecursive(watchFolder string) []string {
	    folders, err := ioutil.ReadDir(watchFolder)
		if err != nil {
			log.Fatal("ERROR", err)
		}

	    for _, folder := range folders {
	    	path := fmt.Sprintf("%s/%s", watchFolder, folder.Name())
			 if !folder.IsDir() {
			 	if w.checkExt(path) {
			 		w.files = append(w.files, path)
				}
			 } else {
			 	w.listRecursive(path)
			 }
		}
		return w.files
}

// Check file extension with the selected filters.
func (w *watcher) checkExt(filename string) bool {
	if len(w.filters) == 0 {
		return true
	}

	for _, filter := range w.filters {
		if fmt.Sprintf(".%s", filter) == filepath.Ext(filename) {
			return true
		}
	}
	return false
}