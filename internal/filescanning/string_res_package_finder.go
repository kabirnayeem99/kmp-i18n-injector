package filescanning

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindStringResGeneratedPackageName(rootDir string) (string, error) {

	var foundPackage string
	err := filepath.Walk(rootDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.IsDir() && !strings.Contains(path, "build") {
			return nil
		}

		if strings.HasSuffix(info.Name(), "String0.something.kt") {

			fileContent, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %v", path, err)
			}

			re := regexp.MustCompile(`(?m)^package\s+([a-zA-Z0-9._]+)`)
			matches := re.FindStringSubmatch(string(fileContent))

			if len(matches) > 1 {
				foundPackage = matches[1]
				return filepath.SkipDir
			}
		}
		return nil
	})

	if foundPackage == "" {
		return "", fmt.Errorf("no String0.something.kt file with a package found")
	}

	return foundPackage, err
}
