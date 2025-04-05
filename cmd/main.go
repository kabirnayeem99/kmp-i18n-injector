package main

import (
	"github.com/kabirnayeem99/kmp-i18n-injector/internal/kmp_validation"
	"github.com/kabirnayeem99/kmp-i18n-injector/internal/string_finding"
)

func main() {
	isValidKmpProject := kmp_validation.IsValidateKmpProject()
	if !isValidKmpProject {
		return
	}
	string_finding.FindStringResources()
}
