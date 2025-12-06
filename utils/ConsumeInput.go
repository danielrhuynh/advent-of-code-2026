package utils

import (
	"bufio"
	"os"
)
func ComsumeInput(path string) ([]string, error ) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var res []string

	for scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return nil, err
		}
		line := scanner.Text()
		res = append(res, line)
	}

	return res, nil
}
