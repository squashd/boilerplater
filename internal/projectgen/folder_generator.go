package projectgen

import (
	"os"
	"path/filepath"

	"github.com/SQUASHD/boilerplater/shared/models"
)

// GenerateProjectStructure creates the directories and files based on the provided project structure.
func GenerateProjectStructure(projectStructure models.ProjectStructure) error {
	rootDir, err := getRootDirectory(projectStructure.RootDirectory)
	if err != nil {
		return err
	}

	// Create subdirectories and files
	for _, dir := range projectStructure.Directories {
		if err := createDirectory(rootDir, dir); err != nil {
			return err
		}
	}

	// Create main files in the root directory
	for _, file := range projectStructure.MainFiles {
		if err := createFile(filepath.Join(rootDir, file)); err != nil {
			return err
		}
	}

	return nil
}

func getRootDirectory(rootDir string) (string, error) {
	if rootDir == "" {
		return os.Getwd()
	}
	// Create the root directory if it's specified
	if err := os.MkdirAll(rootDir, 0755); err != nil {
		return "", err
	}
	return rootDir, nil
}

// createDirectory recursively creates a directory and its subdirectories and files.
func createDirectory(basePath string, dir models.Directory) error {
	dirPath := filepath.Join(basePath, dir.Name)
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return err
	}

	for _, file := range dir.Files {
		if err := createFile(filepath.Join(dirPath, file)); err != nil {
			return err
		}
	}

	for _, subDir := range dir.SubDirs {
		if err := createDirectory(dirPath, subDir); err != nil {
			return err
		}
	}

	return nil
}

// createFile creates an empty file at the specified path.
func createFile(filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	return file.Close()
}
