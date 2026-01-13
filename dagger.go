/*
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

// Package dagger implements a functional toolkit designed to build hierarchical
// and tree-like structures in a composable manner.
//
// It's built using functional programming concepts that some might recognise
// from the Reader monad; however, to ensure the interface is clear, concise,
// and, most importantly, simple to use, the method names have been chosen to
// avoid such specific terminology.
//
// Included below is a sample program that builds a small tree with a variety
// of different nodes:
//
//	package sample
//
// TODO: Include code sample
package dagger

type Storage[T any] interface {
	Write(pos int, nodes ...[]T)
}

type Marker[T any] struct {
	Pos int
	Len int

	// Max specifies the maximum allowable number of nodes that can be stored
	// within the bounds of the Marker.
	Max int
}

func (m *Marker[T]) Write(storage Storage[T], nodes ...[]T) {
	storage.Write(m.Pos, nodes...)
}

// Grow
func (m *Marker[T]) Grow(size int) {
	m.Len += size

	if m.Len > m.Max {
		panic("maximum marker length exceeded")
	}
}

type Buildable[T, U any] interface {
	Build(T) (U, Tree[U])
	BuildTo(T, Tree[U]) U
}

type Tree[T any] interface {
	Grow(size int)
	Len() int
	Nodes() []T
	Push(marker TreeMarker, pos *int, items ...T)
	Read(marker TreeMarker) []T
	Reserve(len int) TreeMarker
}

type Node[T, U any] struct {
	build func(T, Tree[U], TreeMarker, *int, *int) U

	leftPos  int
	rightPos int

	// TODO: Remove these fields in favour of a single count, and use the underlying tree to handle node grouping.
	noOfLeft  int
	noOfRight int
}

type Dagger[T, U any] struct {
	build func(T, Tree[U]) U
}

type Leaf[T, U any] struct {
	build func(T, Tree[U]) U
}

type IdentityLeaf[T, U any] struct {
	identity U
}

func NewLeaf[T, U any](build func(T, Tree[U]) U) *Leaf[T, U] {
	return &Leaf[T, U]{
		build: build,
	}
}

func (b *Leaf[T, U]) Build(data T) (U, Tree[U]) {
	tree := NewTree[U]()
	result := b.BuildTo(data, tree)
	return result, tree
}

func (b *Leaf[T, U]) BuildTo(data T, tree Tree[U]) U {
	return b.build(data, tree)
}

func NewNode[T, U any](build func(T, []U, []U) U) *Node[T, U] {
	return &Node[T, U]{
		build: func(t T, tree Tree[U], marker TreeMarker, leftEnd, rightEnd *int) U {
			list := tree.Read(marker)
			return build(t, list[:*leftEnd], list[*leftEnd:*rightEnd])
		},
	}
}

type BasicTree[T any] struct {
	nodes []T
	pos   int
}

func NewTree[T any]() *BasicTree[T] {
	return &BasicTree[T]{
		nodes: make([]T, 0),
	}
}

// Tree markers define the memory layout for nodes

type TreeMarker struct {
	pos int
	len int
}

func near(n uint32) uint32 {
	n--
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	n++

	return n
}

func (t *BasicTree[T]) Grow(size int) {
	len := len(t.nodes)

	if near(uint32(len)) == near(uint32(len+size)) {
		t.nodes = t.nodes[:len+size]
	} else {
		nodes := make([]T, len+size, near(uint32(len+size)))
		copy(nodes, t.nodes)

		t.nodes = nodes
	}
}

func (t *BasicTree[T]) Read(marker TreeMarker) []T {
	return t.nodes[marker.pos : marker.pos+marker.len]
}

func (t *BasicTree[T]) Len() int {
	return len(t.nodes)
}

func (t *BasicTree[T]) Nodes() []T {
	return t.nodes
}

func (t *BasicTree[T]) Push(marker TreeMarker, pos *int, items ...T) {
	copied := copy(t.Read(marker)[*pos:], items)
	*pos += copied

	if copied < len(items) {
		panic("out of bounds")
	}
}

func (t *BasicTree[T]) Reserve(len int) TreeMarker {
	marker := TreeMarker{
		pos: t.pos,
		len: len,
	}
	t.pos += len

	return marker
}

func (b *Node[T, U]) Build(data T) (U, Tree[U]) {
	tree := NewTree[U]()
	result := b.BuildTo(data, tree)
	return result, tree
}

func (b *Node[T, U]) BuildTo(data T, tree Tree[U]) U {
	tree.Grow(b.noOfLeft + b.noOfRight)

	b.leftPos = 0
	b.rightPos = b.noOfLeft

	return b.build(data, tree, tree.Reserve(b.noOfLeft+b.noOfRight), &b.leftPos, &b.rightPos)
}

func NewDagger[T, U any](build func(T, Tree[U]) U) *Dagger[T, U] {
	return &Dagger[T, U]{
		build: build,
	}
}

func (b *Dagger[T, U]) Build(data T) (U, Tree[U]) {
	tree := NewTree[U]()
	result := b.BuildTo(data, tree)
	return result, tree
}

func (b *Dagger[T, U]) BuildTo(data T, tree Tree[U]) U {
	return b.build(data, tree)
}

// Append adds one or more children to the existing tree.
func (b *Node[T, U]) Append(nodes ...Buildable[T, U]) *Node[T, U] {
	return b.fmap(func(x T, tree Tree[U], marker TreeMarker, leftPos *int) {
		for _, node := range nodes {
			tree.Push(marker, leftPos, node.BuildTo(x, tree))
		}
	}, len(nodes))
}

// Annotate adds one of more annotations to the current node.
func (b *Node[T, U]) Annotate(nodes ...Buildable[T, U]) *Node[T, U] {
	return b.fmapRight(func(x T, tree Tree[U], marker TreeMarker, rightPos *int) {
		for _, node := range nodes {
			tree.Push(marker, rightPos, node.BuildTo(x, tree))
		}
	}, len(nodes))
}

// Lift allows for the dynamic insertion nodes into a tree.
func (b *Node[T, U]) Lift(f func(context T) Buildable[T, U]) *Node[T, U] {
	return b.fmap(func(x T, tree Tree[U], marker TreeMarker, leftPos *int) {
		tree.Push(marker, leftPos, f(x).BuildTo(x, tree))
	}, 1)
}

func (b *Node[T, U]) LiftRight(f func(context T) *Node[T, U]) *Node[T, U] {
	return b.fmapRight(func(x T, tree Tree[U], marker TreeMarker, rightPos *int) {
		tree.Push(marker, rightPos, f(x).BuildTo(x, tree))
	}, 1)
}

// fmap is a utility method that abstracts common behaviour required by
// [Node.Append], [Node.Annotate], and [Node.Lift].
func (b *Node[T, U]) fmap(f func(x T, tree Tree[U], marker TreeMarker, leftPos *int), noAdded int) *Node[T, U] {
	return &Node[T, U]{
		build: func(x T, tree Tree[U], marker TreeMarker, leftPos, rightPos *int) U {
			f(x, tree, marker, leftPos)
			return b.build(x, tree, marker, leftPos, rightPos)
		},

		noOfLeft:  b.noOfLeft + noAdded,
		noOfRight: b.noOfRight,
	}
}

// fmap is a utility method that abstracts common behaviour required by
// [Node.Append], [Node.Annotate], and [Node.Lift].
func (b *Node[T, U]) fmapRight(f func(x T, tree Tree[U], marker TreeMarker, rightPos *int), noAdded int) *Node[T, U] {
	return &Node[T, U]{
		build: func(x T, tree Tree[U], marker TreeMarker, leftPos, rightPos *int) U {
			f(x, tree, marker, rightPos)
			return b.build(x, tree, marker, leftPos, rightPos)
		},

		noOfLeft:  b.noOfLeft,
		noOfRight: b.noOfRight + noAdded,
	}
}

// Local maps the execution context from a parent tree, to be accepted by a
// subtree.
func Local[T1, T2, U any](b *Node[T2, U], f func(T1) T2) *Node[T1, U] {
	return &Node[T1, U]{
		build: func(x T1, tree Tree[U], marker TreeMarker, leftPos, rightPos *int) U {
			return b.build(f(x), tree, marker, leftPos, rightPos)
		},

		noOfLeft:  b.noOfLeft,
		noOfRight: b.noOfRight,
	}
}
