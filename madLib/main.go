package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

func main() {
	noun, err := PromptUser("Enter a noun: ")
	if err != nil {
		log.Fatalf("Error reading prompt: %v", err)
	}

	verb, err := PromptUser("Enter a verb: ")
	if err != nil {
		log.Fatalf("Error reading prompt: %v", err)
	}

	adjective, err := PromptUser("Enter a adjective: ")
	if err != nil {
		log.Fatalf("Error reading prompt: %v", err)
	}

	adverb, err := PromptUser("Enter a adverb: ")
	if err != nil {
		log.Fatalf("Error reading prompt: %v", err)
	}

	fmt.Println("building strings...")
	time.Sleep(1000 * time.Millisecond)

	fmt.Println(BuildingStrings(noun, verb, adjective, adverb, 0))
}

func PromptUser(question string) (string, error) {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print(question)

	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("Error reading input: %v", err)
	}

	input = strings.TrimSpace(input)
	return input, nil
}

func BuildingStrings(noun, verb, adjective, adverb string, stringPicker int) string {
	templateStrings := []string{
		"Do you %s your %s %s %s? That's hilarious!",
		"Why would you %s a %s %s %s?",
		"I watched him %s the %s %s %s.",
		"Please don't %s that %s %s %s.",
		"It’s hard to %s %s %s %s when everyone is watching.",
		"Could you %s your %s %s %s next time?",
		"He decided to %s the %s %s %s for the sake of the mission.",
		"You shouldn't %s %s %s %s in a library.",
		"If you %s that %s %s %s, you'll regret it!",
		"She managed to %s a %s %s %s.",
		"Nobody expects you to %s %s %s %s on a Tuesday.",
	}
	randomNumber := stringPicker

	template := templateStrings[randomNumber]

	result := fmt.Sprintf(template, verb, adjective, noun, adverb)
	return result
}
