package string_finding

import (
	"fmt"
	"os"
	"regexp"

	"github.com/kabirnayeem99/kmp-i18n-injector/internal/file_scanning"
)

// FindStringResources scans Kotlin files for Res.string instances.
func FindStringResources() {
	rootDir, err := file_scanning.GetCurrentDir()
	if err != nil {
		fmt.Printf("⚠️ Error checking project root: %s\n", err)
		return
	}

	fmt.Printf("🌍 Scanning Kotlin files in project root: %s\n", rootDir)

	files, err := file_scanning.FindKotlinFiles(rootDir)
	if err != nil {
		fmt.Printf("⚠️ Error finding Kotlin files.\n")
		return
	}

	instancesByFile := make(map[string][]string)

	for _, filePath := range files {
		instances, err := findResStringInstances(filePath)
		if err != nil {
			fmt.Printf("⚠️ Error reading file %s %s\n", filePath, err)
			continue
		}

		if len(instances) > 0 {
			instancesByFile[filePath] = instances
		}
	}

	if len(instancesByFile) > 0 {
		fmt.Println("Found Res.string instances:")
		for filePath, instances := range instancesByFile {
			fmt.Printf("File: %s\n", filePath)
			for _, instance := range instances {
				fmt.Printf("  - %s\n", instance)
			}
		}
	} else {
		fmt.Println("No Res.string instances found.")
	}
}

// findResStringInstances finds instances of Res.string in a Kotlin file.
func findResStringInstances(filePath string) ([]string, error) {
	var instances []string
	file, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	re := regexp.MustCompile(`Res\.string\.[a-zA-Z_]+`)

	matches := re.FindAllString(string(file), -1)
	instances = append(instances, matches...)

	return instances, nil
}
