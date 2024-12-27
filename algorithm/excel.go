package algorithm

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func ExcelClear(sourceDir string, destDir string) error {
	// Get the current timestamp
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	destDir = filepath.Join(destDir, currentTime)

	// Ensure the destination directory exists
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create destination directory: %w", err)
	}
	if err := copyFile(sourceDir, destDir); err != nil {
		return fmt.Errorf("failed to copy destination directory: %w", err)
	}
	return removeFiles(sourceDir)
}

//func ExcelSumSelf

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

	// If dest exists, merge if both are INI files
	if filepath.Ext(src) == ".ini" && filepath.Ext(dest) == ".ini" {
		return mergeINIFiles(src, dest)
	}

	// If dest exists, merge if both are XLS files
	if filepath.Ext(src) == ".xls" && filepath.Ext(dest) == ".xls" {
		return nil
	}

	// For non-CSV files, return an error or handle differently
	return fmt.Errorf("cannot merge files: %s and %s", src, dest)
}

// Function to merge INI files
func mergeINIFiles(src string, dest string) error {
	// Extract the directory and base name from dest
	destDir := filepath.Dir(dest)
	baseName := strings.TrimSuffix(filepath.Base(dest), ".ini")

	// Determine a unique filename in the destination directory
	newDest := dest
	counter := 1
	for {
		// Check if the file exists
		if _, err := os.Stat(newDest); os.IsNotExist(err) {
			break // File does not exist, we can use this name
		}
		// Increment the counter and try a new name
		newDest = filepath.Join(destDir, fmt.Sprintf("%s_%d.ini", baseName, counter))
		counter++
	}

	// Copy the source file to the new destination
	return copyFile(src, newDest)
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
		//if record
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

func removeFiles(path string) error {
	// Read the directory contents
	entries, err := os.ReadDir(path)
	if err != nil {
		return fmt.Errorf("failed to read directory: %v", err)
	}

	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()

		if entry.IsDir() {
			// If it's a directory, remove it recursively
			err := os.RemoveAll(entryPath)
			if err != nil {
				return fmt.Errorf("failed to delete directory (%s): %v", entryPath, err)
			}
			fmt.Printf("Deleted directory: %s\n", entryPath)
		} else {
			// If it's a file, delete it directly
			err := os.Remove(entryPath)
			if err != nil {
				return fmt.Errorf("failed to delete file (%s): %v", entryPath, err)
			}
			fmt.Printf("Deleted file: %s\n", entryPath)
		}
	}

	return nil
}
