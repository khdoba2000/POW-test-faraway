package io

import (
	"bufio"
	"log"
	"os"
)

func ReadFile(filepath string) ([]string, error) {
	_, err := os.Stat(filepath)
	if err != nil {
		return nil, err
	}
	f, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(f)

	var lines []string
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		i++
	}

	return lines, nil
}
