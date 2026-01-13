package xml

import (
	"github.com/iainjreid/dagger"
)

// Attr ID is just a literal string, so its safe to reuse the text ID.
func NewAttrIdentity[T any](name string, value string) XMLNode[T] {
	return NewTextIdentity[T](" " + name + "=\"" + value + "\"")
}

func AttrIdentity[T any](name string, value string) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf(func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return NewAttrIdentity[T](name, value)
	})
}
