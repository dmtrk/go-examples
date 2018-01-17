package util

import (
	"os"
	"bufio"
	"strings"
	"io"
)

func ParsePropertiesFromFile(filename string) (map[string]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ParseProperties(file), nil
}

func ParseProperties(reader io.Reader) (map[string]string) {
	properties := make(map[string]string)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if (len(line) > 0 && strings.Contains(line, "=")) {
			keyVals := strings.Split(line, "=")
			if (len(keyVals) > 1) {
				properties[strings.TrimSpace(keyVals[0])] = strings.TrimSpace(keyVals[1])
			}
		}
	}
	return properties;
}