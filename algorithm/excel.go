package algorithm

import (
	"encoding/csv"
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
			err := copyOrMerge(srcPath, destPath)
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

// Function to copy or merge files based on their existence
func copyOrMerge(src string, dest string) error {
	// Check if dest file exists
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		// If dest does not exist, simply copy the file
		return copyFile(src, dest)
	}

	// If dest exists, merge if both are CSV files
	if filepath.Ext(src) == ".csv" && filepath.Ext(dest) == ".csv" {
		return mergeCSVFiles(src, dest)
	}

	// For non-CSV files, return an error or handle differently
	return fmt.Errorf("cannot merge non-CSV files: %s and %s", src, dest)
}

// Function to copy a single file from src to dest
func copyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return fmt.Errorf("failed to copy file: %w", err)
	}

	fmt.Printf("Copied %s to %s\n", src, dest)
	return nil
}

// Function to merge two CSV files (entire content, including headers)
func mergeCSVFiles(src string, dest string) error {
	// Open source and destination files
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	defer destFile.Close()

	// Create readers and writer
	srcReader := csv.NewReader(srcFile)
	destWriter := csv.NewWriter(destFile)

	// Append all rows from source to destination
	for {
		record, err := srcReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf("Skipping invalid row: %v\n", err)
			continue // Skip invalid rows
		}

		err = destWriter.Write(record)
		if err != nil {
			return fmt.Errorf("failed to write record to destination: %w", err)
		}
	}

	destWriter.Flush()
	if err := destWriter.Error(); err != nil {
		return fmt.Errorf("failed to flush writer: %w", err)
	}

	fmt.Printf("Merged %s into %s\n", src, dest)
	return nil
}
