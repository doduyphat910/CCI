package main

import (
	"reflect"
	"testing"
)

func TestRouteBetweenNodes(t *testing.T) {
	type args struct {
		graph Node
		a     string
		b     string
	}
	testCases := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test case 1",
			args: args{
				graph: Node{
					Name: "a",
					Node: []Node{
						{
							Name: "c",
							Node: []Node{
								{Name: "b"},
							},
						},
						{
							Name: "d",
							Node: []Node{
								{Name: "b"},
							},
						},
					},
				},
				a: "a",
				b: "b",
			},
			want: []string{"a", "c", "b"},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := routeBetweenNodes(testCase.args.graph, testCase.args.a, testCase.args.b)
			if !reflect.DeepEqual(result, testCase.want) {
				t.Errorf("fail. result:%v, want:%v", result, testCase.want)
			}
		})
	}
}

func TestMinimalTree(t *testing.T) {
	testCases := []struct {
		name string
		arg  []int
		want *NodeInt
	}{
		{
			name: "test 1",
			arg:  []int{1, 2, 3, 4, 5, 6, 7},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := minimalTree(testCase.arg)
			if !reflect.DeepEqual(result, testCase.want) {
				t.Errorf("fail. result:%v, want:%v", result, testCase.want)
			}
		})
	}
}

func TestListOfDepths(t *testing.T) {
	testCases := []struct {
		name string
		arg  *NodeInt
		want []LinkList
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
			want: []LinkList{
				{root: &LinkListNode{Value: 5}},
				{root: &LinkListNode{Value: 3, Next: &LinkListNode{Value: 4}}},
				{root: &LinkListNode{Value: 7, Next: &LinkListNode{Value: 6}}},
				{root: &LinkListNode{Value: 8}},
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			result := listOfDepths(testCase.arg)
			if !reflect.DeepEqual(result, testCase.want) {
				t.Errorf("fail. result:%v, want:%v", result, testCase.want)
			}
		})
	}
}
