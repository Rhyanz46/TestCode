package main

import (
	"bytes"
	"fmt"
	"strings"
	"unicode"
)

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
			fmt.Println(i)
		}
	}

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

func main() {
	kata := "afa"
	if Palindrome(kata) {
		fmt.Println(kata + " >> palindrom")
	} else {
		fmt.Println(kata + " >> not palindrom")
	}

	LeapYear(1980, 2000)
	ReverseWords("arN KeRen Banget aa")
}
