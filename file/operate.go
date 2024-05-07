package file

import (
	"os"
)

func SafeRename(oldpath, newpath string) error {
	err := os.Rename(oldpath, newpath)
	return err
}

func SafeRemove(path string) error {
	err := os.Remove(path)
	return err
}
