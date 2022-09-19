package main

import (
	"fmt"
	"strings"
)

// 1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters.
// What if you cannot use additional data structures

//	func isUnique(s string) bool {
//		var tempS string
//		for i := range s {
//			subString := string(s[i])
//			if subString != tempS {
//				tempS = subString
//				continue
//			}
//			return false
//		}
//		return true
//	}

// O(N^2)
func isUnique(s string) bool {
	var tempS strings.Builder
	for i := range s {
		subString := string(s[i])
		if len(tempS.String()) == 0 {
			tempS.WriteString(subString)
			continue
		}

		for j := range tempS.String() {
			subTempS := string(tempS.String()[j])
			if subString != subTempS {
				tempS.WriteString(subString)
				continue
			}
			return false
		}
	}
	return true
}

// 1.2 Check Permutation: Give two strings, write a method to decide if one is a permutation of the other
func checkPermutation(s1, s2 string) string {
	lengthS1 := len(s1)
	lengthS2 := len(s2)
	if lengthS1 == 0 || lengthS2 == 0 {
		return ""
	}

	if lengthS1 >= lengthS2 {
		sByNumOccurMap := mapStringByNumOccur(s1)
		return checkPermutationByMap(sByNumOccurMap, s2)
	}

	// check s2
	sByNumOccurMap := mapStringByNumOccur(s2)
	return checkPermutationByMap(sByNumOccurMap, s1)
}

func mapStringByNumOccur(s string) map[string]uint {
	var sByNumOccurMap = make(map[string]uint)
	for i := range s {
		numOccur, found := sByNumOccurMap[string(s[i])]
		if !found {
			sByNumOccurMap[string(s[i])] = 1
			continue
		}
		numOccur++
		sByNumOccurMap[string(s[i])] = numOccur
	}

	return sByNumOccurMap
}

func checkPermutationByMap(sByNumOccurMap map[string]uint, s string) string {
	for i := range s {
		numOccur, found := sByNumOccurMap[string(s[i])]
		if !found {
			return ""
		}
		if numOccur == 0 {
			return ""
		}
		numOccur--
	}
	return s
}

// 1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
// has sufficient space at the end to hold the additional characters, and that you are given the "true"
// length of the string. (Note: if implementing in Java, please use a character array so that you can
// perform this operation in place)
func replaceString(s string, length int) string {
	var newS strings.Builder
	for i := range s {
		tempS := string(s[i])
		if tempS == " " {
			tempS = "%20"
		}
		newS.WriteString(tempS)
		if i+1 == length {
			return newS.String()
		}
	}
	return newS.String()
}

// 1.4 Palindrome Permutation. Given a string, write a function to check if it is a permutation of a palindrome.
// A palindrome is a word or phrase that is the same forwards and backwards. A permutation
// is a rearrangement of letters. The palindrome does not need to be limited to just dictionary words
func isPalindrome(palindrome string, permutationS string) bool {
	palindrome = strings.ToLower(palindrome)
	palindByNumOccurMap := mapStringByNumOccur(palindrome)
	for i := range permutationS {
		numOccur, found := palindByNumOccurMap[string(permutationS[i])]
		if !found {
			return false
		}
		if numOccur == 0 {
			return false
		}
		numOccur--
		palindByNumOccurMap[string(permutationS[i])] = numOccur
	}
	for i := range palindByNumOccurMap {
		if palindByNumOccurMap[i] != 0 {
			return false
		}
	}
	return true
}

// 1.5 One Away: There are three types of edits that can be performed on strings: insert a character,
// remove a character, or replace a character. Given two strings, write a function to check if they are
// one edit (or zero edits) away

func isEdit(s1 string, s2 string) bool {
	var diffNum uint
	if len(s1) > len(s2) {
		diffNum = countDiff(s1, s2)
	} else {
		diffNum = countDiff(s2, s1)
	}

	if diffNum > 1 {
		return false
	}
	return true
}

func countDiff(s1 string, s2 string) uint {
	diffCount := 0
	for i := range s1 {
		if diffCount > 1 {
			break
		}
		if len(s1) == len(s2) {
			if s1[i] != s2[i] {
				diffCount++
			}
			continue
		}

		if i+diffCount > len(s1)-1 {
			break
		}
		if i+1 > len(s2) {
			diffCount++
			continue
		}

		if s1[i+diffCount] != s2[i] {
			diffCount++
		}
	}

	return uint(diffCount)
}

// 1.6 String Compression: Implement a method to perform basic string compression using the counts
// of repeated characters. For example, the string aabcccccaaa would become a2b1c5a3. If the
// "compressed" string would not become smaller than original string, your method should return
// the original string. You can assume the string has only uppercase and lowercase letter (a-z)

func compress(originalS string) string {
	var (
		compressedS strings.Builder
		repeatNum   = 1
	)

	for i := range originalS {
		if i == 0 {
			compressedS.WriteString(string(originalS[i]))
			continue
		}

		if originalS[i] == originalS[i-1] {
			repeatNum++
		} else {
			compressedS.WriteString(fmt.Sprintf("%v", repeatNum))
			compressedS.WriteString(string(originalS[i]))
			repeatNum = 1
		}

		if i+1 == len(originalS) {
			compressedS.WriteString(fmt.Sprintf("%v", repeatNum))
		}
	}

	return compressedS.String()
}

// 1.9 String Rotation. Assume you have a method isSubStringwhich checks if one word is a substring
// of another. Given two strings, s1 and s2, write code to check if s2 is a rotation of s1 using only one
// call to isSubstring (e.g., "waterbottle" is a rotation of "erbottlewat").

func isSubString(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	s1ByNumMap := mapStringByNumOccur(s1)
	s2ByNumMap := mapStringByNumOccur(s2)

	for k := range s1ByNumMap {
		if s1ByNumMap[k] != s2ByNumMap[k] {
			return false
		}
	}

	return true
}
