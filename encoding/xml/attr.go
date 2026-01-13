package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLAttrNode[T any] struct {
	prefix string
	f      func(T) string
}

func NewAttr[T any](name string, f func(T) string) XMLAttrNode[T] {
	return XMLAttrNode[T]{
		prefix: " " + name + "=\"",
		f:      f,
	}
}

func Attr[T any](name string, fn func(T) string) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf(func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewAttr(name, fn)
	})
}

func (x XMLAttrNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

var b = []byte{'"'}

func (x XMLAttrNode[T]) Write(w Writer, data T) {
	w.WriteString(x.prefix)
	w.WriteString(x.f(data))
	w.Write(b)
}
