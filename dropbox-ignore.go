package main

import (
	log "github.com/sirupsen/logrus"
	dropbox "github.com/willfantom/dropbox-ignore/dropbox"
	"github.com/willfantom/dropbox-ignore/files"
)

const (
	defaultFile string = "./dpignore"
	defaultRoot string = "."
	gitignoreFile string = "./gitignore"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debugf("running dropbox ignore")
	files, err := files.GetIgnoredList(defaultFile, defaultRoot)
	if err != nil {
		log.Fatalf("could not parse directory for dropbox ignore")
	}
	dropbox.IgnoreFiles(files)
}