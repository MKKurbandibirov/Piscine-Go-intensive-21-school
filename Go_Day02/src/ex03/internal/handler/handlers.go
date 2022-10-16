package handler

import (
	"archive/tar"
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateArchivesDir(archName string) (string, error) {
	archName = strings.TrimPrefix(archName, "/")
	err := os.MkdirAll(archName, os.ModePerm)
	if err != nil {
		return "", err
	}
	if !strings.HasSuffix(archName, "/") {
		archName += "/"
	}
	return archName, nil
}

func createArchiveName(path string) (string, error) {
	info, err := os.Stat(path)
	if err != nil {
		return "", err
	}
	_time := info.ModTime().Unix()
	filename := strings.TrimRight(path, filepath.Ext(path))
	tmp := strings.Split(filename, "/")
	filename = tmp[len(tmp)-1]
	target := fmt.Sprintf("%s_%d.tar.gz", filename, _time)

	return target, nil
}

func OnceHandler(path string, archName string) (string, error) {
	target, err := createArchiveName(path)
	if err != nil {
		return "", err
	}
	if archName != "" {
		target = archName + target
	}

	out, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer out.Close()

	gz := gzip.NewWriter(out)
	defer gz.Close()
	tw := tar.NewWriter(gz)
	defer tw.Close()

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}

	fileInfo, err := file.Stat()
	if err != nil {
		return "", err
	}

	header, err := tar.FileInfoHeader(fileInfo, fileInfo.Name())
	if err != nil {
		return "", err
	}
	header.Name = fileInfo.Name()

	err = tw.WriteHeader(header)
	if err != nil {
		return "", err
	}

	_, err = io.Copy(tw, file)
	if err != nil {
		return "", err
	}
	return target, nil

}
