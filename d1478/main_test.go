package main

import "testing"

type TestCase struct {
	in   *Node
	want string
}

func TestFindDeepestNode(t *testing.T) {
	cases := []TestCase{
		{&Node{
			Val: "A",
			Left: &Node{
				Val: "B",
				Left: &Node{
					Val: "D",
				},
			},
			Right: &Node{
				Val: "C",
			},
		},
			"D"},
		{
			&Node{
				Val: "1",
				Left: &Node{
					Val: "2",
					Left: &Node{
						Val: "4",
					},
					Right: &Node{
						Val: "5",
					},
				},
				Right: &Node{
					Val: "3",
					Left: &Node{
						Val: "6",
					},
					Right: &Node{
						Val: "7",
						Right: &Node{
							Val: "8",
						},
					},
				},
			},
			"8"},
		{
			&Node{
				Val: "1",
				Left: &Node{
					Val: "2",
				},
				Right: &Node{
					Val: "3",
					Left: &Node{
						Val: "6",
					},
				},
			},
			"6",
		},
	}
	for _, tt := range cases {
		got := FindDeepestNode(tt.in)
		if got.Val != tt.want {
			t.Errorf("FindDeepestNode(%v) = %v, want %v", tt.in, got, tt.want)
		}

	}
}
