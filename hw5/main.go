package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type db map[string][]int

func getTextFromFile(filepath string) []string {
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

func containStrings(strings []string, what string) bool {
	for _, word := range strings {
		if word == what {
			return true
		}
	}
	return false
}

func containInts(numbers []int, what int) bool {
	for _, number := range numbers {
		if number == what {
			return true
		}
	}
	return false
}

func getLinesByDirectSearch(word string, inputText []string) ([]string, bool) {
	var foundlines = []string{}
	for _, line := range inputText {
		lineLower := strings.ToLower(line)
		lineLower = strings.TrimSuffix(lineLower, ".")
		lineLower = strings.TrimSuffix(lineLower, ",")
		stringsSlice := strings.Split(lineLower, " ")
		if containStrings(stringsSlice, word) {
			foundlines = append(foundlines, line)
		}
	}
	if len(foundlines) == 0 {
		return []string{"line with word '" + word + "' not found"}, false
	}

	return foundlines, true
}

func buildIndexes(inputText []string) db {
	indexes := make(db)
	for textIndex, textLine := range inputText {
		words := strings.Split(textLine, " ")
		for _, word := range words {
			word = strings.TrimSpace(word)
			word = strings.ToLower(word)
			word = strings.TrimSuffix(word, ".")
			word = strings.TrimSuffix(word, ",")
			if containInts(indexes[word], textIndex) {
				continue
			}
			if word == "" {
				continue
			}
			indexes[word] = append(indexes[word], textIndex)
		}
	}

	return indexes
}

func getLinesFromDb(what string, indexes db, lines []string) ([]string, bool) {
	ids, ok := indexes[what]
	if !ok {
		return []string{"line with word '" + what + "' not found"}, false
	}

	var linesOut = make([]string, len(ids))
	for i, id := range ids {
		linesOut[i] = lines[id]
	}

	return linesOut, true
}

func getRandomWords(indexes db) []string {
	allWords := make([]string, len(indexes))
	i := 0
	for k := range indexes {
		allWords[i] = k
		i++
	}
	var wordsOut []string
	wordsRandomQty := rand.Intn(len(allWords))
	for i = 0; i < wordsRandomQty; i++ {
		iRnd := rand.Intn(len(allWords))
		word := allWords[iRnd]
		if !containStrings(wordsOut, word) {
			wordsOut = append(wordsOut, word)
		}
	}

	return wordsOut
}

func drukerLines(comment string, lines []string, status bool) {
	fmt.Printf("%s", comment)
	if status {
		if len(lines) > 0 {
			for i, line := range lines {
				fmt.Printf("%d) %s\n", i+1, line)
			}
		}
		return
	}
	fmt.Printf("%s\n", lines[0])
}

func drukerTime(comment string, time time.Duration) {
	fmt.Printf("%s %v\n", comment, time)
}

func main() {
	const filepath string = "./lorem.txt"
	var what string
	start4 := time.Now()
	text := getTextFromFile(filepath)
	time4 := time.Since(start4)
	start2 := time.Now()
	indexes := buildIndexes(text)
	time2 := time.Since(start2)
	randomWords := getRandomWords(indexes)

	fmt.Print("What are we looking for? > ")
	fmt.Scanf("%s", &what)

	start1 := time.Now()
	lines1, status1 := getLinesByDirectSearch(what, text)
	time1 := time.Since(start1)

	start3 := time.Now()
	lines2, status2 := getLinesFromDb(what, indexes, text)
	time3 := time.Since(start3)

	start5 := time.Now()
	for i := 0; i < len(randomWords); i++ {
		getLinesByDirectSearch(randomWords[i], text)
	}
	time5 := time.Since(start5)

	start6 := time.Now()
	for i := 0; i < len(randomWords); i++ {
		getLinesFromDb(what, indexes, text)
	}
	time6 := time.Since(start6)

	drukerTime("readfile time: ", time4)
	drukerTime("time of build indexes (hw5): ", time2)

	drukerLines("-> compare lines (hw4):\n", lines1, status1)
	drukerLines("-> with lines (hw5):\n", lines2, status2)

	drukerTime("time of direct search (hw4): ", time1)
	drukerTime("time of index search (hw5): ", time3)

	fmt.Println("random words:", randomWords)
	drukerTime("time of direct search random words: ", time5)
	drukerTime("time of index searc random words: ", time6)
}
