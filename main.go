package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter text: ")
	text, _ := reader.ReadString('\n')

	wordCount := countWords(text)
	sentenceCount := countSentences(text)
	characterCount := countCharacters(text)

	fmt.Println("Word count:", wordCount)
	fmt.Println("Sentence count:", sentenceCount)
	fmt.Println("Character count:", characterCount)
}

func countWords(text string) int {
	words := strings.Fields(text)
	return len(words)
}

func countSentences(text string) int {
	sentences := strings.Split(text, ".")
	return len(sentences)
}

func countCharacters(text string) int {
	return len(text) - strings.Count(text, " ")
}
