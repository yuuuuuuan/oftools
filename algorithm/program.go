package algorithm

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ProgramFirewareSingle(source string) error {
	var err error
	// inputDir ,err := getInputPath()
	// if err !=nil{
	// 	return fmt.Errorf("Error:%e", err)
	// }
	root := source + "\\OIS"
	targets := []string{"A", "B", "C", "D"}
	// Call the function and retrieve matching paths
	paths, err := findSpecificPaths(root, targets)
	if err != nil {
		return fmt.Errorf("Error:%e", err)
	} else {
		for index, value := range paths {
			err = moveFirewareFile(value)
			if err != nil {
				fmt.Printf("%s Error:Move FirewareFile failed at %d\n", getFunctionName(), index)
			}
			iniFile := findSpecificFile(value)
			if iniFile != "" {
				fmt.Printf("%s  %d\n", getFunctionName(), index)

			}
		}
	}
	return nil
}

// func getInputPath() (string, error) {
// 	fmt.Println("please input program dir:")

// 	reader := bufio.NewReader(os.Stdin)
// 	path, _ := reader.ReadString('\n')

// 	path = strings.TrimSpace(path)

// 	if path == "" {
// 		return "", errors.New("dir is void.")
// 	}

// 	return path, nil
// }

func moveFirewareFile(source string) error {
	var srcPath string
	var err error
	fmt.Println("Please input OIS fireware root:")
	fmt.Scanln(&srcPath)
	fmt.Printf("You input OIS fireware root is:%s", srcPath)
	if !strings.HasSuffix(srcPath, ".bin") && !strings.HasSuffix(srcPath, ".hex") {
		fmt.Printf("%s Error:There is no bin or hex file.\n", getFunctionName())
	} else {
		err = copyFile(srcPath, source)
		if err != nil {
			fmt.Printf("%s Error:Move file failed.\n", getFunctionName())
		}
	}
	fmt.Printf("%s:move file Successed.\n", getFunctionName())
	return nil
}

func findSpecificFile(root string) string {
	lastName := filepath.Base(root)
	path := root + "\\" + "OIS_FWDownlod_" + lastName + "_OIS_Self.ini"
	if ifFileExist(path) {
		return path
	}
	return ""
}

func ifFileExist(path string) bool {
	_, err := os.Stat(path)

	if os.IsNotExist(err) {
		return false
	} else {
		return true
	}
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
