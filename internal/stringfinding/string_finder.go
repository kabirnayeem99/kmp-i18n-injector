package stringfinding

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/kabirnayeem99/kmp-i18n-injector/internal/filescanning"
)

const yellow = "\033[33m"
const reset = "\033[0m"

func ScanKotlinFilesForMissingImports() (map[string][]string, error) {
	rootDir, err := filescanning.GetCurrentDir()

	if err != nil {
		fmt.Printf("⚠️ Error checking project root: %s\n", err)
		return nil, err
	}

	fmt.Printf("🌍 Scanning Kotlin files in project root: %s\n", rootDir)

	files, err := filescanning.FindKotlinFiles(rootDir)
	if err != nil {
		fmt.Printf("⚠️ Error finding Kotlin files.\n")
		return nil, err
	}

	missingImportsByFile := make(map[string][]string)

	stringResPackage, err := filescanning.FindStringResGeneratedPackageName(rootDir)

	if err != nil {
		fmt.Printf("⚠️ Error finding string resource package, because, %s\n", err)
		return nil, err
	}

	for _, filePath := range files {
		if strings.Contains(filepath.ToSlash(filePath), "/build/") {
			continue
		}
		missingImports, err := filescanning.CheckMissingResStringImports(filePath, stringResPackage)
		if err != nil {
			fmt.Printf("⚠️ %s\n", err)
			continue
		}

		if len(missingImports) > 0 {
			missingImportsByFile[filePath] = missingImports
			currentDir, _ := filescanning.GetCurrentDir()
			relativePath, err := filepath.Rel(currentDir, filePath)
			if err != nil {
				relativePath = filePath
			}

			fmt.Printf("\n")
			fmt.Printf("🗂️  String resource missing in the file %s%s%s:\n %s%s%s",
				yellow, relativePath, reset,
				yellow, strings.Join(missingImports, ", "), reset,
			)
			fmt.Printf("\n")
		}
	}

	return missingImportsByFile, nil
}
