package filemanagment

import (
	"bufio"
	"encoding/json"
	"errors"
	"os"
)

type FileManagment struct {
	// InputFilePath  string
	// OutputFilePath string `json:"output_file_path"`
	inputFilePath  string
	outputFilePath string
}

func New(iPath, oPath string) FileManagment {
	return FileManagment{
		// InputFilePath:  iPath,
		// OutputFilePath: oPath,
		inputFilePath:  iPath,
		outputFilePath: oPath,
	}
}

func (io FileManagment) LinesReader() ([]string, error) {
	// file, err := os.Open(io.InputFilePath)
	file, err := os.Open(io.inputFilePath)

	if err != nil {
		return nil, errors.New("can't open target file")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	lines := []string{}

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	err = scanner.Err()
	if err != nil {
		return nil, errors.New("can't read one of target file's line")
	}

	return lines, nil
}

func (io FileManagment) WriteJSON(data any) error {
	// file, err := os.Create(io.OutputFilePath)
	file, err := os.Create(io.outputFilePath)

	if err != nil {
		return errors.New("can't create file")
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(data)

	if err != nil {
		return errors.New("can't convert data to JSON")
	}

	return nil
}
