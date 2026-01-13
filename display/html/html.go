// Package html aims to provide an exhaustive set of HTML elements, as described
// by the [HTML living standard].
//
// [HTML living standard]: https://html.spec.whatwg.org/multipage
package html

import (
	"github.com/iainjreid/dagger"
	"github.com/iainjreid/dagger/encoding/xml"
)

type HTML[T any] struct{}

func Fragment[T any](f func(*HTML[T], *xml.XML[T]) *dagger.Node[T, xml.XMLNode[T]]) *dagger.Node[T, xml.XMLNode[T]] {
	return f(&HTML[T]{}, &xml.XML[T]{})
}

// Represents the root (top-level element) of an HTML document, so it is also referred to as the _root element_. All other elements must be descendants of this element.
func (s *HTML[T]) Html(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("html").Annotate(attrs...)
	} else {
		return xml.Element[T]("html")
	}
}

// Specifies the base URL to use for all relative URLs in a document. There can be only one such element in a document.
func (s *HTML[T]) Base(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("base").Annotate(attrs...)
	} else {
		return xml.Element[T]("base")
	}
}

// Contains machine-readable information (metadata) about the document, like its [title](/en-US/docs/Web/HTML/Element/title), [scripts](/en-US/docs/Web/HTML/Element/script), and [style sheets](/en-US/docs/Web/HTML/Element/style).
func (s *HTML[T]) Head(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("head").Annotate(attrs...)
	} else {
		return xml.Element[T]("head")
	}
}

// Specifies relationships between the current document and an external resource. This element is most commonly used to link to CSS but is also used to establish site icons (both "favicon" style icons and icons for the home screen and apps on mobile devices) among other things.
func (s *HTML[T]) Link(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("link").Annotate(attrs...)
	} else {
		return xml.Element[T]("link")
	}
}

// Represents {{Glossary("Metadata","metadata")}} that cannot be represented by other HTML meta-related elements, like {{HTMLElement("base")}}, {{HTMLElement("link")}}, {{HTMLElement("script")}}, {{HTMLElement("style")}} and {{HTMLElement("title")}}.
func (s *HTML[T]) Meta(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("meta").Annotate(attrs...)
	} else {
		return xml.Element[T]("meta")
	}
}

// Contains style information for a document or part of a document. It contains CSS, which is applied to the contents of the document containing this element.
func (s *HTML[T]) Style(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("style").Annotate(attrs...)
	} else {
		return xml.Element[T]("style")
	}
}

// Defines the document's title that is shown in a {{glossary("Browser", "browser")}}'s title bar or a page's tab. It only contains text; HTML tags within the element, if any, are also treated as plain text.
func (s *HTML[T]) Title(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("title").Annotate(attrs...)
	} else {
		return xml.Element[T]("title")
	}
}

// Represents the content of an HTML document. There can be only one such element in a document.
func (s *HTML[T]) Body(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("body").Annotate(attrs...)
	} else {
		return xml.Element[T]("body")
	}
}

// Indicates that the enclosed HTML provides contact information for a person or people, or for an organization.
func (s *HTML[T]) Address(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("address").Annotate(attrs...)
	} else {
		return xml.Element[T]("address")
	}
}

// Represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). Examples include a forum post, a magazine or newspaper article, a blog entry, a product card, a user-submitted comment, an interactive widget or gadget, or any other independent item of content.
func (s *HTML[T]) Article(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("article").Annotate(attrs...)
	} else {
		return xml.Element[T]("article")
	}
}

// Represents a portion of a document whose content is only indirectly related to the document's main content. Asides are frequently presented as sidebars or call-out boxes.
func (s *HTML[T]) Aside(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("aside").Annotate(attrs...)
	} else {
		return xml.Element[T]("aside")
	}
}

// Represents a footer for its nearest ancestor [sectioning content](/en-US/docs/Web/HTML/Content_categories#sectioning_content) or [sectioning root](/en-US/docs/Web/HTML/Element/Heading_Elements) element. A `<footer>` typically contains information about the author of the section, copyright data, or links to related documents.
func (s *HTML[T]) Footer(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("footer").Annotate(attrs...)
	} else {
		return xml.Element[T]("footer")
	}
}

// Represents introductory content, typically a group of introductory or navigational aids. It may contain some heading elements but also a logo, a search form, an author name, and other elements.
func (s *HTML[T]) Header(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("header").Annotate(attrs...)
	} else {
		return xml.Element[T]("header")
	}
}

// Represent six levels of section headings. `<h1>` is the highest section level and `<h6>` is the lowest.
func (s *HTML[T]) H1(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h1").Annotate(attrs...)
	} else {
		return xml.Element[T]("h1")
	}
}
func (s *HTML[T]) H2(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h2").Annotate(attrs...)
	} else {
		return xml.Element[T]("h2")
	}
}
func (s *HTML[T]) H3(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h3").Annotate(attrs...)
	} else {
		return xml.Element[T]("h3")
	}
}
func (s *HTML[T]) H4(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h4").Annotate(attrs...)
	} else {
		return xml.Element[T]("h4")
	}
}
func (s *HTML[T]) H5(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h5").Annotate(attrs...)
	} else {
		return xml.Element[T]("h5")
	}
}
func (s *HTML[T]) H6(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("h6").Annotate(attrs...)
	} else {
		return xml.Element[T]("h6")
	}
}

// Represents a heading grouped with any secondary content, such as subheadings, an alternative title, or a tagline.
func (s *HTML[T]) Hgroup(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("hgroup").Annotate(attrs...)
	} else {
		return xml.Element[T]("hgroup")
	}
}

// Represents the dominant content of the body of a document. The main content area consists of content that is directly related to or expands upon the central topic of a document, or the central functionality of an application.
func (s *HTML[T]) Main(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("main").Annotate(attrs...)
	} else {
		return xml.Element[T]("main")
	}
}

// Represents a section of a page whose purpose is to provide navigation links, either within the current document or to other documents. Common examples of navigation sections are menus, tables of contents, and indexes.
func (s *HTML[T]) Nav(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("nav").Annotate(attrs...)
	} else {
		return xml.Element[T]("nav")
	}
}

// Represents a generic standalone section of a document, which doesn't have a more specific semantic element to represent it. Sections should always have a heading, with very few exceptions.
func (s *HTML[T]) Section(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("section").Annotate(attrs...)
	} else {
		return xml.Element[T]("section")
	}
}

// Represents a part that contains a set of form controls or other content related to performing a search or filtering operation.
func (s *HTML[T]) Search(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("search").Annotate(attrs...)
	} else {
		return xml.Element[T]("search")
	}
}

// Indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation. A URL for the source of the quotation may be given using the `cite` attribute, while a text representation of the source can be given using the {{HTMLElement("cite")}} element.
func (s *HTML[T]) Blockquote(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("blockquote").Annotate(attrs...)
	} else {
		return xml.Element[T]("blockquote")
	}
}

// Provides the description, definition, or value for the preceding term ({{HTMLElement("dt")}}) in a description list ({{HTMLElement("dl")}}).
func (s *HTML[T]) Dd(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("dd").Annotate(attrs...)
	} else {
		return xml.Element[T]("dd")
	}
}

// The generic container for flow content. It has no effect on the content or layout until styled in some way using CSS (e.g., styling is directly applied to it, or some kind of layout model like {{glossary("Flexbox", "flexbox")}} is applied to its parent element).
func (s *HTML[T]) Div(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("div").Annotate(attrs...)
	} else {
		return xml.Element[T]("div")
	}
}

// Represents a description list. The element encloses a list of groups of terms (specified using the {{HTMLElement("dt")}} element) and descriptions (provided by {{HTMLElement("dd")}} elements). Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).
func (s *HTML[T]) Dl(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("dl").Annotate(attrs...)
	} else {
		return xml.Element[T]("dl")
	}
}

// Specifies a term in a description or definition list, and as such must be used inside a {{HTMLElement("dl")}} element. It is usually followed by a {{HTMLElement("dd")}} element; however, multiple `<dt>` elements in a row indicate several terms that are all defined by the immediate next {{HTMLElement("dd")}} element.
func (s *HTML[T]) Dt(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("dt").Annotate(attrs...)
	} else {
		return xml.Element[T]("dt")
	}
}

// Represents a caption or legend describing the rest of the contents of its parent {{HTMLElement("figure")}} element.
func (s *HTML[T]) Figcaption(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("figcaption").Annotate(attrs...)
	} else {
		return xml.Element[T]("figcaption")
	}
}

// Represents self-contained content, potentially with an optional caption, which is specified using the {{HTMLElement("figcaption")}} element. The figure, its caption, and its contents are referenced as a single unit.
func (s *HTML[T]) Figure(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("figure").Annotate(attrs...)
	} else {
		return xml.Element[T]("figure")
	}
}

// Represents a thematic break between paragraph-level elements: for example, a change of scene in a story, or a shift of topic within a section.
func (s *HTML[T]) Hr(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("hr").Annotate(attrs...)
	} else {
		return xml.Element[T]("hr")
	}
}

// Represents an item in a list. It must be contained in a parent element: an ordered list ({{HTMLElement("ol")}}), an unordered list ({{HTMLElement("ul")}}), or a menu ({{HTMLElement("menu")}}). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.
func (s *HTML[T]) Li(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("li").Annotate(attrs...)
	} else {
		return xml.Element[T]("li")
	}
}

// A semantic alternative to {{HTMLElement("ul")}}, but treated by browsers (and exposed through the accessibility tree) as no different than {{HTMLElement("ul")}}. It represents an unordered list of items (which are represented by {{HTMLElement("li")}} elements).
func (s *HTML[T]) Menu(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("menu").Annotate(attrs...)
	} else {
		return xml.Element[T]("menu")
	}
}

// Represents an ordered list of items — typically rendered as a numbered list.
func (s *HTML[T]) Ol(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("ol").Annotate(attrs...)
	} else {
		return xml.Element[T]("ol")
	}
}

// Represents a paragraph. Paragraphs are usually represented in visual media as blocks of text separated from adjacent blocks by blank lines and/or first-line indentation, but HTML paragraphs can be any structural grouping of related content, such as images or form fields.
func (s *HTML[T]) P(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("p").Annotate(attrs...)
	} else {
		return xml.Element[T]("p")
	}
}

// Represents preformatted text which is to be presented exactly as written in the HTML file. The text is typically rendered using a non-proportional, or [monospaced](https://en.wikipedia.org/wiki/Monospaced_font), font. Whitespace inside this element is displayed as written.
func (s *HTML[T]) Pre(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("pre").Annotate(attrs...)
	} else {
		return xml.Element[T]("pre")
	}
}

// Represents an unordered list of items, typically rendered as a bulleted list.
func (s *HTML[T]) Ul(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("ul").Annotate(attrs...)
	} else {
		return xml.Element[T]("ul")
	}
}

// Together with its `href` attribute, creates a hyperlink to web pages, files, email addresses, locations within the current page, or anything else a URL can address.
func (s *HTML[T]) A(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("a").Annotate(attrs...)
	} else {
		return xml.Element[T]("a")
	}
}

// Represents an abbreviation or acronym.
func (s *HTML[T]) Abbr(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("abbr").Annotate(attrs...)
	} else {
		return xml.Element[T]("abbr")
	}
}

// Used to draw the reader's attention to the element's contents, which are not otherwise granted special importance. This was formerly known as the Boldface element, and most browsers still draw the text in boldface. However, you should not use `<b>` for styling text or granting importance. If you wish to create boldface text, you should use the CSS {{cssxref("font-weight")}} property. If you wish to indicate an element is of special importance, you should use the {{HTMLElement("strong")}} element.
func (s *HTML[T]) B(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("b").Annotate(attrs...)
	} else {
		return xml.Element[T]("b")
	}
}

// Tells the browser's bidirectional algorithm to treat the text it contains in isolation from its surrounding text. It's particularly useful when a website dynamically inserts some text and doesn't know the directionality of the text being inserted.
func (s *HTML[T]) Bdi(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("bdi").Annotate(attrs...)
	} else {
		return xml.Element[T]("bdi")
	}
}

// Overrides the current directionality of text, so that the text within is rendered in a different direction.
func (s *HTML[T]) Bdo(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("bdo").Annotate(attrs...)
	} else {
		return xml.Element[T]("bdo")
	}
}

// Produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.
func (s *HTML[T]) Br(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("br").Annotate(attrs...)
	} else {
		return xml.Element[T]("br")
	}
}

// Used to mark up the title of a cited creative work. The reference may be in an abbreviated form according to context-appropriate conventions related to citation metadata.
func (s *HTML[T]) Cite(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("cite").Annotate(attrs...)
	} else {
		return xml.Element[T]("cite")
	}
}

// Displays its contents styled in a fashion intended to indicate that the text is a short fragment of computer code. By default, the content text is displayed using the user agent's default monospace font.
func (s *HTML[T]) Code(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("code").Annotate(attrs...)
	} else {
		return xml.Element[T]("code")
	}
}

// Links a given piece of content with a machine-readable translation. If the content is time- or date-related, the `<time>` element must be used.
func (s *HTML[T]) Data(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("data").Annotate(attrs...)
	} else {
		return xml.Element[T]("data")
	}
}

// Used to indicate the term being defined within the context of a definition phrase or sentence. The ancestor {{HTMLElement("p")}} element, the {{HTMLElement("dt")}}/{{HTMLElement("dd")}} pairing, or the nearest section ancestor of the `<dfn>` element, is considered to be the definition of the term.
func (s *HTML[T]) Dfn(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("dfn").Annotate(attrs...)
	} else {
		return xml.Element[T]("dfn")
	}
}

// Marks text that has stress emphasis. The `<em>` element can be nested, with each nesting level indicating a greater degree of emphasis.
func (s *HTML[T]) Em(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("em").Annotate(attrs...)
	} else {
		return xml.Element[T]("em")
	}
}

// Represents a range of text that is set off from the normal text for some reason, such as idiomatic text, technical terms, and taxonomical designations, among others. Historically, these have been presented using italicized type, which is the original source of the `<i>` naming of this element.
func (s *HTML[T]) I(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("i").Annotate(attrs...)
	} else {
		return xml.Element[T]("i")
	}
}

// Represents a span of inline text denoting textual user input from a keyboard, voice input, or any other text entry device. By convention, the user agent defaults to rendering the contents of a `<kbd>` element using its default monospace font, although this is not mandated by the HTML standard.
func (s *HTML[T]) Kbd(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("kbd").Annotate(attrs...)
	} else {
		return xml.Element[T]("kbd")
	}
}

// Represents text which is marked or highlighted for reference or notation purposes due to the marked passage's relevance in the enclosing context.
func (s *HTML[T]) Mark(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("mark").Annotate(attrs...)
	} else {
		return xml.Element[T]("mark")
	}
}

// Indicates that the enclosed text is a short inline quotation. Most modern browsers implement this by surrounding the text in quotation marks. This element is intended for short quotations that don't require paragraph breaks; for long quotations use the {{HTMLElement("blockquote")}} element.
func (s *HTML[T]) Q(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("q").Annotate(attrs...)
	} else {
		return xml.Element[T]("q")
	}
}

// Used to provide fall-back parentheses for browsers that do not support the display of ruby annotations using the {{HTMLElement("ruby")}} element. One `<rp>` element should enclose each of the opening and closing parentheses that wrap the {{HTMLElement("rt")}} element that contains the annotation's text.
func (s *HTML[T]) Rp(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("rp").Annotate(attrs...)
	} else {
		return xml.Element[T]("rp")
	}
}

// Specifies the ruby text component of a ruby annotation, which is used to provide pronunciation, translation, or transliteration information for East Asian typography. The `<rt>` element must always be contained within a {{HTMLElement("ruby")}} element.
func (s *HTML[T]) Rt(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("rt").Annotate(attrs...)
	} else {
		return xml.Element[T]("rt")
	}
}

// Represents small annotations that are rendered above, below, or next to base text, usually used for showing the pronunciation of East Asian characters. It can also be used for annotating other kinds of text, but this usage is less common.
func (s *HTML[T]) Ruby(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("ruby").Annotate(attrs...)
	} else {
		return xml.Element[T]("ruby")
	}
}

// Renders text with a strikethrough, or a line through it. Use the `<s>` element to represent things that are no longer relevant or no longer accurate. However, `<s>` is not appropriate when indicating document edits; for that, use the {{HTMLElement("del")}} and {{HTMLElement("ins")}} elements, as appropriate.
func (s *HTML[T]) S(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("s").Annotate(attrs...)
	} else {
		return xml.Element[T]("s")
	}
}

// Used to enclose inline text which represents sample (or quoted) output from a computer program. Its contents are typically rendered using the browser's default monospaced font (such as [Courier](<https://en.wikipedia.org/wiki/Courier_(typeface)>) or Lucida Console).
func (s *HTML[T]) Samp(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("samp").Annotate(attrs...)
	} else {
		return xml.Element[T]("samp")
	}
}

// Represents side-comments and small print, like copyright and legal text, independent of its styled presentation. By default, it renders text within it one font size smaller, such as from `small` to `x-small`.
func (s *HTML[T]) Small(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("small").Annotate(attrs...)
	} else {
		return xml.Element[T]("small")
	}
}

// A generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the `class` or `id` attributes), or because they share attribute values, such as `lang`. It should be used only when no other semantic element is appropriate. `<span>` is very much like a div element, but div is a [block-level element](/en-US/docs/Glossary/Block-level_content) whereas a `<span>` is an [inline-level element](/en-US/docs/Glossary/Inline-level_content).
func (s *HTML[T]) Span(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("span").Annotate(attrs...)
	} else {
		return xml.Element[T]("span")
	}
}

// Indicates that its contents have strong importance, seriousness, or urgency. Browsers typically render the contents in bold type.
func (s *HTML[T]) Strong(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("strong").Annotate(attrs...)
	} else {
		return xml.Element[T]("strong")
	}
}

// Specifies inline text which should be displayed as subscript for solely typographical reasons. Subscripts are typically rendered with a lowered baseline using smaller text.
func (s *HTML[T]) Sub(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("sub").Annotate(attrs...)
	} else {
		return xml.Element[T]("sub")
	}
}

// Specifies inline text which is to be displayed as superscript for solely typographical reasons. Superscripts are usually rendered with a raised baseline using smaller text.
func (s *HTML[T]) Sup(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("sup").Annotate(attrs...)
	} else {
		return xml.Element[T]("sup")
	}
}

// Represents a specific period in time. It may include the `datetime` attribute to translate dates into machine-readable format, allowing for better search engine results or custom features such as reminders.
func (s *HTML[T]) Time(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("time").Annotate(attrs...)
	} else {
		return xml.Element[T]("time")
	}
}

// Represents a span of inline text which should be rendered in a way that indicates that it has a non-textual annotation. This is rendered by default as a single solid underline but may be altered using CSS.
func (s *HTML[T]) U(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("u").Annotate(attrs...)
	} else {
		return xml.Element[T]("u")
	}
}

// Represents the name of a variable in a mathematical expression or a programming context. It's typically presented using an italicized version of the current typeface, although that behavior is browser-dependent.
func (s *HTML[T]) Var(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("var").Annotate(attrs...)
	} else {
		return xml.Element[T]("var")
	}
}

// Represents a word break opportunity—a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.
func (s *HTML[T]) Wbr(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("wbr").Annotate(attrs...)
	} else {
		return xml.Element[T]("wbr")
	}
}

// Defines an area inside an image map that has predefined clickable areas. An _image map_ allows geometric areas on an image to be associated with {{Glossary("Hyperlink", "hyperlink")}}.
func (s *HTML[T]) Area(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("area").Annotate(attrs...)
	} else {
		return xml.Element[T]("area")
	}
}

// Used to embed sound content in documents. It may contain one or more audio sources, represented using the `src` attribute or the source element: the browser will choose the most suitable one. It can also be the destination for streamed media, using a {{domxref("MediaStream")}}.
func (s *HTML[T]) Audio(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("audio").Annotate(attrs...)
	} else {
		return xml.Element[T]("audio")
	}
}

// Embeds an image into the document.
func (s *HTML[T]) Img(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("img").Annotate(attrs...)
	} else {
		return xml.Element[T]("img")
	}
}

// Used with {{HTMLElement("area")}} elements to define an image map (a clickable link area).
func (s *HTML[T]) Map(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("map").Annotate(attrs...)
	} else {
		return xml.Element[T]("map")
	}
}

// Used as a child of the media elements, audio and video. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in [WebVTT format](/en-US/docs/Web/API/WebVTT_API) (`.vtt` files)—Web Video Text Tracks.
func (s *HTML[T]) Track(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("track").Annotate(attrs...)
	} else {
		return xml.Element[T]("track")
	}
}

// Embeds a media player which supports video playback into the document. You can also use `<video>` for audio content, but the audio element may provide a more appropriate user experience.
func (s *HTML[T]) Video(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("video").Annotate(attrs...)
	} else {
		return xml.Element[T]("video")
	}
}

// Embeds external content at the specified point in the document. This content is provided by an external application or other source of interactive content such as a browser plug-in.
func (s *HTML[T]) Embed(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("embed").Annotate(attrs...)
	} else {
		return xml.Element[T]("embed")
	}
}

// Represents a nested browsing context, like `<iframe>` but with more native privacy features built in.
func (s *HTML[T]) Fencedframe(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("fencedframe").Annotate(attrs...)
	} else {
		return xml.Element[T]("fencedframe")
	}
}

// Represents a nested browsing context, embedding another HTML page into the current one.
func (s *HTML[T]) Iframe(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("iframe").Annotate(attrs...)
	} else {
		return xml.Element[T]("iframe")
	}
}

// Represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.
func (s *HTML[T]) Object(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("object").Annotate(attrs...)
	} else {
		return xml.Element[T]("object")
	}
}

// Contains zero or more {{HTMLElement("source")}} elements and one {{HTMLElement("img")}} element to offer alternative versions of an image for different display/device scenarios.
func (s *HTML[T]) Picture(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("picture").Annotate(attrs...)
	} else {
		return xml.Element[T]("picture")
	}
}

// Enables the embedding of another HTML page into the current one to enable smoother navigation into new pages.
func (s *HTML[T]) Portal(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("portal").Annotate(attrs...)
	} else {
		return xml.Element[T]("portal")
	}
}

// Specifies multiple media resources for the picture, the audio element, or the video element. It is a void element, meaning that it has no content and does not have a closing tag. It is commonly used to offer the same media content in multiple file formats in order to provide compatibility with a broad range of browsers given their differing support for [image file formats](/en-US/docs/Web/Media/Formats/Image_types) and [media file formats](/en-US/docs/Web/Media/Formats).
func (s *HTML[T]) Source(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("source").Annotate(attrs...)
	} else {
		return xml.Element[T]("source")
	}
}

// Container defining a new coordinate system and [viewport](/en-US/docs/Web/SVG/Attribute/viewBox). It is used as the outermost element of SVG documents, but it can also be used to embed an SVG fragment inside an SVG or HTML document.
func (s *HTML[T]) Svg(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("svg").Annotate(attrs...)
	} else {
		return xml.Element[T]("svg")
	}
}

// The top-level element in MathML. Every valid MathML instance must be wrapped in it. In addition, you must not nest a second `<math>` element in another, but you can have an arbitrary number of other child elements in it.
func (s *HTML[T]) Math(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("math").Annotate(attrs...)
	} else {
		return xml.Element[T]("math")
	}
}

// Container element to use with either the [canvas scripting API](/en-US/docs/Web/API/Canvas_API) or the [WebGL API](/en-US/docs/Web/API/WebGL_API) to draw graphics and animations.
func (s *HTML[T]) Canvas(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("canvas").Annotate(attrs...)
	} else {
		return xml.Element[T]("canvas")
	}
}

// Defines a section of HTML to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.
func (s *HTML[T]) Noscript(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("noscript").Annotate(attrs...)
	} else {
		return xml.Element[T]("noscript")
	}
}

// Used to embed executable code or data; this is typically used to embed or refer to JavaScript code. The `<script>` element can also be used with other languages, such as [WebGL](/en-US/docs/Web/API/WebGL_API)'s GLSL shader programming language and [JSON](/en-US/docs/Glossary/JSON).
func (s *HTML[T]) Script(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("script").Annotate(attrs...)
	} else {
		return xml.Element[T]("script")
	}
}

// Represents a range of text that has been deleted from a document. This can be used when rendering "track changes" or source code diff information, for example. The `<ins>` element can be used for the opposite purpose: to indicate text that has been added to the document.
func (s *HTML[T]) Del(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("del").Annotate(attrs...)
	} else {
		return xml.Element[T]("del")
	}
}

// Represents a range of text that has been added to a document. You can use the `<del>` element to similarly represent a range of text that has been deleted from the document.
func (s *HTML[T]) Ins(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("ins").Annotate(attrs...)
	} else {
		return xml.Element[T]("ins")
	}
}

// Specifies the caption (or title) of a table.
func (s *HTML[T]) Caption(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("caption").Annotate(attrs...)
	} else {
		return xml.Element[T]("caption")
	}
}

// Defines one or more columns in a column group represented by its implicit or explicit parent {{HTMLElement("colgroup")}} element. The `<col>` element is only valid as a child of a {{HTMLElement("colgroup")}} element that has no [`span`](/en-US/docs/Web/HTML/Element/colgroup#span) attribute defined.
func (s *HTML[T]) Col(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("col").Annotate(attrs...)
	} else {
		return xml.Element[T]("col")
	}
}

// Defines a group of columns within a table.
func (s *HTML[T]) Colgroup(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("colgroup").Annotate(attrs...)
	} else {
		return xml.Element[T]("colgroup")
	}
}

// Represents tabular data—that is, information presented in a two-dimensional table comprised of rows and columns of cells containing data.
func (s *HTML[T]) Table(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("table").Annotate(attrs...)
	} else {
		return xml.Element[T]("table")
	}
}

// Encapsulates a set of table rows ({{HTMLElement("tr")}} elements), indicating that they comprise the body of a table's (main) data.
func (s *HTML[T]) Tbody(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("tbody").Annotate(attrs...)
	} else {
		return xml.Element[T]("tbody")
	}
}

// A child of the {{HTMLElement("tr")}} element, it defines a cell of a table that contains data.
func (s *HTML[T]) Td(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("td").Annotate(attrs...)
	} else {
		return xml.Element[T]("td")
	}
}

// Encapsulates a set of table rows ({{HTMLElement("tr")}} elements), indicating that they comprise the foot of a table with information about the table's columns. This is usually a summary of the columns, e.g., a sum of the given numbers in a column.
func (s *HTML[T]) Tfoot(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("tfoot").Annotate(attrs...)
	} else {
		return xml.Element[T]("tfoot")
	}
}

// A child of the {{HTMLElement("tr")}} element, it defines a cell as the header of a group of table cells. The nature of this group can be explicitly defined by the [`scope`](/en-US/docs/Web/HTML/Element/th#scope) and [`headers`](/en-US/docs/Web/HTML/Element/th#headers) attributes.
func (s *HTML[T]) Th(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("th").Annotate(attrs...)
	} else {
		return xml.Element[T]("th")
	}
}

// Encapsulates a set of table rows ({{HTMLElement("tr")}} elements), indicating that they comprise the head of a table with information about the table's columns. This is usually in the form of column headers ({{HTMLElement("th")}} elements).
func (s *HTML[T]) Thead(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("thead").Annotate(attrs...)
	} else {
		return xml.Element[T]("thead")
	}
}

// Defines a row of cells in a table. The row's cells can then be established using a mix of {{HTMLElement("td")}} (data cell) and {{HTMLElement("th")}} (header cell) elements.
func (s *HTML[T]) Tr(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("tr").Annotate(attrs...)
	} else {
		return xml.Element[T]("tr")
	}
}

// An interactive element activated by a user with a mouse, keyboard, finger, voice command, or other assistive technology. Once activated, it performs an action, such as submitting a [form](/en-US/docs/Learn_web_development/Extensions/Forms) or opening a dialog.
func (s *HTML[T]) Button(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("button").Annotate(attrs...)
	} else {
		return xml.Element[T]("button")
	}
}

// Contains a set of {{HTMLElement("option")}} elements that represent the permissible or recommended options available to choose from within other controls.
func (s *HTML[T]) Datalist(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("datalist").Annotate(attrs...)
	} else {
		return xml.Element[T]("datalist")
	}
}

// Used to group several controls as well as labels ({{HTMLElement("label")}}) within a web form.
func (s *HTML[T]) Fieldset(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("fieldset").Annotate(attrs...)
	} else {
		return xml.Element[T]("fieldset")
	}
}

// Represents a document section containing interactive controls for submitting information.
func (s *HTML[T]) Form(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("form").Annotate(attrs...)
	} else {
		return xml.Element[T]("form")
	}
}

// Used to create interactive controls for web-based forms to accept data from the user; a wide variety of types of input data and control widgets are available, depending on the device and user agent. The `<input>` element is one of the most powerful and complex in all of HTML due to the sheer number of combinations of input types and attributes.
func (s *HTML[T]) Input(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("input").Annotate(attrs...)
	} else {
		return xml.Element[T]("input")
	}
}

// Represents a caption for an item in a user interface.
func (s *HTML[T]) Label(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("label").Annotate(attrs...)
	} else {
		return xml.Element[T]("label")
	}
}

// Represents a caption for the content of its parent {{HTMLElement("fieldset")}}.
func (s *HTML[T]) Legend(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("legend").Annotate(attrs...)
	} else {
		return xml.Element[T]("legend")
	}
}

// Represents either a scalar value within a known range or a fractional value.
func (s *HTML[T]) Meter(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("meter").Annotate(attrs...)
	} else {
		return xml.Element[T]("meter")
	}
}

// Creates a grouping of options within a {{HTMLElement("select")}} element.
func (s *HTML[T]) Optgroup(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("optgroup").Annotate(attrs...)
	} else {
		return xml.Element[T]("optgroup")
	}
}

// Used to define an item contained in a select, an {{HTMLElement("optgroup")}}, or a {{HTMLElement("datalist")}} element. As such, `<option>` can represent menu items in popups and other lists of items in an HTML document.
func (s *HTML[T]) Option(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("option").Annotate(attrs...)
	} else {
		return xml.Element[T]("option")
	}
}

// Container element into which a site or app can inject the results of a calculation or the outcome of a user action.
func (s *HTML[T]) Output(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("output").Annotate(attrs...)
	} else {
		return xml.Element[T]("output")
	}
}

// Displays an indicator showing the completion progress of a task, typically displayed as a progress bar.
func (s *HTML[T]) Progress(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("progress").Annotate(attrs...)
	} else {
		return xml.Element[T]("progress")
	}
}

// Represents a control that provides a menu of options.
func (s *HTML[T]) Select(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("select").Annotate(attrs...)
	} else {
		return xml.Element[T]("select")
	}
}

// Represents a multi-line plain-text editing control, useful when you want to allow users to enter a sizeable amount of free-form text, for example, a comment on a review or feedback form.
func (s *HTML[T]) Textarea(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("textarea").Annotate(attrs...)
	} else {
		return xml.Element[T]("textarea")
	}
}

// Creates a disclosure widget in which information is visible only when the widget is toggled into an "open" state. A summary or label must be provided using the {{HTMLElement("summary")}} element.
func (s *HTML[T]) Details(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("details").Annotate(attrs...)
	} else {
		return xml.Element[T]("details")
	}
}

// Represents a dialog box or other interactive component, such as a dismissible alert, inspector, or subwindow.
func (s *HTML[T]) Dialog(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("dialog").Annotate(attrs...)
	} else {
		return xml.Element[T]("dialog")
	}
}

// Specifies a summary, caption, or legend for a details element's disclosure box. Clicking the `<summary>` element toggles the state of the parent {{HTMLElement("details")}} element open and closed.
func (s *HTML[T]) Summary(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("summary").Annotate(attrs...)
	} else {
		return xml.Element[T]("summary")
	}
}

// Part of the [Web Components](/en-US/docs/Web/API/Web_components) technology suite, this element is a placeholder inside a web component that you can fill with your own markup, which lets you create separate DOM trees and present them together.
func (s *HTML[T]) Slot(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("slot").Annotate(attrs...)
	} else {
		return xml.Element[T]("slot")
	}
}

// A mechanism for holding HTML that is not to be rendered immediately when a page is loaded but may be instantiated subsequently during runtime using JavaScript.
func (s *HTML[T]) Template(attrs ...dagger.Buildable[any, xml.XMLNode[T]]) *dagger.Node[any, xml.XMLNode[T]] {
	if len(attrs) > 0 {
		return xml.Element[T]("template").Annotate(attrs...)
	} else {
		return xml.Element[T]("template")
	}
}
