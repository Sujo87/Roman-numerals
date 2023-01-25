package main

import (
	"fmt"
	"strings"
)

func main() {

	var s string

	fmt.Println("enter roman number")
	fmt.Scan(&s)

	if len(s) >= 1 && len(s) <= 15 {
		s = strings.ToUpper(s)

		fmt.Println(s)
		fmt.Println(romanToInt(s))

	} else {
		fmt.Println("Invalid input")

	}

}

func romanToInt(s string) int {

	var rezultat int

	if strings.Contains(s, "IV") {
		rezultat += 4
		s = strings.ReplaceAll(s, "IV", "")
	}

	//Doube digit
	if strings.Contains(s, "IL") {
		rezultat += 40
		s = strings.ReplaceAll(s, "IL", "")
	}

	if strings.Contains(s, "IX") {
		rezultat += 9
		s = strings.ReplaceAll(s, "IX", "")
	}

	if strings.Contains(s, "XC") {
		rezultat += 90
		s = strings.ReplaceAll(s, "XC", "")
	}

	if strings.Contains(s, "CD") {
		rezultat += 400
		s = strings.ReplaceAll(s, "CD", "")
	}

	if strings.Contains(s, "CM") {
		rezultat += 900
		s = strings.ReplaceAll(s, "CM", "")
	}

	//Single digit
	if strings.Contains(s, "I") {
		rezultat = rezultat + (strings.Count(s, "I") * 1)
		s = strings.ReplaceAll(s, "I", "")
	}

	if strings.Contains(s, "X") {
		rezultat = rezultat + (strings.Count(s, "X") * 10)
		s = strings.ReplaceAll(s, "X", "")
	}

	if strings.Contains(s, "L") {
		rezultat = rezultat + (strings.Count(s, "L") * 40)
		s = strings.ReplaceAll(s, "L", "")
	}

	if strings.Contains(s, "C") {
		rezultat = rezultat + (strings.Count(s, "C") * 100)
		s = strings.ReplaceAll(s, "C", "")
	}

	if strings.Contains(s, "D") {
		rezultat = rezultat + (strings.Count(s, "D") * 400)
		s = strings.ReplaceAll(s, "D", "")
	}

	if strings.Contains(s, "M") {
		rezultat = rezultat + (strings.Count(s, "M") * 1000)
		s = strings.ReplaceAll(s, "M", "")
	}

	return rezultat
}
