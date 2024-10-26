package main

import (
	"fmt"
)

func encodeCaesar(plainText string, shift int) string {
	cipherText := ""

	for _, char := range plainText {
		if char >= 'A' && char <= 'Z' {
			cipherText += string((char - 'A' + rune(shift)) % 26 + 'A')
		} else {
			cipherText += string(char)
		}
	}
	return cipherText
}

func main() {
	plainText := "ZOLZLYLQ ZEBTBA JV JXQZE KLQBP"
	shift := 3

	encodedText := encodeCaesar(plainText, shift)
	fmt.Println("Encoded Text:", encodedText)
}

