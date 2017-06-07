package main

import (
	"fmt"
	"os"
	"time"
	"io/ioutil"
	"strings"
	"runtime"
)

// Folders that we want to clean
var folders = [...]string{"/Users/Nemanja/Downloads", "/Users/Nemanja/Documents/ViberDownloads"}
// Current time
// TODO: Implement deleting file only if it's older than x days
var now = time.Now()
// Things that we want to skip
var skip = [...]string{"Udemy"}

func main()  {
	// Detecting system and using file separator for that system
	var fileSeparator string
	if detectSystemIsWindows() {
		fileSeparator = "\\"
	} else {
		fileSeparator = "/"
	}
	fmt.Println("STARTING")
	fmt.Println("Skipping files that contain:", skip)
	fmt.Println("------------------------------------")
	var numDeleted int =  0
	var numOfSkipped int = 0
	for _,initialPath := range folders {
		fmt.Println("Deleting in:", initialPath)
		fmt.Println(".............................")
		files, _ := ioutil.ReadDir(initialPath)
		for _, f := range files {
			// Checking if we have to skip that file or folder
			if checkForSkip(f.Name()) {
				numOfSkipped ++
				continue
			}
			// Building full path
			fullPath := initialPath + fileSeparator + f.Name()
			// Deleting file or folder
			deleteFileOrFolder(fullPath)
			numDeleted ++
		}
		fmt.Println(".............................")
	}
	fmt.Println("------------------------------------")
	fmt.Println("COMPLETED")
	fmt.Println("Number of deleted files:", numDeleted)
	fmt.Println("Number of skipped files:", numOfSkipped)
}

/*
Function used for checking if we should skip file or folder
 */
func checkForSkip(fileToSkip string) bool {
	for _,f := range skip {
		if strings.Contains(fileToSkip, f) {
			return true
		}
	}
	return false
}

/*
Function used for actual deleting file
We are using RemoveAll so we can delete both
files and folders
 */
func deleteFileOrFolder(file string) {
	var err = os.RemoveAll(file)
	checkError(err)
}

/*
Error catching function
 */
func checkError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}
}

/*
Function used at runtime to detect operating system
 */
func detectSystemIsWindows() bool{
	if runtime.GOOS == "windows" {
		return true
	}
	return false
}


