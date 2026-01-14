package xml_test

import (
	"encoding/xml"
	"log"
	"strings"
	"testing"

	pkg "github.com/iainjreid/dagger/encoding/xml"
	"github.com/iainjreid/dagger/encoding/xml/compliance"
)

func TestMain(m *testing.M) {
	compliance.LoadData("./compliance/sample.xml")
	m.Run()
}

func TestPkgXMLMarshalling(t *testing.T) {
	states := compliance.BuildXMLStates()
	bytes, _ := pkg.Marshal(states, &compliance.PkgStates)

	if strings.TrimSpace(string(compliance.StateBytes)) != string(bytes) {
		log.Println(string(bytes))
		t.Fatal("XML encoding result did not match identity")
	}
}

func TestPkgCompXMLMarshalling(t *testing.T) {
	states := compliance.BuildCompXMLStates(compliance.PkgStates)
	bytes, _ := pkg.Marshal(states, &compliance.PkgStates)

	if strings.TrimSpace(string(compliance.StateBytes)) != string(bytes) {
		log.Print(string(bytes))
		t.Fatal("XML encoding result did not match identity")
	}
}

func BenchmarkStdXMLMarshalling(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)

	for b.Loop() {
		xml.Marshal(compliance.StdStates)
	}
}

func BenchmarkPkgPreXMLMarshalling(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildXMLStates()

	for b.Loop() {
		pkg.Marshal(states, &compliance.PkgStates)
	}
}

func BenchmarkPkgPreXMLMarshallingUnsafe(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildXMLStates()

	for b.Loop() {
		pkg.MarshalUnsafe(states, &compliance.PkgStates)
	}
}

func BenchmarkPkgPreXMLStreaming(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildXMLStates()

	for b.Loop() {
		states.Write(&strings.Builder{}, &compliance.PkgStates)
	}
}

func BenchmarkPkgCompXMLMarshalling(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildCompXMLStates(compliance.PkgStates)

	for b.Loop() {
		pkg.Marshal(states, &compliance.PkgStates)
	}
}

func BenchmarkPkgCompXMLMarshallingUnsafe(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildCompXMLStates(compliance.PkgStates)

	for b.Loop() {
		pkg.MarshalUnsafe(states, &compliance.PkgStates)
	}
}

func BenchmarkPkgCompXMLStreaming(b *testing.B) {
	b.SetBytes(compliance.StateBytesLen)
	states := compliance.BuildCompXMLStates(compliance.PkgStates)

	for b.Loop() {
		states.Write(&strings.Builder{}, &compliance.PkgStates)
	}
}
