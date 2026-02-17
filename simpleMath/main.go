package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type Result struct {
	sum       float64
	substract float64
	multiply  float64
	divide    float64
}

func main() {
	firstNum, err := PromptUser("What is the first number?")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}
	secondNum, err := PromptUser("What is the second number?")
	if err != nil {
		log.Fatalf("ERROR: %v", err)
	}

	result, err := run(firstNum, secondNum)
	if err != nil {
		log.Fatalf("Run failed: %v", err)
	}

	fmt.Printf("Sum: %v\n", result.sum)
	fmt.Printf("Substract: %v\n", result.substract)
	fmt.Printf("Multiply: %v\n", result.multiply)
	fmt.Printf("Divide: %v\n", result.divide)
}

func run(num1, num2 string) (Result, error) {
	n1, err := strconv.Atoi(num1)
	if err != nil {
		return Result{}, fmt.Errorf("invalid num1: %w", err)
	}

	n2, err := strconv.Atoi(num2)
	if err != nil {
		return Result{}, fmt.Errorf("invalid num2: %w", err)
	}

	sumVal, err := sum(n1, n2)
	if err != nil {
		return Result{}, err
	}

	subVal, err := substract(n1, n2)
	if err != nil {
		return Result{}, err
	}

	mulVal, err := multiply(n1, n2)
	if err != nil {
		return Result{}, err
	}

	divVal, err := divide(n1, n2)
	if err != nil {
		return Result{}, err
	}

	return Result{
		sum:       sumVal,
		substract: subVal,
		multiply:  mulVal,
		divide:    divVal,
	}, nil
}

func sum(num1, num2 int) (float64, error) {
	return float64(num1 + num2), nil
}

func substract(num1, num2 int) (float64, error) {
	return float64(num1 - num2), nil
}

func multiply(num1, num2 int) (float64, error) {
	return float64(num1 * num2), nil
}

func divide(num1, num2 int) (float64, error) {
	if num2 == 0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}
	return float64(num1 / num2), nil
}

func PromptUser(prompt string) (string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print(prompt + " ")
	scanner.Scan()
	text := scanner.Text()

	if err := scanner.Err(); err != nil {
		return "", fmt.Errorf("Cannot reading input: %v", err)
	}

	return text, nil
}
