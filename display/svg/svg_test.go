package svg_test

import (
	"bytes"
	stdxml "encoding/xml"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/iainjreid/dagger"
	"github.com/iainjreid/dagger/display/svg"
	"github.com/iainjreid/dagger/encoding/xml"
)

var node = svg.Fragment(func(svg *svg.SVG[any], xml *xml.XML[any]) *dagger.Node[any, xml.XMLNode[any]] {
	return svg.Svg(
		xml.AttrIdentity("xmlns", "http://www.w3.org/2000/svg"),
		xml.AttrIdentity("width", "1000"),
		xml.AttrIdentity("height", "600"),
		xml.AttrIdentity("viewBox", "0 0 5 5"),
	).Append(
		svg.Rect(
			xml.AttrIdentity("id", "shade_1"),
			xml.AttrIdentity("fill", "#000"),
			xml.AttrIdentity("width", "5"),
			xml.AttrIdentity("height", "1"),
		),
		svg.Rect(
			xml.AttrIdentity("id", "shade_2"),
			xml.AttrIdentity("fill", "#444"),
			xml.AttrIdentity("width", "5"),
			xml.AttrIdentity("height", "1"),
			xml.AttrIdentity("y", "1"),
		),
		svg.Rect(
			xml.AttrIdentity("id", "shade_3"),
			xml.AttrIdentity("fill", "#888"),
			xml.AttrIdentity("width", "5"),
			xml.AttrIdentity("height", "1"),
			xml.AttrIdentity("y", "2"),
		),
		svg.Rect(
			xml.AttrIdentity("id", "shade_4"),
			xml.AttrIdentity("fill", "#ccc"),
			xml.AttrIdentity("width", "5"),
			xml.AttrIdentity("height", "1"),
			xml.AttrIdentity("y", "3"),
		),
		svg.Rect(
			xml.AttrIdentity("id", "shade_5"),
			xml.AttrIdentity("fill", "#fff"),
			xml.AttrIdentity("width", "5"),
			xml.AttrIdentity("height", "1"),
			xml.AttrIdentity("y", "4"),
		),
	)
})

var target, _ = node.Build(nil)

func TestSVGMarshal(t *testing.T) {
	file, _ := os.Open("sample.svg")
	svgBytes, _ := io.ReadAll(file)
	file.Close()

	target, _ := node.Build(nil)
	xmlBytes, _ := xml.Marshal(target, nil)

	if bytes.Compare(svgBytes, xmlBytes) != 0 {
		t.Errorf("got=%s, want=%s", string(xmlBytes), string(svgBytes))
	}
}

func BenchmarkSVGMarshal(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xml.Marshal(target, nil)
	}
}

func BenchmarkSVGMarshalUnsafe(b *testing.B) {
	for n := 0; n < b.N; n++ {
		xml.MarshalUnsafe(target, nil)
	}
}

func BenchmarkSVGWriteTo(b *testing.B) {
	for n := 0; n < b.N; n++ {
		target.Write(&strings.Builder{}, nil)
	}
}

type SVG struct {
	XMLName stdxml.Name `xml:"svg"`
	Rects   []Rect      `xml:"rect"`
}

type Rect struct {
	XMLName stdxml.Name `xml:"rect"`
	Id      string      `xml:"id,attr"`
	Fill    string      `xml:"fill,attr"`
}

func BenchmarkStdSVGMarshal(b *testing.B) {
	file, _ := os.Open("sample.svg")
	bytes, _ := io.ReadAll(file)
	var wild SVG
	stdxml.Unmarshal(bytes, &wild)
	file.Close()
	b.ResetTimer()

	for n := 0; n < b.N; n++ {
		stdxml.Marshal(wild)
	}
}
