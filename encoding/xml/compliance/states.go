// Package compliance implements
package compliance

import (
	"encoding/xml"
	"io"
	"os"
	"strconv"

	"github.com/iainjreid/dagger"
	pkg "github.com/iainjreid/dagger/encoding/xml"
)

type XMLStates struct {
	XMLName xml.Name `xml:"states"`

	States []*XMLState `xml:"state"`
}

type States struct {
	States []*State
}

type XMLState struct {
	XMLName xml.Name `xml:"state"`

	Id   int    `xml:"id,attr"`
	Type string `xml:"type,attr"`
	Name string `xml:"name"`
	Code string `xml:"code"`

	Country *XMLCountry `xml:"country"`
	Coords  *XMLCoords  `xml:"coords"`
}

type State struct {
	Id   int
	Type string
	Name string
	Code string

	Country *Country
	Coords  *Coords
}

type XMLCountry struct {
	XMLName xml.Name `xml:"country"`

	Id   int    `xml:"id,attr"`
	Name string `xml:"name"`
	Code string `xml:"code"`
}

type Country struct {
	Id   int
	Name string
	Code string
}

type XMLCoords struct {
	XMLName xml.Name `xml:"coords"`

	Latitude  float64 `xml:"latitude"`
	Longitude float64 `xml:"longitude"`
}

type Coords struct {
	Latitude  float64
	Longitude float64
}

var StdStates XMLStates
var PkgStates States

var StateBytes []byte

func LoadData(path string) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	StateBytes, _ = io.ReadAll(file)

	xml.Unmarshal(StateBytes, &StdStates)

	for _, state := range StdStates.States {
		PkgStates.States = append(PkgStates.States, &State{
			Id:   state.Id,
			Type: state.Type,
			Name: state.Name,
			Code: state.Code,

			Country: &Country{
				Id:   state.Country.Id,
				Name: state.Country.Name,
				Code: state.Country.Code,
			},
			Coords: &Coords{
				Latitude:  state.Coords.Latitude,
				Longitude: state.Coords.Longitude,
			},
		})
	}

	file.Close()
}

func BuildXMLStates() pkg.XMLNode[*States] {
	return pkg.NewElement[*States]("states", []pkg.XMLNode[*States]{
		pkg.NewScope[*States, []*State](pkg.NewIterator[*State](BuildXMLState()), func(states *States) []*State {
			return states.States
		}),
	}, nil)
}

var StatesBuilder = pkg.Element[*States]("states").Append(
	pkg.Scope(pkg.Iterator[*State](stateBuilder), func(states *States) []*State {
		return states.States
	}),
)

var stateBuilder = pkg.Fragment(func(xml *pkg.XML[*State]) *dagger.Node[any, pkg.XMLNode[*State]] {
	return xml.Element("state").Annotate(
		xml.Attr("id", func(state *State) string { return strconv.Itoa(state.Id) }),
		xml.Attr("type", func(state *State) string { return state.Type }),
	).Append(
		xml.Element("name").Append(
			xml.Text(func(state *State) string { return state.Name }),
		),
		xml.Element("code").Append(
			xml.Text(func(state *State) string { return state.Code }),
		),
		xml.Element("country").Annotate(
			xml.Attr("id", func(state *State) string { return strconv.Itoa(state.Country.Id) }),
		).Append(
			xml.Element("name").Append(
				xml.Text(func(state *State) string { return state.Country.Name }),
			),
			xml.Element("code").Append(
				xml.Text(func(state *State) string { return state.Country.Code }),
			),
		),
		xml.Element("coords").Append(
			xml.Element("latitude").Append(
				xml.TextLiteral(func(state *State) []byte { return strconv.AppendFloat(nil, state.Coords.Latitude, 'f', 2, 64) }),
			),
			xml.Element("longitude").Append(
				xml.TextLiteral(func(state *State) []byte { return strconv.AppendFloat(nil, state.Coords.Longitude, 'f', 2, 64) }),
			),
		),
	)
})

func BuildCompXMLStates(states States) pkg.XMLNode[*States] {
	node, _ := StatesBuilder.Build(&states)
	return node
}

func BuildXMLState() pkg.XMLNode[*State] {
	return pkg.NewElement("state", []pkg.XMLNode[*State]{
		pkg.NewElement("name", []pkg.XMLNode[*State]{
			pkg.NewText[*State](func(state *State) string { return state.Name }),
		}, nil),
		pkg.NewElement("code", []pkg.XMLNode[*State]{
			pkg.NewText[*State](func(state *State) string { return state.Code }),
		}, nil),
		pkg.NewElement("country", []pkg.XMLNode[*State]{
			pkg.NewElement("name", []pkg.XMLNode[*State]{
				pkg.NewText[*State](func(state *State) string { return state.Country.Name }),
			}, nil),
			pkg.NewElement("code", []pkg.XMLNode[*State]{
				pkg.NewText[*State](func(state *State) string { return state.Country.Code }),
			}, nil),
		}, []pkg.XMLNode[*State]{
			pkg.NewAttr[*State]("id", func(state *State) string { return strconv.Itoa(state.Country.Id) }),
		}),
		pkg.NewElement("coords", []pkg.XMLNode[*State]{
			pkg.NewElement("latitude", []pkg.XMLNode[*State]{
				pkg.NewTextLiteral[*State](func(state *State) []byte { return strconv.AppendFloat(nil, state.Coords.Latitude, 'f', 2, 64) }),
			}, nil),
			pkg.NewElement("longitude", []pkg.XMLNode[*State]{
				pkg.NewTextLiteral[*State](func(state *State) []byte { return strconv.AppendFloat(nil, state.Coords.Longitude, 'f', 2, 64) }),
			}, nil),
		}, nil),
	}, []pkg.XMLNode[*State]{
		pkg.NewAttr[*State]("id", func(state *State) string { return strconv.Itoa(state.Id) }),
		pkg.NewAttr[*State]("type", func(state *State) string { return state.Type }),
	})
}
