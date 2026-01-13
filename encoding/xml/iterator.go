package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLIterator[T any] struct {
	node XMLNode[T]
}

func NewIterator[T any](node XMLNode[T]) XMLIterator[T] {
	return XMLIterator[T]{
		node: node,
	}
}

func Iterator[T any](node *dagger.Node[any, XMLNode[T]]) *dagger.Leaf[any, XMLNode[[]T]] {
	return dagger.NewLeaf[any, XMLNode[[]T]](func(data any, tree dagger.Tree[XMLNode[[]T]]) XMLNode[[]T] {
		_node, _ := node.Build(data)
		return NewIterator[T](_node)
	})
}

func (x XMLIterator[T]) WriteTo(w Writer, data []T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x XMLIterator[T]) Write(w Writer, data []T) {
	for i := 0; i < len(data); i++ {
		x.node.Write(w, data[i])
	}
}

type Scoped[T1, T2 any] struct {
	b XMLNode[T2]
	f func(T1) T2
}

func NewScope[T1, T2 any](b XMLNode[T2], f func(T1) T2) XMLNode[T1] {
	return &Scoped[T1, T2]{
		b: b,
		f: f,
	}
}

// Scope maps the execution context from a parent tree, to be accepted by a
// subtree.
func Scope[T1, T2 any](node dagger.Buildable[any, XMLNode[T2]], f func(T1) T2) *dagger.Leaf[any, XMLNode[T1]] {
	return dagger.NewLeaf[any, XMLNode[T1]](func(x any, tree dagger.Tree[XMLNode[T1]]) XMLNode[T1] {
		_node, _ := node.Build(x)
		return NewScope(_node, f)
	})
}

func (x *Scoped[T1, T2]) WriteTo(w Writer, data T1) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x *Scoped[T1, T2]) Write(w Writer, data T1) {
	x.b.Write(w, x.f(data))
}
