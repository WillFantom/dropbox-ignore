// +build !windows

package dropbox

import (
	"github.com/pkg/xattr"
)

const (
	ignoredXattr string = "com.dropbox.ignored"
)

func osIgnoreFile(filePath string) error {
	if err := xattr.Set(filePath, ignoredXattr, []byte("1")); err != nil {
		return err
	}
	return nil
}

func osUnIgnoreFile(filePath string) error {
	if err := xattr.Remove(filePath, ignoredXattr); err != nil {
		return err
	}
	return nil
}