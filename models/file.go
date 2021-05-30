package models

import "os"

type FilesMetadata struct {
	Path string
	FileInfo os.FileInfo
}
