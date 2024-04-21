package explorer

import (
	"fmt"
	"io/fs"
	"log"
	"os"
)

func OpenDir(dir string) []fs.DirEntry {
	files, err := os.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	return files
}
