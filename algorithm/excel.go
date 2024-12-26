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

// mergeINIFiles merges the contents of the source INI file into the destination INI file.
// If there are duplicate keys, the keys from the source file are renamed with _1, _2, etc.
func mergeINIFiles(src, dest string) error {
	// Load the source INI file
	srcCfg, err := ini.Load(src)
	if err != nil {
		return fmt.Errorf("failed to load source INI file: %w", err)
	}

	// Load or create the destination INI file
	destCfg, err := ini.Load(dest)
	if err != nil {
		// If the destination file doesn't exist, create a new INI object
		if os.IsNotExist(err) {
			destCfg = ini.Empty()
		} else {
			return fmt.Errorf("failed to load destination INI file: %w", err)
		}
	}

	// Iterate through all sections and keys in the source file
	for _, srcSection := range srcCfg.Sections() {
		destSection, err := destCfg.GetSection(srcSection.Name())
		if err != nil {
			// If the section doesn't exist in the destination, create it
			destSection, _ = destCfg.NewSection(srcSection.Name())
		}

		// Iterate through all keys in the source section
		for _, srcKey := range srcSection.Keys() {
			originalKey := srcKey.Name()
			newKey := originalKey
			counter := 1

			// Ensure the key doesn't already exist in the destination section
			for destSection.HasKey(newKey) {
				newKey = fmt.Sprintf("%s_%d", originalKey, counter)
				counter++
			}

			// Add the key-value pair to the destination section
			_, err := destSection.NewKey(newKey, srcKey.Value())
			if err != nil {
				return fmt.Errorf("failed to add key %s to section %s: %w", newKey, srcSection.Name(), err)
			}
		}
	}

	// Save the merged configuration back to the destination file
	err = destCfg.SaveTo(dest)
	if err != nil {
		return fmt.Errorf("failed to save destination INI file: %w", err)
	}

	return nil
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
