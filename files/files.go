package files

import (
	"os"
	"path/filepath"

	gitignore "github.com/sabhiram/go-gitignore"

	log "github.com/sirupsen/logrus"
)

func GetIgnoredList(ignorePath string, baseDir string, maxDepth int) ([]*File, error) {
	var err error
	log.Debugf("Loading dropbox ignore file: %s", ignorePath)
	gi, err := gitignore.CompileIgnoreFile(ignorePath)
	if err != nil {
		log.Errorf("ignore file not found %s", ignorePath)
		return nil, err
	}
	log.Debugf("Parsing directory from root: %s", baseDir)
	files, err := getAllFiles(baseDir)
	if err != nil {
		log.Errorf("problem reading the given directory's contents")
		return nil, err
	}
	for _, file := range files {
		file.Ignored = gi.MatchesPath(file.RelativePath)
		log.Debugf("File Name: %s || Ignored: %v\n", file.RelativePath, file.Ignored)
	}
	return files, nil
}

func getAllFiles(root string) ([]*File, error) {
	var files []*File
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		files = append(files, NewFile(path, info.IsDir(), false))
		return nil
	})
	return files, err
}
