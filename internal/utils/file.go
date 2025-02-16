package utils

import (
	"fmt"
	"io"
	"os"
)

func ReadJSONFile(filename string) ([]byte, error) {
	// Open the file
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	// Read the file content
	fileContent, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	return fileContent, nil
}

func CreateFile(filename string, extention string) (*os.File, error) {
	// Check if file exists
	if _, err := os.Stat(filename); os.IsNotExist(err) {

		file, err := os.Create(fmt.Sprintf("%s.%s", filename, extention))
		if err != nil {
			return nil, fmt.Errorf("failed to create file: %w", err)
		}
		return file, nil
	}
	// Open existing file
	file, err := os.OpenFile(filename, os.O_RDWR, 0644)
	if err != nil {
		return nil, fmt.Errorf("failed to open file: %w", err)
	}
	return file, nil
}
