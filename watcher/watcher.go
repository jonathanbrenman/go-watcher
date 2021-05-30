package watcher

import (
	"fmt"
	"go-watcher/models"
	"io/ioutil"
	"log"
	"os"
	"time"
	"path/filepath"
)

type Watcher interface {
	Add(root string)
	List(root string)
}

type watcher struct {
	root string
	files []models.FilesMetadata
	delay time.Duration
	debug bool
	filters []string
	history []models.FilesMetadata
}

// Add new watcher
func (w *watcher) Add(root string) {
	w.root = root
	for {
		start := time.Now()
		if len(w.files) > 0 {
			w.history = w.files
		}
		w.files = []models.FilesMetadata{}
		w.files = w.listRecursive(root)
		w.logDebugMessage(fmt.Sprintf(
			"Start watching folder %s with configs: debug = %v, delay = %s, only files with extensions %s",
			root,
			w.debug,
			w.delay,
			w.filters,
		))
		for _, file := range w.files {
			found := false
			for _, hfile := range w.history {
				if file.Path == hfile.Path {
					found = true
					break
				}
			}
			if !found {
				fmt.Println("-----------------------------------------")
				log.Println("New file detected!")
				log.Println("Path: ", file.Path)
				log.Println("Name: ", file.FileInfo.Name())
				log.Println("Size: ", w.toHumanFormat(file.FileInfo.Size()))
				log.Println("Modified at: ", file.FileInfo.ModTime().Format("2006-01-02 15:04:05"))
				fmt.Println("-----------------------------------------")
			}
		}
		duration := time.Since(start)
		w.logDebugMessage("checking for new files ->", duration.String())
		time.Sleep(w.delay)
	}
}

// List files in directory recursive
func (w *watcher) List(root string) {
	for _, file := range w.listRecursive(root) {
		log.Println("Path: ", file.Path)
		log.Println("Name: ", file.FileInfo.Name())
		log.Println("Size: ", w.toHumanFormat(file.FileInfo.Size()))
		log.Println("Modified at: ", file.FileInfo.ModTime().Format("2006-01-02 15:04:05"))
		fmt.Println()
	}
}

// List nested files
func (w *watcher) listRecursive(watchFolder string) []models.FilesMetadata {
	    folders, err := ioutil.ReadDir(watchFolder)
		if err != nil {
			log.Fatal("ERROR", err)
		}

	    for _, folder := range folders {
	    	path := fmt.Sprintf("%s/%s", watchFolder, folder.Name())
			 if !folder.IsDir() {
			 	if w.checkExt(path) {
			 		fileInfo, err := os.Lstat(path)
			 		if err != nil {
			 			fmt.Println(err)
					} else {
						var fileMetadata models.FilesMetadata
						fileMetadata.FileInfo = fileInfo
						fileMetadata.Path = path
						w.files = append(w.files, fileMetadata)
					}
				}
			 } else {
			 	w.listRecursive(path)
			 }
		}
		return w.files
}

// Check file extension with the selected filters.
func (w *watcher) checkExt(filename string) bool {
	if len(w.filters) == 1 {
		return true
	}

	for _, filter := range w.filters {
		if fmt.Sprintf(".%s", filter) == filepath.Ext(filename) {
			return true
		}
	}
	return false
}

// Format byte size to human readable format
func (w *watcher) toHumanFormat(b int64) string {
        const unit = 1000
        if b < unit {
                return fmt.Sprintf("%d B", b)
        }
        div, exp := int64(unit), 0
        for n := b / unit; n >= unit; n /= unit {
                div *= unit
                exp++
        }
        return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "kMGTPE"[exp])
}

func (w *watcher) logDebugMessage(message ...string) {
	if w.debug {
		log.Println(message)
	}
}