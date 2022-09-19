package main

import (
	"reflect"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	type args struct {
		row    int
		column int
	}

	testCases := []struct {
		name string
		args args
		want [][]int8
	}{
		{
			name: "case 1",
			args: args{
				row:    3,
				column: 4,
			},
			want: [][]int8{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := zeroMatrix(testCase.args.row, testCase.args.column)
			if !reflect.DeepEqual(result, testCase.want) {
				t.Errorf("fail. Result:%v. Want:%v", result, testCase.want)
			}
		})
	}
}
