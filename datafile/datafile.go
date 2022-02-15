//package datafile allows reading data samples from files
package datafile

import (
	"bufio"
	"os"
)

//getStrings read a string from each line of a file
func GetStrings(filename string) ([]string, error) {
	var lines []string

	file, error := os.Open(filename)
	if error != nil {
		return nil, error
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}
	err := file.Close()
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

	return lines, nil
}
