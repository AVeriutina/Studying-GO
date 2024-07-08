package main

import (
	"fmt"
	"unicode"
)

func isVowel(ch rune) {
	ch = unicode.ToLower(ch)
	set := map[rune]struct{}{
		'а': {},
		'у': {},
		'о': {},
		'и': {},
		'э': {},
		'ы': {},
		'я': {},
		'ю': {},
		'е': {},
		'ё': {},
	}
	_, ok := set[ch]
	if ok {
		fmt.Println("Это гласная буква")
	} else {
		fmt.Println("Это согласная буква")
	}

}

func main() {
	var ch rune
	fmt.Scanf("%c", &ch)
	isVowel(ch)
}
