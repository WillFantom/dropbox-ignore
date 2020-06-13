package main

import (
	"fmt"
	"os"
	"path/filepath"

	gitignore "github.com/monochromegane/go-gitignore"
)

func main() {
	var err error
	gi, err := gitignore.NewGitIgnore("./.gitignore", ".")
	if err != nil {
		panic(err)
	}
	paths, err := FilePathWalkDir(".")
	if err != nil {
		panic(err)
	}
	for _, path := range paths {
		isAccepted := gi.Match(path, IsDirectory(path))
		fmt.Printf("%s -> %v\n", path, isAccepted)
	}
	fmt.Printf("%v", gi)
}

func FilePathWalkDir(root string) ([]string, error) {
	var files []string
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if !info.IsDir() {
					files = append(files, path)
			}
			return nil
	})
	return files, err
}

func IsDirectory(path string) (bool) {
	fileInfo, err := os.Stat(path)
	if err != nil{
		panic(err)
	}
	return fileInfo.IsDir()
}