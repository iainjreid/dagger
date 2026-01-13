package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLAttrLiteralNode[T any] struct {
	prefix string
	f      func(T) []byte
}

func NewAttrLiteral[T any](name string, f func(T) []byte) XMLAttrLiteralNode[T] {
	return XMLAttrLiteralNode[T]{
		prefix: " " + name + "=\"",
		f:      f,
	}
}

func AttrLiteral[T any](name string, fn func(T) []byte) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf(func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewAttrLiteral[T](name, fn)
	})
}

func (x XMLAttrLiteralNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x XMLAttrLiteralNode[T]) Write(w Writer, data T) {
	w.WriteString(x.prefix)
	w.Write(x.f(data))
	w.Write([]byte{'"'})
}
