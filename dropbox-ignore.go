package main

import (
	log "github.com/sirupsen/logrus"
	dropbox "github.com/willfantom/dropbox-ignore/dropbox"
	"github.com/willfantom/dropbox-ignore/files"
)

func main() {
	log.SetLevel(log.DebugLevel)
	log.Debugf("running dropbox ignore")
	files, err := files.GetIgnoredList(".dpignore", ".", 1)
	if err != nil {
		log.Fatalf("could not parse directory for dropbox ignore")
	}
	dropbox.IgnoreFiles(files)
}