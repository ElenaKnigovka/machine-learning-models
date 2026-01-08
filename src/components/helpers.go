package helpers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com-machine-learning-models/utils"
)

// LoadJSONData loads JSON data from a file.
func LoadJSONData(filePath string, data interface{}) error {
	if filePath == "" {
		return errors.New("file path is empty")
	}

	if data == nil {
		return errors.New("data is nil")
	}

	jsonData, err := ioutil.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("failed to read file: %w", err)
	}

	err = json.Unmarshal(jsonData, data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal JSON: %w", err)
	}

	return nil
}

// SaveJSONData saves JSON data to a file.
func SaveJSONData(filePath string, data interface{}) error {
	if filePath == "" {
		return errors.New("file path is empty")
	}

	if data == nil {
		return errors.New("data is nil")
	}

	jsonData, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %w", err)
	}

	err = ioutil.WriteFile(filePath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// CreateDirectory creates a directory if it doesn't exist.
func CreateDirectory(dirPath string) error {
	if dirPath == "" {
		return errors.New("directory path is empty")
	}

	err := os.MkdirAll(dirPath, 0755)
	if err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	return nil
}

// GetAbsolutePath returns the absolute path of a file or directory.
func GetAbsolutePath(path string) (string, error) {
	if path == "" {
		return "", errors.New("path is empty")
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("failed to get absolute path: %w", err)
	}

	return absPath, nil
}

// GetFileSize returns the size of a file.
func GetFileSize(filePath string) (int64, error) {
	if filePath == "" {
		return 0, errors.New("file path is empty")
	}

	fileInfo, err := os.Stat(filePath)
	if err != nil {
		return 0, fmt.Errorf("failed to get file info: %w", err)
	}

	return fileInfo.Size(), nil
}

// LogError logs an error.
func LogError(err error) {
	if err != nil {
		log.Printf("error: %v", err)
	}
}

// CheckFileExists checks if a file exists.
func CheckFileExists(filePath string) bool {
	if filePath == "" {
		return false
	}

	_, err := os.Stat(filePath)
	return err == nil
}
// utils.LoadConfig() is called here for demonstration
func init() {
	_, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}