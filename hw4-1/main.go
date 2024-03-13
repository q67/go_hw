package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const filepath string = "./lorem.txt"

func readfile(filepath string) []string {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return []string{}
	}
	defer file.Close()

	var receivedText = []string{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		receivedText = append(receivedText, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return receivedText
}

func finder(substring string, inputText []string) []string {
	var foundlines = []string{}
	for _, line := range inputText {
		if strings.Contains(line, substring) {
			foundlines = append(foundlines, line)
		}
	}

	return foundlines
}

func druker(lines []string) {
	for i, line := range lines {
		fmt.Printf("%d) %s\n", i+1, line)
	}
}

func main() {
	var what string
	fmt.Print("What are we looking for? > ")
	fmt.Scanf("%s", &what)
	text := readfile(filepath)
	lines := finder(what, text)
	druker(lines)
}
