package linker

import (
	"errors"
	"log"
	"os"
	"path/filepath"
)

func Link(dirs []string, target string, replace bool) error {
	os.MkdirAll(target, 0755)

	for _, dir := range dirs {
		fullTargetPath := filepath.Join(target, filepath.Base(dir))
		if replace {
			os.RemoveAll(fullTargetPath)
		}

		if error := os.Symlink(dir, fullTargetPath); error == nil {
			log.Printf("Linked %s => %s", dir, fullTargetPath)
		} else if errors.Is(error, os.ErrExist) {
			log.Printf("A symlink for %s already exists (%s)", dir, fullTargetPath)
		} else {
			log.Printf("There was an error in creating a symlink: %s", error.Error())
		}
	}

	return nil
}
