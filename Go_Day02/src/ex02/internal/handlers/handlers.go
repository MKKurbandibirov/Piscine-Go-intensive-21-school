package handlers

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
)

func GetArgs() []string {
	var args []string

	args = append(args, os.Args[2:]...)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		args = append(args, scanner.Text())
	}
	return args
}

func Handler(command string, args []string) (string, error) {
	cmd := exec.Command(command, args...)

	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return "", err
	}
	return out.String(), nil
}
