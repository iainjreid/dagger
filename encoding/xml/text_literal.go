package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLTextLiteralNode[T any] struct {
	f func(T) []byte
}

func NewTextLiteral[T any](f func(T) []byte) XMLTextLiteralNode[T] {
	return XMLTextLiteralNode[T]{
		f,
	}
}

func TextLiteral[T any](f func(T) []byte) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf[any, XMLNode[T]](func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewTextLiteral[T](f)
	})
}

func (x XMLTextLiteralNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x XMLTextLiteralNode[T]) Write(w Writer, data T) {
	w.Write(x.f(data))
}
