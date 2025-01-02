package algorithm

import (
	"fmt"
	"os"
	"path/filepath"
)

func ProgramFirewareSingle(source string, destdirs []string) error {
	root := source + "\\OIS"
	targets := []string{"A", "B", "C", "D"}
	// Call the function and retrieve matching paths
	paths, err := findSpecificPaths(root, targets)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return nil
}

// FindSpecificPaths searches for files or directories named A, B, or C in the specified directory
func findSpecificPaths(root string, targets []string) ([]string, error) {
	var result []string
	targetSet := make(map[string]struct{})

	// Create a set of target names for quick lookup
	for _, target := range targets {
		targetSet[target] = struct{}{}
	}

	// Walk through the directory and its subdirectories
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		// Check if the current file or directory name matches the targets
		if _, found := targetSet[info.Name()]; found {
			result = append(result, path)
		}
		return nil
	})

	return result, err
}
