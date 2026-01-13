package html

import (
	"github.com/iainjreid/dagger"
	"github.com/iainjreid/dagger/encoding/xml"
)

type Attr[T any] struct{}

func (s *Attr[T]) Id(name string) *dagger.Leaf[any, xml.XMLNode[T]] {
	return xml.AttrIdentity[T]("id", name)
}

func (s *Attr[T]) Class(...name []string) *dagger.Leaf[any, xml.XMLNode[T]] {
	return xml.AttrIdentity[T]("class", name)
}
