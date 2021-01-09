package sorting

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func splitAndTypecastToInt(inputStr string) []int {
	stringSlice := strings.Fields(inputStr)
	typecastedIntegers := make([]int, 0)
	for _, value := range stringSlice {
		typecastedString, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal("Error while typecasting the user inputted splitted strings to integers", err)
		}
		typecastedIntegers = append(typecastedIntegers, typecastedString)
	}
	return typecastedIntegers
}

func getUserInput() string {
	fmt.Printf("Enter the numbers to be sorted with spaces between each\n>")
	scanner := bufio.NewReader(os.Stdin)
	stringOfNumbers, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading the user input", err)
	}
	return stringOfNumbers
}

// ObtainAndParseInput accepts the strings from the user
// ,splits it values separated by spaces, typecasts the
// splitted string slice to integers of base 10.
func ObtainAndParseInput() []int {
	return splitAndTypecastToInt(getUserInput())
}
