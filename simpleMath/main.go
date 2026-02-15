package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	firstNum, err := PromptUser("What is the first number?")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	secondNum, err := PromptUser("What is the first number?")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	sum, subs, multiply, divide, err := run(firstNum, secondNum)
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	// print the result
	fmt.Printf("Sum: %d\n", sum)
	fmt.Printf("Subs: %d\n", subs)
	fmt.Printf("Multiply: %d\n", multiply)
	fmt.Printf("Divide: %d\n", divide)
}

func run(num1, num2 string) (int, int, int, int, error) {
	n1, err := strconv.Atoi(num1)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("ERROR WMHEHEHEH!")
	}
	n2, err := strconv.Atoi(num2)
	if err != nil {
		return 0, 0, 0, 0, fmt.Errorf("ERROR WMHEHEHEH!")
	}

	return n1 + n2, n1 - n2, n1 * n2, n1 / n2, nil
}

func PromptUser(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + " ")
	scanner.Scan()
	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("cannot reading input: %v", err)
	}

	return text, nil
}
