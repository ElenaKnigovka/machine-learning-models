package machinelearningmodels

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/gorilla/mux"
)

func LoadJSONFile(filePath string, v interface{}) error {
	data, err := ioutil.ReadFile(filePath)
	if err!= nil {
		return err
	}
	return json.Unmarshal(data, v)
}

func WriteJSONFile(filePath string, v interface{}) error {
	data, err := json.MarshalIndent(v, "", "  ")
	if err!= nil {
		return err
	}
	return ioutil.WriteFile(filePath, data, 0644)
}

func GetEnvironmentVariableOrDefault(envVar string, defaultValue string) string {
	value := os.Getenv(envVar)
	if value == "" {
		return defaultValue
	}
	return value
}

func GetAbsolutePath(relativePath string) string {
	_, err := filepath.Abs(relativePath)
	if err!= nil {
		log.Fatal(err)
	}
	return relativePath
}

func GetRouteParam(r *mux.Router, param string) string {
	vars := mux.Vars(r)
	return vars[param]
}

func IsDirectoryExist(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetDirectoryFiles(dirPath string) []string {
	files, err := ioutil.ReadDir(dirPath)
	if err!= nil {
		log.Fatal(err)
	}
	var filePaths []string
	for _, file := range files {
		filePaths = append(filePaths, filepath.Join(dirPath, file.Name()))
	}
	return filePaths
}

func GetFileNameWithoutExtension(filePath string) string {
	ext := filepath.Ext(filePath)
	return strings.TrimSuffix(filePath, ext)
}

func GetEnvironmentVariables() map[string]string {
	vars := make(map[string]string)
	for _, env := range os.Environ() {
		parts := strings.SplitN(env, "=", 2)
		vars[parts[0]] = parts[1]
	}
	return vars
}

func IsFileExist(filePath string) bool {
	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func GetDirectoryPath(filePath string) string {
	dirPath := filepath.Dir(filePath)
	if dirPath == "." {
		return ""
	}
	return dirPath
}