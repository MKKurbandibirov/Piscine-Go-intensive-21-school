package main

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"

	"day02/internal/domain"
)

func Exec(fl domain.Flags) []string {
	var files []string
	err := filepath.Walk(fl.RootName, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !fl.Dir && !fl.SymLink && !fl.File && info.Mode()&os.ModePerm != 0 {
			files = append(files, path)
			return nil
		}
		if fl.Dir && info.Mode()&os.ModeDir != 0 && info.Mode()&os.ModePerm != 0 {
			files = append(files, path)
			return nil
		}
		if fl.SymLink && info.Mode()&os.ModeSymlink != 0 && info.Mode()&os.ModePerm != 0 {
			files = append(files, path)
			return nil
		}
		if fl.File && info.Mode()&os.ModeDir == 0 && info.Mode()&os.ModeSymlink == 0 && info.Mode()&os.ModePerm != 0 {
			if fl.Ext != "" && filepath.Ext(path) == "."+fl.Ext {
				files = append(files, path)
				return nil
			} else if fl.Ext == "" {
				files = append(files, path)
				return nil
			}
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return files
}

func main() {
	flags := domain.NewFlags()
	var files = Exec(*flags)

	for _, file := range files {
		fmt.Println(file)
	}
}
