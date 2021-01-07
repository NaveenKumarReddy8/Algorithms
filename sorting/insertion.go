package sorting

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func performInsertionSort(toBeSorted []int) []int {
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
	fmt.Printf("Enter the numbers to be sorted with spaces between each\n>")
	scanner := bufio.NewReader(os.Stdin)
	stringOfNumbers, err := scanner.ReadString('\n')
	if err != nil {
		log.Fatal("Error while reading the user input", err)
	}
	return stringOfNumbers
}

// Insertion .
func Insertion() {
	inputStr := getUserInput()
	integerInput := splitAndTypecastToInt(inputStr)
	startTime := time.Now()
	sortedSlice := performInsertionSort(integerInput)
	fmt.Printf("Time taken for Insertion sort for %d elements is %s\n", len(integerInput), time.Since(startTime))
	fmt.Println("The sorted slice is: ", sortedSlice)
}
