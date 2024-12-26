package algorithm

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/ini.v1"
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
			err := copyOrMergeCSV(srcPath, destPath)
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

// Function to copy or merge a CSV file
func copyOrMergeCSV(src string, dest string) error {
	// Check if destination file exists
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		// Destination file does not exist, just copy the file
		return copyFile(src, dest)
	}

	// Destination file exists, check if it's a CSV file
	if filepath.Ext(dest) == ".csv" {
		// Merge the two CSV files
		return mergeCSVFiles(src, dest)
	}

	// Destination file exists, check if it's a INI file
	if filepath.Ext(dest) == ".ini" {
		// Merge the two CSV files
		return mergeINIFiles(src, dest)
	}

	// If the file exists but is not a CSV, return an error or handle it differently
	return fmt.Errorf("destination file %s already exists and is not a CSV or INI file", dest)
}

// mergeINIFiles merges the contents of src INI file into the destination directory
// while handling duplicate filenames by appending incremental suffixes (_1, _2, etc.).
func mergeINIFiles(src string, destDir string) error {
	// Get the base filename of the source file
	srcFilename := filepath.Base(src)
	destPath := filepath.Join(destDir, srcFilename)

	// If a file with the same name exists, generate a new unique filename
	destPath = ensureUniqueFilename(destPath)

	// Load the source INI file
	srcFile, err := ini.Load(src)
	if err != nil {
		return fmt.Errorf("failed to load source INI file: %w", err)
	}

	// Save the INI file to the unique destination path
	err = srcFile.SaveTo(destPath)
	if err != nil {
		return fmt.Errorf("failed to save INI file to destination: %w", err)
	}

	fmt.Printf("Merged %s into %s\n", src, destPath)
	return nil
}

// ensureUniqueFilename checks if a file exists and appends a suffix (_1, _2, ...) until the filename is unique.
func ensureUniqueFilename(filePath string) string {
	dir := filepath.Dir(filePath)
	ext := filepath.Ext(filePath)
	base := filepath.Base(filePath[:len(filePath)-len(ext)])

	counter := 1
	newPath := filePath

	for {
		if _, err := os.Stat(newPath); os.IsNotExist(err) {
			// File does not exist, return this path
			return newPath
		}

		// File exists, generate a new filename with a counter
		newPath = filepath.Join(dir, fmt.Sprintf("%s_%d%s", base, counter, ext))
		counter++
	}
}


// Function to merge two CSV files
func mergeCSVFiles(src string, dest string) error {
	// Open the source CSV file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	// Open the destination CSV file in append mode
	destFile, err := os.OpenFile(dest, os.O_RDWR|os.O_APPEND, os.ModePerm)
	if err != nil {
		return fmt.Errorf("failed to open destination file: %w", err)
	}
	defer destFile.Close()

	// Read from the source CSV
	srcReader := csv.NewReader(srcFile)
	destReader := csv.NewReader(destFile)

	// Read the headers from both files
	srcHeader, err := srcReader.Read()
	if err != nil {
		return fmt.Errorf("failed to read source file header: %w", err)
	}

	destHeader, err := destReader.Read()
	if err != nil {
		return fmt.Errorf("failed to read destination file header: %w", err)
	}

	// Ensure headers match
	if !equalHeaders(srcHeader, destHeader) {
		return fmt.Errorf("CSV headers do not match: %v != %v", srcHeader, destHeader)
	}

	// Write the source CSV content to the destination CSV
	destWriter := csv.NewWriter(destFile)
	for {
		record, err := srcReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return fmt.Errorf("failed to read source CSV content: %w", err)
		}

		// Write the record to the destination file
		if err := destWriter.Write(record); err != nil {
			return fmt.Errorf("failed to write record to destination file: %w", err)
		}
	}
	destWriter.Flush()
	return destWriter.Error()
}

// Helper function to check if two CSV headers are equal
func equalHeaders(header1, header2 []string) bool {
	if len(header1) != len(header2) {
		return false
	}
	for i := range header1 {
		if strings.TrimSpace(header1[i]) != strings.TrimSpace(header2[i]) {
			return false
		}
	}
	return true
}

// Function to copy a single file from source to destination
func copyFile(src string, dest string) error {
	// Open the source file
	srcFile, err := os.Open(src)
	if err != nil {
		return fmt.Errorf("failed to open source file: %w", err)
	}
	defer srcFile.Close()

	// Create the destination file
	destFile, err := os.Create(dest)
	if err != nil {
		return fmt.Errorf("failed to create destination file: %w", err)
	}
	defer destFile.Close()

	// Copy the content from source file to destination file
	_, err = io.Copy(destFile, srcFile)
	return err
}
