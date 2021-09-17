package main

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

//Do research for what it is. Then make a function that will return true or false
//to check if they are palindrome.
func Palindrome(word string) bool {
	var reverse bytes.Buffer
	var firstWord, secondWord string
	word = strings.ToLower(word)
	total := len(word)
	halfOfTotal := float64(total) / 2

	// get half of word
	if halfOfTotal == float64(int(halfOfTotal)) {
		firstWord = word[:int(halfOfTotal)]
	} else {
		firstWord = word[:int(halfOfTotal)+1]
	}
	secondWord = word[int(halfOfTotal):]

	// reverse half of word
	for j := len(secondWord) - 1; j >= 0; j-- {
		reverse.WriteString(string(secondWord[j]))
	}
	if firstWord != reverse.String() {
		return false
	}
	return true
}

func LeapYear(firstYear int, secondYear int) {

	// check if
	if firstYear < 1000 || firstYear > 9999 || secondYear < 1000 || secondYear > 9999 {
		panic("wrong year")
	}

	if firstYear >= secondYear {
		panic("wrong year range")
	}

	for i := firstYear; i <= secondYear; i++ {
		if i%400 == 0 || i%4 == 0 && i%100 != 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

}

func ReverseWords(words string) {
	var reversedWords, result bytes.Buffer
	upperIndexList := map[int][]interface{}{}
	wordsSplit := strings.Split(words, " ")
	totalSplit := len(wordsSplit)

	for i, word := range wordsSplit {
		var reverseWord bytes.Buffer
		for j := len(word) - 1; j >= 0; j-- {
			w := string(word[j])
			if unicode.IsUpper(rune(word[j])) {
				upperIndexList[i] = append(upperIndexList[i], j)
			}
			reverseWord.WriteString(w)
		}
		reversedWords.WriteString(reverseWord.String() + " ")
	}

	// set upper letter
	for i, word := range strings.Split(strings.ToLower(reversedWords.String()), " ") {
		for j, char := range word {
			upper := false
			for _, value := range upperIndexList[i] {
				if value == j {
					upper = true
				}
			}
			if upper {
				result.WriteString(strings.ToUpper(string(char)))
			} else {
				result.WriteString(string(char))
			}
		}
		if totalSplit-1 > i {
			result.WriteString(" ")
		}
	}
	fmt.Println("\"" + words + "\"")
	fmt.Println("\"" + result.String() + "\"")
}

func NearestFibonacci(data []int) {
	// sum all list int
	var max, distance int
	//var fibSelected int
	for _, value := range data {
		max += value
	}

	// fibonacci
	fibData := []int{1, 1}
	for i := 1; true; i++ {
		stop := false
		fibValue := fibData[i-1] + fibData[i]
		fibData = append(fibData, fibValue)
		currentDistance := max - fibValue

		if fibValue >= max {
			numberSplit := strings.Split(strconv.Itoa(currentDistance), "-")
			if len(numberSplit) == 2 {
				currentDistance = int(^uint32(int32(currentDistance) - 1))
			}
			stop = true
		}

		if i != 1 {
			if distance > currentDistance {
				distance = currentDistance
				//fibSelected = fibValue
			}
		} else {
			distance = currentDistance
			//fibSelected = fibValue
		}

		if stop {
			break
		}
	}

	//
	//fmt.Println("input result :" ,max)
	//fmt.Println("near from : ", fibSelected)
	//fmt.Println("distance : ", distance)
	//fmt.Println("fibData ", fibData)
	fmt.Println(distance)
}

func FizzBuzz(n int) {
	var result []interface{}
	for i := 1; i < n; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, i)
		}
	}
	fmt.Println(result)
}

func main() {
	var kata string
	kata = "Radar"
	if Palindrome(kata) {
		fmt.Println(kata + " \t\t\t>> palindrome")
	} else {
		fmt.Println(kata + " \t\t\t>> not palindrome")
	}
	kata = "Malam"
	if Palindrome(kata) {
		fmt.Println(kata + " \t\t\t>> palindrome")
	} else {
		fmt.Println(kata + " \t\t\t>> not palindrome")
	}
	kata = "Kasur ini rusak"
	if Palindrome(kata) {
		fmt.Println(kata + " \t>> palindrome")
	} else {
		fmt.Println(kata + " \t>> not palindrome")
	}
	kata = "Ibu Ratna antar ubi"
	if Palindrome(kata) {
		fmt.Println(kata + " \t>> palindrome")
	} else {
		fmt.Println(kata + " \t>> not palindrome")
	}
	kata = "Malas"
	if Palindrome(kata) {
		fmt.Println(kata + " \t\t\t>> palindrome")
	} else {
		fmt.Println(kata + " \t\t\t>> not palindrome")
	}
	kata = "Makan nasi goreng"
	if Palindrome(kata) {
		fmt.Println(kata + " \t>> palindrome")
	} else {
		fmt.Println(kata + " \t>> not palindrome")
	}
	kata = "Balonku ada lima"
	if Palindrome(kata) {
		fmt.Println(kata + " \t>> palindrome")
	} else {
		fmt.Println(kata + " \t>> not palindrome")
	}
	fmt.Println("---------------------")
	LeapYear(1900, 2020)
	fmt.Println("---------------------")
	ReverseWords("I am A Great human")
	fmt.Println("---------------------")
	NearestFibonacci([]int{15, 1, 3})
	fmt.Println("---------------------")
	FizzBuzz(16)
}
