package handlers

import (
	"bufio"
	"fmt"
	"os"
)

func getFilePath(filename string) (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/internal/tmp/%s", path, filename), nil

}

func ReadFileAndSendToChannel(filename string, rawMsgChannel chan string) error {
	path, err := getFilePath(filename)
	if err != nil {
		return err
	}

	file, err := os.OpenFile(path, os.O_RDONLY, 0644)
	if err != nil {
		return err
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	for scanner.Scan() {
		rawMsgChannel <- scanner.Text()
	}
	close(rawMsgChannel)

	if scanner.Err() != nil {
		return scanner.Err()
	}

	return nil
}
