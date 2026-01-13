package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLTextNode[T any] struct {
	f func(T) string
}

func NewText[T any](f func(T) string) XMLNode[T] {
	return XMLTextNode[T]{
		f,
	}
}

func Text[T any](f func(T) string) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf[any, XMLNode[T]](func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewText[T](f)
	})
}

func (x XMLTextNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x XMLTextNode[T]) Write(w Writer, data T) {
	w.WriteString(x.f(data))
}
