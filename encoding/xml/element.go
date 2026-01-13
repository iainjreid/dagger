package xml

import (
	"github.com/iainjreid/dagger"
)

type XMLElementNode[T any] struct {
	prefix        []byte
	postfix       []byte
	children      []XMLNode[T]
	childrenLen   int
	attributes    []XMLNode[T]
	attributesLen int
}

// NewElement creates a new *XMLElementNode struct.
func NewElement[T any](name string, children []XMLNode[T], attributes []XMLNode[T]) *XMLElementNode[T] {
	return &XMLElementNode[T]{
		prefix:        []byte("<" + name),
		postfix:       []byte("</" + name + ">"),
		children:      children,
		childrenLen:   len(children),
		attributes:    attributes,
		attributesLen: len(attributes),
	}
}

func Element[T any](name string) *dagger.Node[any, XMLNode[T]] {
	return dagger.NewNode(func(_ any, ln []XMLNode[T], rn []XMLNode[T]) XMLNode[T] {
		return NewElement(name, ln, rn)
	})
}

func (x *XMLElementNode[T]) WriteTo(w Writer, data T) (int64, error) {
	cw := NewCollectedWriter(w)
	x.Write(cw, data)
	return cw.Result()
}

var gt = []byte{'>'}
var slash = []byte{'/'}
var slashgt = []byte{'/', '>'}

func (x *XMLElementNode[T]) Write(w Writer, data T) {
	w.Write(x.prefix)

	if x.attributesLen > 0 {
		for _, attribute := range x.attributes {
			attribute.Write(w, data)
		}
	}

	if x.childrenLen > 0 {
		w.Write(gt)

		for _, child := range x.children {
			child.Write(w, data)
		}

		w.Write(x.postfix)
	} else {
		w.Write(slashgt)
	}
}
