// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-
//
//	ROT
//	Author: Ondřej Tuček (ondrejtucek9@gmail.com)				27.11.2023
//
// -----------------------------------------------------------------------
//
//	Desc. ROT is ciphering algorithm that is based on the caesar cipher
//	This program allows you to decode and encode payload in ROT.
//
// =-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-=-==-=-=-=-=-=-=-=-=-=-=-=-
package rot

import (
	"fmt"
	"strings"
)

func getEnglishAlphabet() map[rune]int {
	var alphabet = make(map[rune]int, 25)
	var decASCII int = 65

	for i := 1; i <= 26; i++ {
		alphabet[rune(decASCII)] = i
		decASCII += 1
	}

	return alphabet
}

func Encode(payload string, shift int) string {
	alphabet := getEnglishAlphabet()
	runeStr := []rune(strings.ToUpper(payload))

	for i, value := range runeStr {
		index, ok := alphabet[value]

		if !ok {
			fmt.Printf("warning: incompatible character %c with alphabet\n", value)
			continue
		}

		if index+shift > 26 {
			amount := (index + shift) - 26
			runeStr[i] = 64 + rune(amount)
			continue
		}

		runeStr[i] += rune(shift)
	}

	return string(runeStr)
}

func Decode(payload string, shift int) string {
	alphabet := getEnglishAlphabet()
	str := []rune(strings.ToUpper(payload))

	for i, value := range str {
		index, ok := alphabet[value]

		if !ok {
			fmt.Printf("warning: incompatible character %c with alphabet\n", value)
			continue
		}

		amount := index - shift

		if amount < 0 {
			str[i] = rune(90 - amount*-1)
			continue
		}

		str[i] -= rune(shift)
	}

	return string(str)
}
