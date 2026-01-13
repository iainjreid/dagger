package xml

import (
	"github.com/iainjreid/dagger"
)

// TODO: Should we provide a comment identity, and a comment FUNC?
type XMLCommentNode[T any] struct {
	comment []byte
}

func Comment[T any](comment string) *dagger.Leaf[any, XMLNode[T]] {
	return dagger.NewLeaf[any, XMLNode[T]](func(any, dagger.Tree[XMLNode[T]]) XMLNode[T] {
		return &XMLCommentNode[T]{
			comment: []byte(comment),
		}
	})
}

func (x *XMLCommentNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

var open = []byte("<!-- ")
var close = []byte(" -->")

func (x *XMLCommentNode[T]) Write(w Writer, data T) {
	w.Write(open)
	w.Write(x.comment)
	w.Write(close)
}
