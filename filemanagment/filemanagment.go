package filemanagment

import (
	"bufio"
	"errors"
	"os"
)

func LinesReader(path string) ([]string, error) {
	file, err := os.Open(path)

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
