package symbols

import "github.com/iainjreid/dagger/encoding/xml"

// PlusMinus is a special MathML character entity that is equivilant to using
// the unicode character "&#00B1", otherwise known as the [plus-minus sign].
//
// [plus-minus sign]: https://www.compart.com/en/unicode/U+00B1
var PlusMinus = xml.TextIdentity[any]("&PlusMinus;")
