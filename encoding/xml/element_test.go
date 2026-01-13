package xml_test

import (
	stdxml "encoding/xml"
	"testing"

	"github.com/iainjreid/dagger"
	"github.com/iainjreid/dagger/encoding/xml"
)

type Catalog struct {
	XMLName stdxml.Name `xml:"catalog"`
	Plants  []Plant     `xml:"plant"`
}

type Plant struct {
	XMLName      stdxml.Name `xml:"PLANT"`
	Common       string      `xml:"COMMON"`
	Botanical    string      `xml:"BOTANICAL"`
	Zone         string      `xml:"ZONE"`
	Light        string      `xml:"LIGHT"`
	Price        string      `xml:"PRICE"`
	Availability string      `xml:"AVAILABILITY"`
}

var catalog Catalog

// func TestMain(m *testing.M) {
// 	file, _ := os.Open("catalog.xml")
// 	bytes, _ := io.ReadAll(file)

// 	defer file.Close()

// 	stdxml.Unmarshal(bytes, &catalog)

// 	m.Run()
// }

func reduce[T, U any](slice []T, f func(U, T) U, initial U) U {
	acc := initial
	for _, v := range slice {
		acc = f(acc, v)
	}
	return acc
}

// var target = reduce(catalog.Plants, func(acc *fst.NodeBuilder[any, xml.XMLNode, xml.XMLAttr], plant Plant) *fst.Builder[any, xml.XMLNode, xml.XMLAttr] {
// 	return acc.Append(
// 		xml.Element[any]("COMMON").Append(xml.TextIdentity[any](plant.Common)),
// 		xml.Element[any]("BOTANICAL").Append(xml.TextIdentity[any](plant.Botanical)),
// 		xml.Element[any]("ZONE").Append(xml.TextIdentity[any](plant.Zone)),
// 		xml.Element[any]("LIGHT").Append(xml.TextIdentity[any](plant.Light)),
// 		xml.Element[any]("PRICE").Append(xml.TextIdentity[any](plant.Price)),
// 		xml.Element[any]("AVAILABILITY").Append(xml.TextIdentity[any](plant.Availability)),
// 	)
// }, xml.Element[any]("CATALOG")).Build(nil)

func TestStruct(t *testing.T) {
	var subject, _ = xml.Fragment(func(xml *xml.XML[any]) *dagger.Node[any, xml.XMLNode[any]] {
		return xml.Element("parent").Append(
			xml.Element("child").Append(
				xml.TextIdentity("Text"),
			),
		)
	}).Build(nil)

	var expected = "<parent><child>Text</child></parent>"

	if string(xml.MarshalUnsafe(subject, nil)) != expected {
		t.Fatal("encoded tree should match expected string")
	}
}

// func BenchmarkXML(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		xml.Marshal(target)
// 	}
// }

// func BenchmarkXMLUnsafe(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		xml.MarshalUnsafe(target)
// 	}
// }

// func BenchmarkXMLWriteTo(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		target.WriteTo(io.Discard)
// 	}
// }

// func BenchmarkStdXML(b *testing.B) {
// 	for n := 0; n < b.N; n++ {
// 		stdxml.Marshal(catalog)
// 	}
// }
