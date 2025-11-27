package msp

import (
	"os"
	"path/filepath"
	"slices"
)

func IsFile(path string) bool {
	if path == "" {
		return false
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return !fileInfo.IsDir()
}

func IsDir(path string) bool {
	if path == "" {
		return false
	}
	fileInfo, err := os.Stat(path)
	if err != nil {
		return false
	}
	return fileInfo.IsDir()
}

func GetFilePathsInDirWithExts(dirPath string, exts []string) ([]string, error) {
	ret := make([]string, 0)

	err := filepath.Walk(dirPath,
		func(filePath string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(filePath, info.Size())
			//fmt.Println(filePath)
			if IsFile(filePath) && slices.Contains(exts, filepath.Ext(filePath)) {
				ret = append(ret, filePath)
			}
			return nil
		})
	if err != nil {
		return ret, err
	}
	return ret, nil
}
