package main

import (
	"flag"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func main() {
	// Define and parse the command line argument for the folder path
	folderPath := flag.String("path", "", "Path to the folder to clean up .git directories")
	flag.Parse()

	if *folderPath == "" {
		fmt.Println("Please specify a folder path using the -path option.")
		os.Exit(1)
	}

	// Call the function to remove .git directories
	err := removeGitFolders(*folderPath)
	if err != nil {
		fmt.Printf("Error removing .git directories: %s\n", err)
		os.Exit(1)
	}

	fmt.Println("Completed removing .git directories.")
}

// removeGitFolders searches for and removes .git directories recursively
func removeGitFolders(path string) error {
	return filepath.WalkDir(path, func(currentPath string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		// Skip non-directory files immediately
		if !d.IsDir() {
			return nil
		}

		// If a .git directory is found, remove it
		if d.Name() == ".git" {
			fmt.Printf("Removing .git directory: %s\n", currentPath)
			if err := os.RemoveAll(currentPath); err != nil {
				return err
			}
			// Skip walking into the .git directory since we've removed it
			return filepath.SkipDir
		}

		return nil
	})
}
