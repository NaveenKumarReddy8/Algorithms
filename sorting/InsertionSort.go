package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func insertionSort(toBeSorted []int) []int {
	lengthOfSlice := len(toBeSorted)
	if lengthOfSlice == 1 {
		return toBeSorted
	}

	for index := 1; index < lengthOfSlice; index++ {
		key := toBeSorted[index]
		previousIndex := index - 1
		for ; previousIndex >= 0; previousIndex-- {
			if toBeSorted[previousIndex] > key {
				toBeSorted[previousIndex+1] = toBeSorted[previousIndex]
			} else {
				break
			}
		}
		toBeSorted[previousIndex+1] = key
	}
	return toBeSorted
}

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
	fmt.Println("Enter the numbers to be sorted with spaces between each")
	scanner := bufio.NewReader(os.Stdin)
	stringOfNumbers, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading the user input", err)
	}
	return stringOfNumbers
}

func main() {
	inputStr := getUserInput()
	integerInput := splitAndTypecastToInt(inputStr)
	fmt.Println("The input is: ", integerInput)
	sortedSlice := insertionSort(integerInput)
	fmt.Println("The sorted slice is: ", sortedSlice)
}
