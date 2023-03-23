package main

import (
	"fmt"
	"strconv"
	"unicode"
)

func unpackString(strings string) string {
	runeString := []rune(strings)
	stringEdited := ""
	var k int

	if len(runeString) == 0 || unicode.IsDigit(runeString[0]) {
		return ""
	}

	for i := 1; i < len(runeString); i++ {
		if unicode.IsLetter(runeString[i-1]) && unicode.IsDigit(runeString[i]) {
			s, _ := strconv.Atoi(string(runeString[i]))
			for j := 0; j < s; j++ {
				stringEdited += string(runeString[i-1])
				k += j + 1
			}
		} else if unicode.IsLetter(runeString[i-1]) == true {
			stringEdited += string(runeString[i-1])
			k += 1
		}
	}

	if unicode.IsLetter(runeString[len(runeString)-1]) == true {
		stringEdited += string(runeString[len(runeString)-1])
	}

	return stringEdited
}

func main() {
	firstString := "a4bc2d5e"
	fmt.Println(unpackString(firstString))

	secondString := "abcd"
	fmt.Println(unpackString(secondString))

	thirdString := "45"
	fmt.Println(unpackString(thirdString))

	fourthString := ""
	fmt.Println(unpackString(fourthString))
}
