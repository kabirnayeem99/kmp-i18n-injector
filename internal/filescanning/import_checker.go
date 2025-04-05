package filescanning

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func CheckMissingResStringImports(filePath, stringResPackage string) ([]string, error) {

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

	var missingImports []string

	for key, isImported := range resourceKeys {
		if isImported {
			continue
		}

		importPattern := fmt.Sprintf(`import\s+[a-zA-Z0-9._]+(?:\s+[a-zA-Z0-9._]+)*\.%s\.%s\s*`,
			regexp.QuoteMeta(stringResPackage), regexp.QuoteMeta(key))
		reImport := regexp.MustCompile(importPattern)

		if reImport.MatchString(string(fileContent)) {
			resourceKeys[key] = true
			continue
		}

		if strings.Contains(string(fileContent), fmt.Sprintf("import %s.%s", stringResPackage, key)) {
			resourceKeys[key] = true
			continue
		}

		missingImports = append(missingImports, fmt.Sprintf("import %s.%s", stringResPackage, key))
	}

	if len(missingImports) > 0 {
		return missingImports, nil
	}

	// todo: add key to the xml file
	// todo: and add the import %s.%s to the top of the file after package

	return nil, nil
}
