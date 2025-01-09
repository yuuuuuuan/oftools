package algorithm

import (
	"io"
	"oftools/oflog"
	"os"
	"runtime"
)

// GetFunctionName retrieves the name of the currently executing function
func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	funcObj := runtime.FuncForPC(pc)
	return funcObj.Name()
}

// Function to copy a single file from src to dest
func copyFile(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		oflog.Print.Errorf("failed to open source file.")
		return err
	}
	defer srcFile.Close()

	destFile, err := os.Create(dest)
	if err != nil {
		oflog.Print.Errorf("failed to create destination file.")
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		oflog.Print.Errorf("failed to copy file.")
		return err
	}

	oflog.Print.Infof("Copied %s to %s", src, dest)
	return nil
}

func createDir(destDir string) error {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		oflog.Print.Errorf("failed to create destination directory.")
		return err
	}
	return nil
}

func readDir(path string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(path)
	if err != nil {
		oflog.Print.Errorf("failed to read destination directory.")
		return nil, err
	}
	return entries, nil
}

func removePath(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		oflog.Print.Errorf("failed to remove all directory.")
		return err
	}
	return nil
}

func removeFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		oflog.Print.Errorf("failed to remove file.")
		return err
	}
	return nil
}
