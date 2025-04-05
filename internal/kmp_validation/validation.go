package kmp_validation

import (
	"fmt"
	"log"
	"os"

	"github.com/kabirnayeem99/kmp-i18n-injector/internal/file_scanning"
)

func IsValidateKmpProject() bool {
	fmt.Printf("🌍 Welcome to the KMP i18n Injector!\n")

	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("❌ Failed to get current working directory: %s\n", err)
		return false
	}

	isRoot, err := file_scanning.IsKMPProjectRoot(wd)
	if err != nil {
		log.Fatalf("⚠️ Error checking project root: %s\n", err)
		return false
	}

	if !isRoot {
		fmt.Printf("⚠️ Not a KMP root. Please ensure you're in the root directory of a KMP project.\n")
		return false
	}

	fmt.Printf("✅ Kotlin Multiplatform project root detected! You're good to go!\n")

	return true
}
