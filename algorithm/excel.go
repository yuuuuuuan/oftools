package algorithm

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

func ExcelSumMult(sourceDirs []string, destDir string) error {
	// Merge each source folder into the destination folder
	for _, sourceDir := range sourceDirs {
		fmt.Printf("Merging folder: %s\n", sourceDir)
		err := ExcelSumSinger(sourceDir, destDir)
		if err != nil {
			log.Fatal("Error merging folder:", err)
		}
	}
	return nil
}

// Function to merge folder contents from source to destination
func ExcelSumSinger(sourceDir string, destDir string) error {
	// Walk through the source folder and copy each file/folder to the destination
	err := filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Get the relative path of the current file/folder
		relPath, _ := filepath.Rel(sourceDir, srcPath)
		destPath := filepath.Join(destDir, relPath)

		// If it's a file, copy it to the destination
		if !info.IsDir() {
			err := copyFile(srcPath, destPath)
			if err != nil {
				return err
			}
		} else {
			// If it's a directory, create the directory in the destination
			err := os.MkdirAll(destPath, os.ModePerm)
			if err != nil {
				return err
			}
		}
		return nil
	})

	return err
}

// Function to copy a single file from source to destination
func copyFile(src string, dest string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer destFile.Close()

	// Copy the content from source file to destination file
	_, err = io.Copy(destFile, srcFile)
	return err
}
