package utils

import (
	"bufio"
	"os"
)
func ComsumeInputNewLines(path string) ([][]string, error ) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var res [][]string

	for scanner.Scan() {
		newLine := []string{}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := scanner.Text()
		newLine = append(newLine, line)
		res = append(res, newLine)
	}

	return res, nil
}
