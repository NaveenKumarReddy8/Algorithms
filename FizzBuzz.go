/*
Fizz-Buzz algorithm:
Fizz and Buzz refer to any number that's a multiple of 3 and 5 respectively. In other words, if a number is divisible
by 3, it is substituted with fizz; if a number is divisible by 5, it is substituted with buzz. If a number is simultaneously
a multiple of 3 AND 5, the number is replaced with "fizz buzz." In essence, it emulates the famous children game
"fizz buzz".
*/

package main

import (
	"fmt"
	"strconv"
)

func stringOrConvertToInt(input string) interface{} {
	intValue, err := strconv.Atoi(input)
	if err != nil {
		return input
	} else {
		return intValue
	}
}

func getInput() []int {
	// Creating a slice to which integers would be appended which are given as inputs by the user.
	elements := make([]int, 0, 10)
	for {
		var userInput string
		fmt.Println("Type any integer value to proceed, type 'done' if feeding the integers is completed")
		fmt.Print(">")
		_, err := fmt.Scanf("%s", &userInput)
		if err != nil {
			fmt.Println(err)
		} else {
			input := stringOrConvertToInt(userInput)
			integerValue, ok := input.(int)
			if ok {
				elements = append(elements, integerValue)
				continue
			} else {
				stringValue, ok := input.(string)
				if ok {
					if stringValue == "done" {
						break
					} else {
						fmt.Println("Invalid input! Please type 'done' if completed with feeding the integers.")
					}
				}
			}
		}
	}
	return elements

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
	elements := getInput()
	fmt.Printf("The elements are %v \n", elements)
	fizz, buzz := getFizzBuzz()
	fmt.Printf("The Fizz is: %d, The Buzz is: %d \n", fizz, buzz)
	fizzBuzzer(elements, fizz, buzz)
}
