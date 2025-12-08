package utils

import (
	"bufio"
	"os"
)

func ConsumeInputMatrix(path string) ([][]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var matrix [][]string
	for scanner.Scan() {
		line := scanner.Text()
		row := make([]string, 0, len(line))
		for _, r := range line {
			row = append(row, string(r))
		}
		matrix = append(matrix, row)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matrix, nil
}
