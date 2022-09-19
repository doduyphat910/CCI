package main

import "testing"

func TestIsUnique(t *testing.T) {
	testCases := []struct {
		name string
		arg  string
		want bool
	}{
		{
			name: "unique",
			arg:  "tes",
			want: true,
		},
		{
			name: "not unique",
			arg:  "testttt",
			want: false,
		},
		{
			name: "not unique",
			arg:  "test",
			want: false,
		},
		{
			name: "not unique",
			arg:  "",
			want: true,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			result := isUnique(test.arg)
			if result != test.want {
				t.Errorf("fail. result:%v, want:%v", result, test.want)
			}
		})
	}
}

func TestPermutation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}

	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "have permutation s1",
			args: args{
				s1: "ab",
				s2: "adb",
			},
			want: "ab",
		},
		{
			name: "have permutation s2",
			args: args{
				s1: "abc",
				s2: "ac",
			},
			want: "ac",
		},
		{
			name: "not have permutation",
			args: args{
				s1: "abc",
				s2: "a123",
			},
			want: "",
		},
		{
			name: "not have permutation",
			args: args{
				s1: "",
				s2: "a123",
			},
			want: "",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			permutationS := checkPermutation(testCase.args.s1, testCase.args.s2)
			if permutationS != testCase.want {
				t.Errorf("fail. Result:%s. Want:%s", permutationS, testCase.want)
			}
		})
	}
}

func TestReplaceString(t *testing.T) {
	type args struct {
		s      string
		length int
	}
	testCases := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base",
			args: args{
				s:      "Mr John Smith   ",
				length: 13,
			},
			want: "Mr%20John%20Smith",
		},
		{
			name: "",
			args: args{
				s:      "Mr John Smith   ",
				length: 15,
			},
			want: "Mr%20John%20Smith%20%20",
		},
	}

	for _, test := range testCases {
		resultS := replaceString(test.args.s, test.args.length)
		if resultS != test.want {
			t.Errorf("Fail: result:%s, want:%s", resultS, test.want)
		}
	}
}

func TestIsPalindrome(t *testing.T) {
	type args struct {
		palindrome   string
		permutationS string
	}
	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "valid",
			args: args{
				palindrome:   "Tact Coa",
				permutationS: "taco cat",
			},
			want: true,
		},
		{
			name: "valid2",
			args: args{
				palindrome:   "Tact Coa",
				permutationS: "atco cta",
			},
			want: true,
		},
		{
			name: "invalid1",
			args: args{
				palindrome:   "Tact Coa",
				permutationS: "atco ctaa",
			},
			want: false,
		},
		{
			name: "invalid2",
			args: args{
				palindrome:   "Tact Coaaa",
				permutationS: "atco cta",
			},
			want: false,
		},
		{
			name: "invalid",
			args: args{
				palindrome:   "Tact Coa",
				permutationS: "test",
			},
			want: false,
		},
	}

	for _, testCase := range testCases {
		result := isPalindrome(testCase.args.palindrome, testCase.args.permutationS)
		if result != testCase.want {
			t.Errorf("fail. want=%v, result=%v", testCase.want, result)
		}
	}
}

func TestIsEdit(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1 remove",
			args: args{
				s1: "pale",
				s2: "ple",
			},
			want: true,
		},
		{
			name: "1 insert",
			args: args{
				s1: "pales",
				s2: "pale",
			},
			want: true,
		},
		{
			name: "1 edit",
			args: args{
				s1: "pale",
				s2: "bale",
			},
			want: true,
		},
		{
			name: "2 edit",
			args: args{
				s1: "pale",
				s2: "bake",
			},
			want: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isEdit(testCase.args.s1, testCase.args.s2)
			if result != testCase.want {
				t.Errorf("fail. want=%v, result=%v", testCase.want, result)
			}
		})
	}
}

func TestCompress(t *testing.T) {
	testCases := []struct {
		name      string
		originalS string
		want      string
	}{
		{
			name:      "Example",
			originalS: "aabcccccaaa",
			want:      "a2b1c5a3",
		},
		{
			name:      "Example",
			originalS: "aabccccca",
			want:      "a2b1c5a1",
		},
		{
			name:      "Example",
			originalS: "aabccccc",
			want:      "a2b1c5",
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := compress(testCase.originalS)
			if result != testCase.want {
				t.Errorf("fail. want=%v, result=%v", testCase.want, result)
			}
		})
	}
}

func TestStringRotation(t *testing.T) {
	type args struct {
		s1 string
		s2 string
	}
	testCases := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "ex1",
			args: args{
				s1: "waterbottle",
				s2: "erbottlewat",
			},
			want: true,
		},
		{
			name: "ex2",
			args: args{
				s1: "waterbottle",
				s2: "erbottlawat",
			},
			want: false,
		},
		{
			name: "ex2",
			args: args{
				s1: "waterbottleee",
				s2: "erbottlawat",
			},
			want: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isSubString(testCase.args.s1, testCase.args.s2)
			if result != testCase.want {
				t.Errorf("fail. want=%v, result=%v", testCase.want, result)
			}
		})
	}
}
