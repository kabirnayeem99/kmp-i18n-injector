package filescanning

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func CheckMissingResStringImports(filePath string) ([]string, error) {
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("❌ failed to read file %s: %v", filePath, err)
	}

	re := regexp.MustCompile(`Res\.string\.[a-zA-Z_]+`)
	matches := re.FindAllString(string(fileContent), -1)

	resourceKeys := make(map[string]bool)
	for _, match := range matches {
		key := strings.Split(match, ".")[2]
		resourceKeys[key] = false
	}

	currentDir, _ := GetCurrentDir()
	relativePath, err := filepath.Rel(currentDir, filePath)
	if err != nil {
		relativePath = filePath
	}

	stringResPackage, err := FindStringResGeneratedPackageName(relativePath)

	if err != nil {
		return nil, err
	}

	var missingImports []string

	for key := range resourceKeys {
		importRegex := fmt.Sprintf(`import %s.%s`, relativePath, key)

		importFound := false
		reImport := regexp.MustCompile(importRegex)
		if reImport.MatchString(string(fileContent)) {
			importFound = true
		}

		if !importFound {
			missingImports = append(missingImports, fmt.Sprintf("import %s.%s", stringResPackage, key))
		} else {
			resourceKeys[key] = true
		}
	}

	if len(missingImports) > 0 {
		return missingImports, nil
	}

	return nil, nil
}
