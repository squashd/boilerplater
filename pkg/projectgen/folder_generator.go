// projectgen is for the currently unimplemented boilerplate project structure and functions
package projectgen

import (
	"os"
	"path/filepath"

	"github.com/SQUASHD/boilerplater/internal/shared/models"
)

func GenerateProjectStructure(projectStructure models.ProjectStructure) error {
	rootDir, err := GetRootDirectory(projectStructure.RootDirectory)
	if err != nil {
		return err
	}

	for _, dir := range projectStructure.Directories {
		if err := createDirectory(rootDir, dir); err != nil {
			return err
		}
	}

	for _, file := range projectStructure.MainFiles {
		if err := createFile(filepath.Join(rootDir, file)); err != nil {
			return err
		}
	}

	return nil
}

func GetRootDirectory(rootDir string) (string, error) {
	if rootDir == "" {
		return os.Getwd()
	}
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
