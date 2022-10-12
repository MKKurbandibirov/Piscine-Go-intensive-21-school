package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

func OpenFiles(old, new string) (*os.File, *os.File, error) {
	oldFile, err := os.Open(old)
	if err != nil {
		return nil, nil, err
	}
	newFile, err := os.Open(new)
	if err != nil {
		return nil, nil, err
	}
	return oldFile, newFile, nil
}

func ReadFS(fs *os.File) ([]string, error) {
	scanner := bufio.NewScanner(fs)
	var result []string
	for scanner.Scan() {
		line := scanner.Text()
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		result = append(result, line)
	}
	fs.Seek(0, 0)
	return result, nil
}

func FSContain(fs []string, s string) bool {
	for _, val := range fs {
		if val == s {
			return true
		}
	}
	return false
}

func CompareFS(old, new string) error {
	closeFS := func(oldFS, newFS *os.File) {
		oldFS.Close()
		newFS.Close()
	}
	var file []string
	var scanner *bufio.Scanner

	compareFS := func(old, new string, closeFS func(oldFS, newFS *os.File), msg string) error {
		oldFS, newFS, err := OpenFiles(old, new)
		if err != nil {
			return err
		}
		scanner = bufio.NewScanner(newFS)
		file, err = ReadFS(oldFS)
		if err != nil {
			return err
		}
		for scanner.Scan() {
			line := scanner.Text()
			if err := scanner.Err(); err != nil {
				return err
			}
			if !FSContain(file, line) {
				fmt.Printf("%s %s\n", msg, line)
			}
		}
		closeFS(oldFS, newFS)
		return nil
	}

	if err := compareFS(old, new, closeFS, "ADDED"); err != nil {
		return err
	}
	if err := compareFS(new, old, closeFS, "REMOVED"); err != nil {
		return err
	}
	return nil
}

func main() {
	oldName := flag.String("old", "", "Enter a old DB file name")
	newName := flag.String("new", "", "Enter a new DB file name")
	flag.Parse()

	if *oldName != "" || *newName != "" {
		err := CompareFS(*oldName, *newName)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
