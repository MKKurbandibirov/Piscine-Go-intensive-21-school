package handler

import (
	"archive/tar"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func OnceHandler(path string) (string, error) {
	abs, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}
	filename := strings.TrimRight(abs, filepath.Ext(path))
	target := fmt.Sprintf("%s_%d.tar", filename, time.Now().Unix())
	tarfile, err := os.Create(target)
	if err != nil {
		return "", err
	}
	defer tarfile.Close()

	tw := tar.NewWriter(tarfile)
	defer tw.Close()

	file, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer file.Close()

	info, err := os.Lstat(path)
	if err != nil {
		return "", err
	}

	//header, err := tar.FileInfoHeader(info, info.Name())
	//if err != nil {
	//	return "", err
	//}

	//if baseDir != "" {
	//	header.Name = filepath.Join(baseDir, strings.TrimPrefix(path, source))
	//}

	_, err = io.Copy(tw, file)

	return target, nil

}
