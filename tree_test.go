package main

import (
	"testing"
)

func TestIsBalanced(t *testing.T) {
	testCases := []struct {
		name string
		arg  *NodeInt
		want bool
	}{
		{
			name: "Case1",
			arg: &NodeInt{
				Value: 5,
				LeftNode: &NodeInt{
					Value:    3,
					LeftNode: &NodeInt{Value: 7},
					RightNode: &NodeInt{
						Value:     6,
						RightNode: &NodeInt{Value: 8},
					},
				},
				RightNode: &NodeInt{Value: 4},
			},
			want: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := isBalanced(testCase.arg)
			if result != testCase.want {
				t.Errorf("fail. result:%v, want:%v", result, testCase.want)
			}
		})
	}
}

func TestIsBST(t *testing.T) {
	testCases := []struct {
		name string
		arg  *NodeInt
		want bool
	}{
		// {
		// 	name: "Case1",
		// 	arg: &NodeInt{
		// 		Value: 5,
		// 		LeftNode: &NodeInt{
		// 			Value:    3,
		// 			LeftNode: &NodeInt{Value: 7},
		// 			RightNode: &NodeInt{
		// 				Value:     6,
		// 				RightNode: &NodeInt{Value: 8},
		// 			},
		// 		},
		// 		RightNode: &NodeInt{Value: 4},
		// 	},
		// 	want: false,
		// },
		{
			name: "Case2",
			arg: &NodeInt{
				Value: 7,
				LeftNode: &NodeInt{
					Value:    5,
					LeftNode: &NodeInt{Value: 3},
					// RightNode: &NodeInt{
					// 	Value:     6,
					// 	RightNode: &NodeInt{Value: 8},
					// },
				},
				RightNode: &NodeInt{Value: 8},
			},
			want: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			isBst, _ := isBST(testCase.arg)
			if isBst != testCase.want {
				t.Errorf("fail. result:%v, want:%v", isBst, testCase.want)
			}
		})
	}
}
