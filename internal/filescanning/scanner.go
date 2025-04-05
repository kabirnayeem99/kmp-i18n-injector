package filescanning

import (
	"errors"
	"io/fs"
	"os"
	"path/filepath"
	"strings"
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

	err := filepath.WalkDir(rootDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() && (d.Name() == "build" || d.Name() == "generated") {
			return fs.SkipDir
		}

		if !d.IsDir() && strings.EqualFold(filepath.Ext(d.Name()), ".kt") {
			files = append(files, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return files, nil
}
