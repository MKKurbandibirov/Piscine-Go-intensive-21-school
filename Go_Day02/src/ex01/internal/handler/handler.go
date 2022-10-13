package handler

import (
	"bufio"
	"day02/internal/domain"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"sync"
	"unicode/utf8"
)

func HandleWords(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		return -1
	}
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	var result int
	for scanner.Scan() {
		result++
	}
	return result
}

func HandleLines(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		return -1
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return -1
	}
	return len(strings.Split(string(data), "\n"))
}

func HandleSymbols(fileName string) int {
	file, err := os.Open(fileName)
	if err != nil {
		return -1
	}
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return -1
	}
	return utf8.RuneCountInString(string(data))
}

func Handle(flags domain.Flags) {
	wg := new(sync.WaitGroup)
	mu := new(sync.Mutex)
	if flags.W {
		for _, val := range flags.Files {
			var count int
			wg.Add(1)
			go func(val string) {
				defer wg.Done()
				count = HandleWords(val)
				mu.Lock()
				fmt.Printf("%d\t%s\n", count, val)
				mu.Unlock()
			}(val)
		}
	} else if flags.L {
		for _, val := range flags.Files {
			var count int
			wg.Add(1)
			go func(val string) {
				defer wg.Done()
				count = HandleLines(val)
				mu.Lock()
				fmt.Printf("%d\t%s\n", count, val)
				mu.Unlock()
			}(val)
		}
	} else if flags.M {
		for _, val := range flags.Files {
			var count int
			wg.Add(1)
			go func(val string) {
				defer wg.Done()
				count = HandleSymbols(val)
				mu.Lock()
				fmt.Printf("%d\t%s\n", count, val)
				mu.Unlock()
			}(val)
		}
	}
	wg.Wait()
}
