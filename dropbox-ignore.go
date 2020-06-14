package main

import (
	"flag"
	"fmt"

	log "github.com/sirupsen/logrus"
	dropbox "github.com/willfantom/dropbox-ignore/dropbox"
	"github.com/willfantom/dropbox-ignore/files"
)

const (
	defaultFile string = "./.dbignore"
	defaultRoot string = "."
	gitignoreFile string = "./.gitignore"
)

func main() {
	fmt.Println("Dropbox Ignore...")
	useGitIgnore := flag.Bool("g", false, "use .gitignore file rather then .dbignore")
	debugLevel := flag.Bool("d", false, "get debug level logs")
	doUpdate := flag.Bool("u", false, "also mark files to sync that don't match")
	flag.Parse()
	if *debugLevel {
		log.SetLevel(log.DebugLevel)
	}
	ignoreFile := defaultFile
	if *useGitIgnore {
		ignoreFile = gitignoreFile
	}
	files, err := files.GetIgnoredList(ignoreFile, defaultRoot)
	if err != nil {
		log.Fatalf("could not parse directory for dropbox ignore")
	}
	count := dropbox.IgnoreFiles(files)
	fmt.Printf("...Ignored %d files \n", count)
	if *doUpdate {
		updateCount := dropbox.SyncFiles(files)
		fmt.Printf("...Un-Ignored %d files \n", updateCount)
	}
	
}