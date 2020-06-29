package main

import "fmt"

/*
 * The problem statement
 *
 * Time Limit: 45min + 15min
 *
 * We were working with rhyme schemes. To tell if the
 * two words rhyme, we have to compare their rhyme-
 * patterns. like
 * We need your help to write a program that can find
 * out the rhyme-pattern of any word.
 *
 * Vowel
 * A vowel is any of: `a`, `e`, `i`, `o`, `u` or `y`.
 * NOTE: we consider `​y`​ as a vowel too, as long as
 * it is not at the start or end of a word. So, as an
 * example, the `y`​ in `rhythm`​ is considered a vowel
 *
 * ​Rhyme-pattern​
 * A ​rhyme-pattern​ is a substring of a word such that:
 *   1. The word ends with that substring.
 *   2. The first letter of the substring is always a
 *       vowel.
 *   3. The substring contains exactly one contiguous
 *       string of vowel(s).
 *   4. The substring must either be the whole word,
 *       or the letter immediately preceding the
 *       start of the substring must be a non-vowel.
 *
 * For example,
 *  the rhyme-pattern​ of `star` would be `ar`,
 *  the rhyme-pattern​ of `rainbow` would be `ow`,
 *  the rhyme-pattern​ of `noise` would be `e`,
 *  the rhyme-pattern​ of `s​pying​` would be `​ying​`,
 *  and the rhyme-pattern​ of `​all​` would be `​all​`.
 *
 * Input will:
 *  1. always be in lower case
 *  2. always have vowels
 *  3. have no other character than [a-z]
 *
 * Task: you need to implement the function
 * `GetRhymePattern` below to return the rhyme-pattern
 * as described.
 *
 * [Bonus marks for good commenting, brevity, and modularity]
 */

func main() {

	//fmt.Println(GetRhymePattern("star"))
	//fmt.Println(GetRhymePattern("s​pying​"))
	s := GetRhymePattern("​all​")
	fmt.Println(s)

}

func isVovel(ch rune, index, length int) bool {
	if ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u' {
		return true
	}
	if ch == 'y' {
		if index == 0 || index == length-1 {
			return false
		}
		return true
	}
	return false
}

func GetRhymePattern(word string) string {
	length := len(word)

	chArr := []rune(word)

	i := 0
	for i < len(chArr) {
		if isVovel(chArr[i], i, length) {
			temp := i
			vowel := true
			for temp < len(word) {
				if isVovel(chArr[temp], temp, length) {
					if !vowel {
						break
					}
					temp++
					continue
				}
				vowel = false
				temp++
			}
			if temp == length {
				if isVovel(chArr[temp-1], temp-1, length) && !vowel {
					return string(chArr[temp-1])
				} else {
					return string(chArr[i:])
				}
			}

			i = temp
			continue
		}

		i++
	}
	return ""
}
