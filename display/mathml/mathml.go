// Package mathml aims to provide an exhaustive set of MathML elements, as
// described by the [MathML working draft].
//
// Links to the specification are available throughout the documentation, with
// anchors to the closest section where possible.
//
// [MathML working draft]: https://www.w3.org/TR/2023/WD-mathml-core-20231127
package mathml

import (
	"github.com/iainjreid/dagger/encoding/xml"
)

// AnnotationXml is an element that contains an annotation for the associated
// MathML expression in XML format. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/annotation-xml
var AnnotationXml = xml.Element[any]("annotation-xml")

// Annotation is an element that contains an annotation for the associated
// MathML expression in a textual format. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/annotation
var Annotation = xml.Element[any]("annotation")

// Maction is an element that allows for certain user interactions with the
// associated MathML expression without the use of JavaScript. [Learn more].
//
// Deprecated: Maction should no longer be used (although some browsers may
// still support it). It exists for completeness, and historical compatibility.
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/annotation-xml
var Maction = xml.Element[any]("maction")

// Math is the top-level MathML element. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/math
var Math = xml.Element[any]("math")

// Menclose is an element that renders its child elements within or alongside a
// complementary geometric form. This could be a circle, square, line, etc. or
// any chosen combination of the supported forms. [Learn more].
//
// Deprecated: Menclose is a non-standard feature, and does not work on all
// browsers.
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/menclose
var Menclose = xml.Element[any]("menclose")

// Merror is an element that can be used to display an error message within a
// [Math] block. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/merror
var Merror = xml.Element[any]("merror")

// Mfenced is an element that overrides the opening, closing, and delimiting
// characters used to display a MathML expression. [Learn more].
//
// Deprecated: Mfenced should no longer be used (although some browsers may
// still support it). It exists for completeness, and historical compatibility.
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mfenced
var Mfenced = xml.Element[any]("mfenced")

// Mfrac is an element that can be used to display fractions and fraction-like
// objects (such as binomial coefficients and Legendre symbols). [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mfrac
var Mfrac = xml.Element[any]("mfrac")

// Mi is an element that is used to help identify its contents as a mathematical
// function, variable, or sybolic constant. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mi
var Mi = xml.Element[any]("mi")

// Mmultiscripts is an element used to annotate an expression with an arbitrary
// number of subscripts and superscripts. The expression to be annotated is
// assigned as its first child. [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mmultiscripts
var Mmultiscripts = xml.Element[any]("mmultiscripts")

// Mn is an element used to display a numeric literal. This could be an integer,
// a decimal, a number written scientific notation, or even a number expressed
// in plain text such as "thirty-six". [Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mn
var Mn = xml.Element[any]("mn")

// Mo is an element used to display mathematical operators and the special
// character entities supported by MathML. [Learn more].
//
// The [symbols] package aims to provide an exhastive list of special character
// entities, and explains why they should be preferred over the traditional
// operators.
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mo
var Mo = xml.Element[any]("mo")

// Mover is an element used to display an overscript above an expression.
//
// References:
//   - [W3C]
//   - [MDN]
//
// [W3C]: https://www.w3.org/TR/mathml-core/#dfn-mover
// [MDN]: https://developer.mozilla.org/en-US/docs/Web/MathML/Element/mover
var Mover = xml.Element[any]("mover")

// Mpadded is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mpadded
var Mpadded = xml.Element[any]("mpadded")

// Mphantom is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mphantom
var Mphantom = xml.Element[any]("mphantom")

// Mprescripts is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mprescripts
var Mprescripts = xml.Element[any]("mprescripts")

// Mroot is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mroot
var Mroot = xml.Element[any]("mroot")

// Mrow is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mrow
var Mrow = xml.Element[any]("mrow")

// Ms is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/ms
var Ms = xml.Element[any]("ms")

// Mspace is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mspace
var Mspace = xml.Element[any]("mspace")

// Msqrt is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/msqrt
var Msqrt = xml.Element[any]("msqrt")

// Mstyle is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mstyle
var Mstyle = xml.Element[any]("mstyle")

// Msub is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/msub
var Msub = xml.Element[any]("msub")

// Msubsup is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/msubsup
var Msubsup = xml.Element[any]("msubsup")

// Msup is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/msup
var Msup = xml.Element[any]("msup")

// Mtable is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mtable
var Mtable = xml.Element[any]("mtable")

// Mtd is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mtd
var Mtd = xml.Element[any]("mtd")

// Mtext is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mtext
var Mtext = xml.Element[any]("mtext")

// Mtr is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/mtr
var Mtr = xml.Element[any]("mtr")

// Munder is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/munder
var Munder = xml.Element[any]("munder")

// Munderover is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/munderover
var Munderover = xml.Element[any]("munderover")

// Semantics is[Learn more].
//
// [Learn more]: https://developer.mozilla.org/docs/Web/MathML/Element/semantics
var Semantics = xml.Element[any]("semantics")
