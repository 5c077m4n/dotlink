package pathfinder

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func PathFind(args cli.Args) ([]string, error) {
	dirs := []string{}
	if args.Present() {
		for _, path := range args.Slice() {
			if absPath, error := filepath.Abs(path); error == nil {
				dirs = append(dirs, absPath)
			} else {
				log.Printf("Could not find file %s", path)
			}
		}
	} else {
		dirList, error := os.ReadDir("./")
		if error != nil {
			return nil, fmt.Errorf("could not read current directory content: %s", error.Error())
		}

		for _, dir := range dirList {
			absPath, error := filepath.Abs(dir.Name())
			if error != nil {
				fmt.Printf("Could not get a full file path: %s", error.Error())
				continue
			}

			dirs = append(dirs, absPath)
		}
	}

	return dirs, nil
}
