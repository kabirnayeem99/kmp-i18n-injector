package filescanning

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func FindStringResGeneratedPackageName(relativePath string) (string, error) {

	var foundPackage string

	re := regexp.MustCompile(`^String0\..+\.kt$`)

	err := filepath.Walk(relativePath, func(path string, info os.FileInfo, err error) error {

		if err != nil {
			return err
		}

		if info.IsDir() && (strings.Contains(path, "iosApp") || strings.Contains(path, "iosMain") || strings.Contains(path, ".git") || strings.Contains(path, ".kotlin") || strings.Contains(path, ".gradle")) {
			return filepath.SkipDir
		}

		if !info.IsDir() && re.MatchString(info.Name()) {

			fileContent, err := os.ReadFile(path)
			if err != nil {
				return fmt.Errorf("failed to read file %s: %v", path, err)
			}

			packageRe := regexp.MustCompile(`(?m)^package\s+([a-zA-Z0-9._]+)`)
			matches := packageRe.FindStringSubmatch(string(fileContent))

			if len(matches) > 1 {
				foundPackage = matches[1]
				return filepath.SkipDir
			}
		}
		return nil
	})

	if foundPackage == "" {
		return "", fmt.Errorf("no String0.<anything>.kt file with a package found in %s", relativePath)
	}

	return foundPackage, err
}
