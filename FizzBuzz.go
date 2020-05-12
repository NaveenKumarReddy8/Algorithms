/*
Fizz-Buzz algorithm:
Fizz and Buzz refer to any number that's a multiple of 3 and 5 respectively. In other words, if a number is divisible
by 3, it is substituted with fizz; if a number is divisible by 5, it is substituted with buzz. If a number is simultaneously
a multiple of 3 AND 5, the number is replaced with "fizz buzz." In essence, it emulates the famous children game
"fizz buzz".
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func typeCastToIntegers(elementString string) []int {
	elements := make([]int, 0)
	splittedIntegerStrings := strings.Fields(elementString)
	for _, value := range splittedIntegerStrings {
		value = strings.TrimSpace(value)
		integerValue, err := strconv.Atoi(value)
		if err != nil {
			fmt.Println(err)
		} else {
			elements = append(elements, integerValue)
		}
	}
	return elements
}

func getInput() string {
	elements := make([]int, 0)
	var integerString string
	buffer := bufio.NewScanner(os.Stdin)
	fmt.Printf("Please enter the integer values separated by a space\n>")
	if buffer.Scan() {
		integerString = buffer.Text()
	}
	integerString = strings.TrimSpace(integerString)
	return integerString
}

func getFizzBuzz() (fizz, buzz int) {
	fmt.Println("Enter the value for Fizz")
	_, err := fmt.Scanf("%d", &fizz)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Enter the value for Buzz")
	_, err = fmt.Scanf("%d", &buzz)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func fizzBuzzer(elements []int, fizz, buzz int) {
	fmt.Println("The results are:")
	for _, value := range elements {
		remainderByFizz := value % fizz
		remainderByBuzz := value % buzz
		if remainderByFizz == 0 && remainderByBuzz == 0 {
			fmt.Printf("%d -> FizzBuzz \n", value)
		} else if remainderByFizz == 0 {
			fmt.Printf("%d -> Fizz \n", value)
		} else if remainderByBuzz == 0 {
			fmt.Printf("%d -> Buzz \n", value)
		} else {
			fmt.Println(value)
		}
	}
}

func main() {
	elementString := getInput()
	fmt.Printf("The elements are %v \n", elements)
	elements := typeCastToIntegers(elementString)
	fizz, buzz := getFizzBuzz()
	fmt.Printf("The Fizz is: %d, The Buzz is: %d \n", fizz, buzz)
	fizzBuzzer(elements, fizz, buzz)
}
