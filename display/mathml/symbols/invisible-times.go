package symbols

import (
	"github.com/iainjreid/dagger/encoding/xml"
)

// InvisibleTimes is a special MathML character entity that is used to indicate
// to the renderer that spacing rules should be applied between the neighbouring
// characters. [Learn more].
//
// [Learn more]: https://www.w3.org/TR/MathML2/chapter2.html#fund.pres
var InvisibleTimes = xml.TextIdentity[any]("&InvisibleTimes;")
