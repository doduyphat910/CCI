package main

import "strings"

//1.1 Is Unique: Implement an algorithm to determine if a string has all unique characters.
//What if you cannot use additional data structures
func isUnique(s string) bool {
	var tempS string
	for i := range s {
		subString := string(s[i])
		if subString != tempS {
			tempS = subString
			continue
		}
		return false
	}
	return true
}

//1.2 Check Permutation: Give two strings, write a method to decide if one is a permutation of the other
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

//1.3 URLify: Write a method to replace all spaces in a string with '%20'. You may assume that the string
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

//1.4 Palindrome Permutation. Given a string, write a function to check if it is a permutation of a palindrome.
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
