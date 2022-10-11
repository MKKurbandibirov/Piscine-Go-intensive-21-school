package main

import (
	"bufio"
	"day00/mean"
	"day00/median"
	"day00/mode"
	"day00/sd"
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strconv"
)

const (
	MAX = 100000
	MIN = -100000
)

func ReadSequence(reader io.Reader) ([]int, error) {
	var input = make([]int, 0)
	scanner := bufio.NewScanner(reader)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		text := scanner.Text()
		err := scanner.Err()
		if err == io.EOF {
			if len(input) == 0 {
				return nil, errors.New("empty string")
			}
			break
		}
		if err != nil {
			return nil, err
		}
		n, err := strconv.Atoi(text)
		if err != nil {
			return nil, errors.New("letter character")
		}
		if n > MAX || n < MIN {
			return nil, errors.New("number out of bounds")
		}
		input = append(input, n)
	}
	return input, nil
}

func main() {
	arr, err := ReadSequence(os.Stdin)
	if err != nil {
		log.Fatal("reading err: %w", err)
	}
	sort.Ints(arr)

	mean := mean.ComputeMean(arr)
	fmt.Printf("Mean: %.2f\n", mean)
	median := median.ComputeMedian(arr)
	fmt.Printf("Median: %.2f\n", median)
	mode := mode.ComputeMode(arr)
	fmt.Printf("Mode: %d\n", mode)
	sd := sd.ComputeSD(arr, mean)
	fmt.Printf("SD: %.2f\n", sd)
}
