/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package dagger_test

import (
	"log"
	"reflect"
	"testing"

	"github.com/iainjreid/dagger"
)

type TestNode interface {
}

type TestNodeImpl struct {
	name  string
	left  []TestNode
	right []TestNode
}

func NewTestNode(name string) *dagger.Node[string, TestNode] {
	return dagger.NewNode(func(_ string, left, right []TestNode) TestNode {
		return &TestNodeImpl{
			name,
			left,
			right,
		}
	})
}

type DaggerNodeImpl struct {
	name string
}

func NewTestDaggerNode(name string) *dagger.Node[string, TestNode] {
	return dagger.NewNode(func(_ string, _, _ []TestNode) TestNode {
		return DaggerNodeImpl{
			name,
		}
	})
}

// TestAppend calls [Builder.Append], to ensure that it correctly adds a new
// child to the parent node.
func TestAppend(t *testing.T) {
	subject := NewTestNode("parent").Append(NewTestNode("child"))

	expected := &TestNodeImpl{
		name: "parent",
		left: []TestNode{
			&TestNodeImpl{
				name:  "child",
				left:  []TestNode{},
				right: []TestNode{},
			},
		},
		right: []TestNode{},
	}

	tree := dagger.NewTree[TestNode]()
	actual := subject.BuildTo("68yvwz", tree)

	if !reflect.DeepEqual(actual, expected) {
		log.Println(actual)
		t.Fatal("result should be equal to expected output")
	}
}

func TestAppendMany(t *testing.T) {
	subject := NewTestNode("parent").Append(
		NewTestNode("child1"),
		NewTestNode("child2"),
		NewTestNode("child3"),
	)

	expected := &TestNodeImpl{
		name: "parent",
		left: []TestNode{
			&TestNodeImpl{
				name:  "child1",
				left:  []TestNode{},
				right: []TestNode{},
			},
			&TestNodeImpl{
				name:  "child2",
				left:  []TestNode{},
				right: []TestNode{},
			},
			&TestNodeImpl{
				name:  "child3",
				left:  []TestNode{},
				right: []TestNode{},
			},
		},
		right: []TestNode{},
	}

	tree := dagger.NewTree[TestNode]()
	actual := subject.BuildTo("68yvwz", tree)

	if tree.Len() != 3 {
		t.Fatal("incorrect tree size")
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Fatal("result should be equal to expected output")
	}
}

func TestAppendManyNested(t *testing.T) {
	subject := NewTestNode("parent").Append(
		NewTestNode("child1").Append(
			NewTestNode("child2"),
			NewTestNode("child3"),
		),
	)

	expected := &TestNodeImpl{
		name: "parent",
		left: []TestNode{
			&TestNodeImpl{
				name: "child1",
				left: []TestNode{
					&TestNodeImpl{
						name:  "child2",
						left:  []TestNode{},
						right: []TestNode{},
					},
					&TestNodeImpl{
						name:  "child3",
						left:  []TestNode{},
						right: []TestNode{},
					},
				},
				right: []TestNode{},
			},
		},
		right: []TestNode{},
	}

	tree := dagger.NewTree[TestNode]()
	actual := subject.BuildTo("68yvwz", tree)

	if tree.Len() != 3 {
		t.Fatal("incorrect tree size")
	}

	if !reflect.DeepEqual(actual, expected) {
		log.Println(actual)
		t.Fatal("result should be equal to expected output")
	}
}

func TestAppendDeeplyNested(t *testing.T) {
	subject := NewTestNode("parent").Append(
		NewTestNode("child1").Append(
			NewTestDaggerNode("child2"),
			NewTestDaggerNode("child3"),
		),
		NewTestNode("child4").Append(
			NewTestDaggerNode("child5"),
			NewTestDaggerNode("child6"),
		),
		NewTestNode("child7").Append(
			NewTestDaggerNode("child8"),
			NewTestDaggerNode("child9"),
		),
	)

	expected := &TestNodeImpl{
		name: "parent",
		left: []TestNode{
			&TestNodeImpl{
				name: "child1",
				left: []TestNode{
					DaggerNodeImpl{
						name: "child2",
					},
					DaggerNodeImpl{
						name: "child3",
					},
				},
				right: []TestNode{},
			},
			&TestNodeImpl{
				name: "child4",
				left: []TestNode{
					DaggerNodeImpl{
						name: "child5",
					},
					DaggerNodeImpl{
						name: "child6",
					},
				},
				right: []TestNode{},
			},
			&TestNodeImpl{
				name: "child7",
				left: []TestNode{
					DaggerNodeImpl{
						name: "child8",
					},
					DaggerNodeImpl{
						name: "child9",
					},
				},
				right: []TestNode{},
			},
		},
		right: []TestNode{},
	}

	tree := dagger.NewTree[TestNode]()
	actual := subject.BuildTo("68yvwz", tree)

	if tree.Len() != 9 {
		t.Fatal("incorrect tree size")
	}

	if !reflect.DeepEqual(actual, expected) {
		log.Println(actual)
		t.Fatal("result should be equal to expected output")
	}
}

// TestLift calls [Builder.Lift], to ensure that dynamic nodes can be added
// using the provided build context.
// func TestLift(t *testing.T) {
// 	subject := NewTestNode("parent").Lift(func(str string) dagger.Buildable[string, &TestNodeImpl] {
// 		return NewTestNode(str)
// 	})

// 	expected := &TestNodeImpl{
// 		name: "parent",
// 		left: []&TestNodeImpl{
// 			{
// 				name:  "x8azmu",
// 				left:  []TestNode{},
// 				right: []TestNode{},
// 			},
// 		},
// 		right: []TestNode{},
// 	}

// 	actual, _ := subject.BuildOut("x8azmu")

// 	if !reflect.DeepEqual(actual, expected) {
// 		t.Fatal("result should be equal to expected output")
// 	}
// }
