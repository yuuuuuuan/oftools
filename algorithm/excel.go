package algorithm

import (
	"encoding/csv"
	"fmt"
	"io"
	"oftools/oflog"
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
	if err := createDir(destDir); err != nil {
		oflog.Print.Errorf("%s Error:failed at algorithm.createDir!", getFunctionName())
		return err
	}
	if err := copyFile(sourceDir, destDir); err != nil {
		oflog.Print.Errorf("%s Error:failed at algorithm.copyFile!", getFunctionName())
		return err
	}
	return removeFiles(sourceDir)
}

func ExcelSumMult(sourceDirs []string, destDir string) error {
	// Merge each source folder into the destination folder
	for _, sourceDir := range sourceDirs {
		oflog.Print.Infof("%s:Merging folder: %s", getFunctionName(), sourceDir)
		err := ExcelSumSinger(sourceDir, destDir)
		if err != nil {
			oflog.Print.Fatalf("%s Error:failed at algorithm.ExcelSumSinger!", getFunctionName())
			return err
		}
	}
	return ExcelSumSelf(destDir)
}

// ExcelSumSelf organizes `.csv` and `.txt` files into separate folders `sumcsv` and `sumtxt`.
func ExcelSumSelf(sourceDir string) error {
	oflog.Print.Debugf("%s ============> Debug", getFunctionName())
	// Define the target directories
	sumCSVDir := filepath.Join(sourceDir, ".sumcsv")
	sumTXTDir := filepath.Join(sourceDir, ".sumtxt")

	// Create target directories if they don't exist
	if err := os.MkdirAll(sumCSVDir, 0755); err != nil {
		oflog.Print.Errorf("failed to create directory %s", sumCSVDir)
		return err
	}
	if err := os.MkdirAll(sumTXTDir, 0755); err != nil {
		oflog.Print.Errorf("failed to create directory %s", sumCSVDir)
		return err
	}

	// Walk through the source directory
	err := filepath.Walk(sourceDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			oflog.Print.Errorf("Walk through the source folder failed.")
			return err
		}

		// Skip directories and the target folders
		if info.IsDir() || path == sumCSVDir || path == sumTXTDir {
			oflog.Print.Infof("Skipping directories and the target folders.")
			return nil
		}

		// Determine the file type and move to the corresponding folder
		ext := strings.ToLower(filepath.Ext(info.Name()))
		var targetDir string
		if ext == ".csv" {
			targetDir = sumCSVDir
		} else if ext == ".txt" {
			targetDir = sumTXTDir
		} else {
			// Ignore unsupported file types
			return nil
		}

		// Move the file
		targetPath := filepath.Join(targetDir, info.Name())
		err = copyOrMerge(path, targetPath)
		if err != nil {
			oflog.Print.Errorf("%s Error:failed at algorithm.copyOrMerge!", getFunctionName())
			return err
		}
		oflog.Print.Infof("Moved file: %s -> %s", path, targetPath)

		return nil
	})

	if err != nil {
		return err
	}

	oflog.Print.Infof("Files have been organized into sumcsv and sumtxt.")
	return nil
}

// Function to merge folder contents from source to destination
func ExcelSumSinger(sourceDir string, destDir string) error {
	oflog.Init()
	// Walk through the source folder and copy each file/folder to the destination
	err := filepath.Walk(sourceDir, func(srcPath string, info os.FileInfo, err error) error {
		if err != nil {
			oflog.Print.Errorf("Walk through the source folder failed:%v", err)
			return err
		}

		// Get the relative path of the current file/folder
		relPath, _ := filepath.Rel(sourceDir, srcPath)
		destPath := filepath.Join(destDir, relPath)

		// If it's a file, copy it to the destination
		if !info.IsDir() {
			err := copyOrMerge(srcPath, destPath)
			if err != nil {
				oflog.Print.Errorf("%s Error:failed at algorithm.copyOrMerge!", getFunctionName())
				return err
			}
		} else {
			// If it's a directory, create the directory in the destination
			err := createDir(destPath)
			if err != nil {
				oflog.Print.Errorf("%s Error:failed at algorithm.createDir!", getFunctionName())
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}

	oflog.Print.Infof("%s have been organized into %s.", sourceDir, destDir)
	return nil
}

// Function to copy or merge files based on their existence
func copyOrMerge(src string, dest string) error {
	oflog.Print.Debugf("%s ============> Debug", getFunctionName())
	oflog.Init()
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

	// If dest exists, merge if both are XLS files
	if filepath.Ext(src) == ".txt" && filepath.Ext(dest) == ".txt" {
		return nil
	}
	// For non-CSV files, return an error or handle differently
	oflog.Print.Errorf("no supported files %s.", src)
	return nil
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
	oflog.Init()
	oflog.Print.Debugf("%s ============> Debug", getFunctionName())
	// Open source and destination files
	srcFile, err := os.Open(src)
	if err != nil {
		oflog.Print.Errorf("failed to open source file: %s", src)
		return err
	}
	defer srcFile.Close()

	destFile, err := os.OpenFile(dest, os.O_APPEND|os.O_WRONLY, os.ModePerm)
	if err != nil {
		oflog.Print.Errorf("failed to open destination file: %s", dest)
		return err
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

		// if err != nil {

		// 	oflog.Print.Infof("Skipping invalid row.")
		// 	continue // Skip invalid rows
		// }

		//Skipping Head
		if record[0] == "No." || record[0] == "" {
			oflog.Print.Infof("Skipping Head row.")
			continue // Skip invalid rows
		}
		err = destWriter.Write(record)
		if err != nil {
			oflog.Print.Errorf("failed to write record to destination.")
			return err
		}
		oflog.Print.Infof("Fuse ID %s move success.", record[1])
	}

	destWriter.Flush()
	if err := destWriter.Error(); err != nil {
		oflog.Print.Errorf("failed to flush writer.")
		return err
	}

	//oflog.Print.Infof("Merged %s into %s", src, dest)
	return nil
}

func removeFiles(path string) error {
	oflog.Init()
	// Read the directory contents
	entries, err := readDir(path)
	if err != nil {
		oflog.Print.Errorf("%s Error:failed at algorithm.readDir!", getFunctionName())
		return err
	}

	for _, entry := range entries {
		entryPath := path + "/" + entry.Name()

		if entry.IsDir() {
			// If it's a directory, remove it recursively
			err := removePath(entryPath)
			if err != nil {
				oflog.Print.Errorf("%s Error:failed at algorithm.removePath!", getFunctionName())
				return err
			}
			oflog.Print.Infof("%s Info:Deleted directory: %s", getFunctionName(), entryPath)
		} else {
			// If it's a file, delete it directly
			err := removeFile(entryPath)
			if err != nil {
				oflog.Print.Errorf("%s Error:failed at algorithm.removeFile!", getFunctionName())
				return err
			}
			oflog.Print.Infof("%s Info:Deleted file: %s", getFunctionName(), entryPath)
		}
	}

	return nil
}
