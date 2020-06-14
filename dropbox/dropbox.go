package dropbox

import (
	log "github.com/sirupsen/logrus"
	files "github.com/willfantom/dropbox-ignore/files"
)

// IgnoreFiles tells dropbox to ignore any file flagged by the gitignore style parser 
// using extended file attributes
func IgnoreFiles(f []*files.File) int {
	var err error
	count := 0
	for _, file := range f {
		if file.Ignored {
			err = osIgnoreFile(file.RelativePath)
			if err != nil {
				log.Errorf("😢 file could not be marked as ignored | %s", file.RelativePath)
			} else {
				count ++
			}
		}
	}
	return count	
}

func SyncFiles(f []*files.File) int {
	var err error
	count := 0
	for _, file := range f {
		if !file.Ignored {
			err = osUnIgnoreFile(file.RelativePath)
			if err != nil {
				log.Errorf("😢 file could not be marked to sync | %s", file.RelativePath)
			} else {
				count++
			}
		}
	}
	return count	
}
