package main

import (
	"github.com/kabirnayeem99/kmp-i18n-injector/internal/kmpvalidation"
	"github.com/kabirnayeem99/kmp-i18n-injector/internal/stringfinding"
)

func main() {
	isValidKmpProject := kmpvalidation.IsValidateKmpProject()
	if !isValidKmpProject {
		return
	}
	stringfinding.ScanKotlinFilesForMissingImports()
}
