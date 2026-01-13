// Package svg aims to provide an axhaustive set of SVG elements, as described
// by the [SVG candidate Recomendation]
//
// [SVG Candidate Recomendation]: https://www.w3.org/TR/2011/REC-SVG11-20110816/
package svg

import (
	"github.com/iainjreid/dagger"
	"github.com/iainjreid/dagger/encoding/xml"
)

type SVG[T any] struct{}

func Fragment[T any](f func(*SVG[T], *xml.XML[T]) *dagger.Node[T, xml.XMLNode[T]]) *dagger.Node[T, xml.XMLNode[T]] {
	return f(&SVG[T]{}, &xml.XML[T]{})
}

func (s *SVG[T]) A(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("a").Annotate(attrs...)
	} else {
		return xml.Element[T]("a")
	}
}

func (s *SVG[T]) Animate(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("animate").Annotate(attrs...)
	} else {
		return xml.Element[T]("animate")
	}
}

func (s *SVG[T]) AnimateMotion(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("animateMotion").Annotate(attrs...)
	} else {
		return xml.Element[T]("animateMotion")
	}
}

func (s *SVG[T]) AnimateTransform(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("animateTransform").Annotate(attrs...)
	} else {
		return xml.Element[T]("animateTransform")
	}
}

func (s *SVG[T]) Circle(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("circle").Annotate(attrs...)
	} else {
		return xml.Element[T]("circle")
	}
}

func (s *SVG[T]) ClipPath(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("clipPath").Annotate(attrs...)
	} else {
		return xml.Element[T]("clipPath")
	}
}

func (s *SVG[T]) Cursor(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("cursor").Annotate(attrs...)
	} else {
		return xml.Element[T]("cursor")
	}
}

func (s *SVG[T]) Defs(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("defs").Annotate(attrs...)
	} else {
		return xml.Element[T]("defs")
	}
}

func (s *SVG[T]) Desc(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("desc").Annotate(attrs...)
	} else {
		return xml.Element[T]("desc")
	}
}

func (s *SVG[T]) Ellipse(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("ellipse").Annotate(attrs...)
	} else {
		return xml.Element[T]("ellipse")
	}
}

func (s *SVG[T]) FeBlend(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feBlend").Annotate(attrs...)
	} else {
		return xml.Element[T]("feBlend")
	}
}

func (s *SVG[T]) FeColorMatrix(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feColorMatrix").Annotate(attrs...)
	} else {
		return xml.Element[T]("feColorMatrix")
	}
}

func (s *SVG[T]) FeComponentTransfer(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feComponentTransfer").Annotate(attrs...)
	} else {
		return xml.Element[T]("feComponentTransfer")
	}
}

func (s *SVG[T]) FeComposite(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feComposite").Annotate(attrs...)
	} else {
		return xml.Element[T]("feComposite")
	}
}

func (s *SVG[T]) FeConvolveMatrix(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feConvolveMatrix").Annotate(attrs...)
	} else {
		return xml.Element[T]("feConvolveMatrix")
	}
}

func (s *SVG[T]) FeDiffuseLighting(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feDiffuseLighting").Annotate(attrs...)
	} else {
		return xml.Element[T]("feDiffuseLighting")
	}
}

func (s *SVG[T]) FeDisplacementMap(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feDisplacementMap").Annotate(attrs...)
	} else {
		return xml.Element[T]("feDisplacementMap")
	}
}

func (s *SVG[T]) FeDistantLight(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feDistantLight").Annotate(attrs...)
	} else {
		return xml.Element[T]("feDistantLight")
	}
}

func (s *SVG[T]) FeDropShadow(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feDropShadow").Annotate(attrs...)
	} else {
		return xml.Element[T]("feDropShadow")
	}
}

func (s *SVG[T]) FeFlood(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feFlood").Annotate(attrs...)
	} else {
		return xml.Element[T]("feFlood")
	}
}

func (s *SVG[T]) FeFuncA(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feFuncA").Annotate(attrs...)
	} else {
		return xml.Element[T]("feFuncA")
	}
}

func (s *SVG[T]) FeFuncB(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feFuncB").Annotate(attrs...)
	} else {
		return xml.Element[T]("feFuncB")
	}
}

func (s *SVG[T]) FeFuncG(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feFuncG").Annotate(attrs...)
	} else {
		return xml.Element[T]("feFuncG")
	}
}

func (s *SVG[T]) FeFuncR(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feFuncR").Annotate(attrs...)
	} else {
		return xml.Element[T]("feFuncR")
	}
}

func (s *SVG[T]) FeGaussianBlur(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feGaussianBlur").Annotate(attrs...)
	} else {
		return xml.Element[T]("feGaussianBlur")
	}
}

func (s *SVG[T]) FeImage(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feImage").Annotate(attrs...)
	} else {
		return xml.Element[T]("feImage")
	}
}

func (s *SVG[T]) FeMerge(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feMerge").Annotate(attrs...)
	} else {
		return xml.Element[T]("feMerge")
	}
}

func (s *SVG[T]) FeMergeNode(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feMergeNode").Annotate(attrs...)
	} else {
		return xml.Element[T]("feMergeNode")
	}
}

func (s *SVG[T]) FeMorphology(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feMorphology").Annotate(attrs...)
	} else {
		return xml.Element[T]("feMorphology")
	}
}

func (s *SVG[T]) FeOffset(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feOffset").Annotate(attrs...)
	} else {
		return xml.Element[T]("feOffset")
	}
}

func (s *SVG[T]) FePointLight(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("fePointLight").Annotate(attrs...)
	} else {
		return xml.Element[T]("fePointLight")
	}
}

func (s *SVG[T]) FeSpecularLighting(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feSpecularLighting").Annotate(attrs...)
	} else {
		return xml.Element[T]("feSpecularLighting")
	}
}

func (s *SVG[T]) FeSpotLight(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feSpotLight").Annotate(attrs...)
	} else {
		return xml.Element[T]("feSpotLight")
	}
}

func (s *SVG[T]) FeTile(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feTile").Annotate(attrs...)
	} else {
		return xml.Element[T]("feTile")
	}
}

func (s *SVG[T]) FeTurbulence(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("feTurbulence").Annotate(attrs...)
	} else {
		return xml.Element[T]("feTurbulence")
	}
}

func (s *SVG[T]) Filter(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("filter").Annotate(attrs...)
	} else {
		return xml.Element[T]("filter")
	}
}

func (s *SVG[T]) FontFaceFormat(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font-face-format").Annotate(attrs...)
	} else {
		return xml.Element[T]("font-face-format")
	}
}

func (s *SVG[T]) FontFaceName(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font-face-name").Annotate(attrs...)
	} else {
		return xml.Element[T]("font-face-name")
	}
}

func (s *SVG[T]) FontFaceSrc(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font-face-src").Annotate(attrs...)
	} else {
		return xml.Element[T]("font-face-src")
	}
}

func (s *SVG[T]) FontFaceUri(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font-face-uri").Annotate(attrs...)
	} else {
		return xml.Element[T]("font-face-uri")
	}
}

func (s *SVG[T]) FontFace(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font-face").Annotate(attrs...)
	} else {
		return xml.Element[T]("font-face")
	}
}

func (s *SVG[T]) Font(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("font").Annotate(attrs...)
	} else {
		return xml.Element[T]("font")
	}
}

func (s *SVG[T]) ForeignObject(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("foreignObject").Annotate(attrs...)
	} else {
		return xml.Element[T]("foreignObject")
	}
}

func (s *SVG[T]) G(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("g").Annotate(attrs...)
	} else {
		return xml.Element[T]("g")
	}
}

func (s *SVG[T]) Glyph(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("glyph").Annotate(attrs...)
	} else {
		return xml.Element[T]("glyph")
	}
}

func (s *SVG[T]) GlyphRef(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("glyphRef").Annotate(attrs...)
	} else {
		return xml.Element[T]("glyphRef")
	}
}

func (s *SVG[T]) Hkern(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("hkern").Annotate(attrs...)
	} else {
		return xml.Element[T]("hkern")
	}
}

func (s *SVG[T]) Image(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("image").Annotate(attrs...)
	} else {
		return xml.Element[T]("image")
	}
}

func (s *SVG[T]) Line(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("line").Annotate(attrs...)
	} else {
		return xml.Element[T]("line")
	}
}

func (s *SVG[T]) LinearGradient(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("linearGradient").Annotate(attrs...)
	} else {
		return xml.Element[T]("linearGradient")
	}
}

func (s *SVG[T]) Marker(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("marker").Annotate(attrs...)
	} else {
		return xml.Element[T]("marker")
	}
}

func (s *SVG[T]) Mask(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("mask").Annotate(attrs...)
	} else {
		return xml.Element[T]("mask")
	}
}

func (s *SVG[T]) Metadata(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("metadata").Annotate(attrs...)
	} else {
		return xml.Element[T]("metadata")
	}
}

func (s *SVG[T]) MissingGlyph(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("missing-glyph").Annotate(attrs...)
	} else {
		return xml.Element[T]("missing-glyph")
	}
}

func (s *SVG[T]) Mpath(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("mpath").Annotate(attrs...)
	} else {
		return xml.Element[T]("mpath")
	}
}

func (s *SVG[T]) Path(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("path").Annotate(attrs...)
	} else {
		return xml.Element[T]("path")
	}
}

func (s *SVG[T]) Pattern(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("pattern").Annotate(attrs...)
	} else {
		return xml.Element[T]("pattern")
	}
}

func (s *SVG[T]) Polygon(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("polygon").Annotate(attrs...)
	} else {
		return xml.Element[T]("polygon")
	}
}

func (s *SVG[T]) Polyline(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("polyline").Annotate(attrs...)
	} else {
		return xml.Element[T]("polyline")
	}
}

func (s *SVG[T]) RadialGradient(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("radialGradient").Annotate(attrs...)
	} else {
		return xml.Element[T]("radialGradient")
	}
}

func (s *SVG[T]) Rect(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("rect").Annotate(attrs...)
	} else {
		return xml.Element[T]("rect")
	}
}

func (s *SVG[T]) Script(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("script").Annotate(attrs...)
	} else {
		return xml.Element[T]("script")
	}
}

func (s *SVG[T]) Set(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("set").Annotate(attrs...)
	} else {
		return xml.Element[T]("set")
	}
}

func (s *SVG[T]) Stop(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("stop").Annotate(attrs...)
	} else {
		return xml.Element[T]("stop")
	}
}

func (s *SVG[T]) Style(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("style").Annotate(attrs...)
	} else {
		return xml.Element[T]("style")
	}
}

func (s *SVG[T]) Svg(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("svg").Annotate(attrs...)
	} else {
		return xml.Element[T]("svg")
	}
}

func (s *SVG[T]) Switch(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("switch").Annotate(attrs...)
	} else {
		return xml.Element[T]("switch")
	}
}

func (s *SVG[T]) Symbol(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("symbol").Annotate(attrs...)
	} else {
		return xml.Element[T]("symbol")
	}
}

func (s *SVG[T]) Text(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("text").Annotate(attrs...)
	} else {
		return xml.Element[T]("text")
	}
}

func (s *SVG[T]) TextPath(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("textPath").Annotate(attrs...)
	} else {
		return xml.Element[T]("textPath")
	}
}

func (s *SVG[T]) Title(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("title").Annotate(attrs...)
	} else {
		return xml.Element[T]("title")
	}
}

func (s *SVG[T]) Tref(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("tref").Annotate(attrs...)
	} else {
		return xml.Element[T]("tref")
	}
}

func (s *SVG[T]) Tspan(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("tspan").Annotate(attrs...)
	} else {
		return xml.Element[T]("tspan")
	}
}

func (s *SVG[T]) Use(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("use").Annotate(attrs...)
	} else {
		return xml.Element[T]("use")
	}
}

func (s *SVG[T]) View(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("view").Annotate(attrs...)
	} else {
		return xml.Element[T]("view")
	}
}

func (s *SVG[T]) Vkern(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("vkern").Annotate(attrs...)
	} else {
		return xml.Element[T]("vkern")
	}
}
