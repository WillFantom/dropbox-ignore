package dropbox

import (
	files "github.com/willfantom/dropbox-ignore/files"
)

// IgnoreFiles tells dropbox to ignore any file flagged by the gitignore style parser 
// using extended file attributes
func IgnoreFiles(f []*files.File) {
	osIgnoreFiles(f)
}
