package main

import "fmt"

func main() {
	sum := GetSumOfMultiples(20)
	fmt.Println(sum)
}

/*
	If we list all the natural numbers
	below 10 that are multiples of 3 or 5,
	we get 3, 5, 6 and 9. The sum of these multiples is 23.
	Find the sum of all the multiples of 3 or 5 below 1000.
*/

/*
	1. divide number by 3, store quotient1
	2. divide number by 5, store quotient2
	3. iterate through result1 to get multiples of 3, store multiples
	4. iterate through result2 to get multiples of 3, store multiples
*/

func GetSumOfMultiples(number int) int {
	quotient1 := number / 3
	quotient2 := number / 5

	multiplesOf3 := getMultiplesOfNumber(number, quotient1, 3)
	multiplesOf5 := getMultiplesOfNumber(number, quotient2, 5)

	// remove dividend from multiples
	multiplesOf3 = popDividend(multiplesOf3, number)
	multiplesOf5 = popDividend(multiplesOf5, number)

	sumOfMultiplesOf3 := sumArray(multiplesOf3)
	sumOfMultiplesOf5 := sumArray(multiplesOf5)

	return sumOfMultiplesOf3 + sumOfMultiplesOf5
}

func getMultiplesOfNumber(dividend int, quotient int, divisor int) []int {
	var multiples = make([]int, quotient)
	for i := 0; i < quotient; i++ {
		multiples[i] = (i + 1) * divisor
	}
	fmt.Println(multiples)
	return multiples
}

func sumArray(number []int) int {
	length := len(number)
	var sum int
	for i := 0; i < length; i++ {
		sum += number[i]
	}
	return sum
}

func popDividend(number []int, dividend int) []int {
	lastIndex := len(number) - 1
	lastNumber := number[lastIndex]
	popedArr := number
	if dividend == lastNumber {
		newIndex := lastIndex
		popedArr = popedArr[:newIndex]
	}
	fmt.Println(popedArr)
	return popedArr
}
