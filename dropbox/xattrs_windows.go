// +build windows

package dropbox

import (
	"github.com/pkg/xattr"
	log "github.com/sirupsen/logrus"
	files "github.com/willfantom/dropbox-ignore/files"
)

const (
	ignoredXattr string = "com.dropbox.ignored"
)

// TODO: this.
func osIgnoreFiles(f []*files.File) {
	log.Debugf("tagging files with ignore xattr")

	for _, file := range f {
		if err := xattr.Set(file.RelativePath, ignoredXattr, []byte("1")); err != nil {
			log.Fatal(err)
		}
	}

}