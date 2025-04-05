package file_scanning

import (
	"errors"
	"os"
	"path/filepath"
)

func IsKMPProjectRoot(path string) (bool, error) {
	files := []string{
		"settings.gradle.kts",
		"build.gradle.kts",
		"gradlew",
		"gradle.properties",
	}
	dirs := []string{
		"composeApp",
		"iosApp",
		"gradle",
	}

	fileMatches := 0
	for _, name := range files {
		fp := filepath.Join(path, name)
		if _, err := os.Stat(fp); err == nil {
			fileMatches++
		} else if !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}

	dirMatches := 0
	for _, name := range dirs {
		dp := filepath.Join(path, name)
		info, err := os.Stat(dp)
		if err == nil && info.IsDir() {
			dirMatches++
		} else if err != nil && !errors.Is(err, os.ErrNotExist) {
			return false, err
		}
	}

	return fileMatches >= 2 && dirMatches >= 1, nil
}

func GetCurrentDir() (string, error) {
	return os.Getwd()
}

func FindKotlinFiles(rootDir string) ([]string, error) {
	var files []string

	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && filepath.Ext(info.Name()) == ".kt" {
			files = append(files, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return files, nil
}
