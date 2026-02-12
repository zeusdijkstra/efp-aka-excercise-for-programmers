package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	length := StringPrompt("What is the length of the room in feet?")
	width := StringPrompt("What is the width of the room in feet?")
	fmt.Printf("You entered dimensions of %s feet by %s feet.\n", length, width)

	intLength, err := strconv.Atoi(length)
	if err != nil {
		log.Fatalf("invalid length: %v", err)
	}

	intWidth, err := strconv.Atoi(width)
	if err != nil {
		log.Fatalf("invalid length: %v", err)
	}

	fmt.Println("calculating...")
	time.Sleep(1000 * time.Millisecond)

	area, err := calculateFeet(intLength, intWidth)
	if err != nil {
		log.Fatalf("calculation failed: %v", err)
	}

	toMeters, err := feetToMeters(area)
	if err != nil {
		log.Fatalf("conversion failed: %v", err)
	}

	fmt.Printf("The area is %d feet\n", area)
	fmt.Printf("%.2f meters\n", toMeters)
}

func calculateFeet(length, width int) (int, error) {
	if length < 0 || width < 0 {
		return 0, ErrNegativeValue
	}

	if length == 0 || width == 0 {
		return 0, ErrZeroValues
	}

	result := length * width

	return result, nil
}

func feetToMeters(feet int) (float64, error) {
	// handle error
	if feet == 0 {
		return 0, ErrZeroValues
	}
	if feet < 0 {
		return 0, ErrNegativeValue
	}

	converter := 0.3048
	result := converter * float64(feet)

	return result, nil
}

func StringPrompt(label string) string {
	var s string
	r := bufio.NewReader(os.Stdin)
	for {
		fmt.Fprint(os.Stdout, label+" ")
		s, _ = r.ReadString('\n')
		if s != "" {
			break
		}
	}
	return strings.TrimSpace(s)
}
