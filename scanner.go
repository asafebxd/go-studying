package main

import (
	"bufio"
	"os"
)

func saveMessage(filename, message string) error {
	content := []byte(message)

	err := os.WriteFile(filename, content, 0644)
	if err != nil {
		return err
	}

	return nil
}

// func saveMessage(filename, message string) error {
// 	return os.WriteFile(filename, []byte(message), 0644)
// } Shorter version

func readMessage(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func appendLog(filename, message string) error {
	file, err := os.OpenFile(
		filename,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0644,
	)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(message + "\n")
	return err
}

func countLines(filename string) (int, error) {
	lineCount := 0
	file, err := os.Open(filename)
	if err != nil {
		return lineCount, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		lineCount++
	}

	if err := scanner.Err(); err != nil {
		return lineCount, err
	}

	return lineCount, nil
}

func scanner() {
	
}