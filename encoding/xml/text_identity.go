package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLTextIdentityNode[T any] struct {
	b string
}

// NewTextIdentity creates a new XMLTextNode[T] struct.
func NewTextIdentity[T any](text string) XMLNode[T] {
	return XMLTextIdentityNode[T]{
		b: text,
	}
}

func TextIdentity[T any](text string) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf[any, XMLNode[T]](func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewTextIdentity[T](text)
	})
}

func (x XMLTextIdentityNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

func (x XMLTextIdentityNode[T]) Write(w Writer, _ T) {
	w.WriteString(x.b)
}
