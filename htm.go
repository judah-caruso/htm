package htm

import (
	"fmt"
	"strings"
)

// Element represents a piece of html code.
type Element interface {
	Render() string
}

// Make returns a new html element. Useful for creating new/non-standard elements.
func Make(tag string, body ...Element) Element {
	return build(tag, false, false).withBody(body)
}

// MakeSelfClosing returns a new self closing html element.
func MakeSelfClosing(tag string, body ...Element) Element {
	return build(tag, true, false).withBody(body)
}

// Empty returns an empty element. Useful in conditionals.
func Empty() Element {
	return &builder{}
}

// Fragment returns a new element only used to contain others elements.
func Fragment(body ...Element) Element {
	return build("__fragment", false, true).withBody(body)
}

// Attr returns a new attribute. Useful for creating new/non-standard attributes.
func Attr(key, value string) Element {
	return build("__attr", false, false).withAttr(key, value)
}

// Id returns a new id attribute.
func Id(id string) Element {
	return Attr("id", id)
}

// Class returns a new class attribute.
func Class(class string) Element {
	return Attr("class", class)
}

// Name returns a new name attribute.
func Name(name string) Element {
	return Attr("name", name)
}

// Type returns a new type attribute.
func Type(typ string) Element {
	return Attr("type", typ)
}

// Style returns a new style attribute.
func Style(style string) Element {
	return Attr("style", style)
}

// Rel returns a new rel attribute.
func Rel(rel string) Element {
	return Attr("rel", rel)
}

// Href returns a new href attribute.
func Href(href string) Element {
	return Attr("href", href)
}

// Alt returns a new alt attribute.
func Alt(alt string) Element {
	return Attr("alt", alt)
}

// Src returns a new src attribute.
func Src(src string) Element {
	return Attr("src", src)
}

// If returns an element based on the given condition.
func If(cond bool, whenTrue, whenFalse Element) Element {
	if cond {
		return whenTrue
	} else {
		return whenFalse
	}
}

// When returns an element only if the given condition is true.
// Equivalent to: If(cond, ..., Empty())
func When(cond bool, whenTrue Element) Element {
	return If(cond, whenTrue, Empty())
}

// Map takes a list of values and returns their concatenation as Fragment.
func Map[T any](values []T, iter func(T) Element) Element {
	return MapIdx(values, func(v T, _ int) Element {
		return iter(v)
	})
}

// MapIdx takes a list of values and returns their concatenation as Fragment.
func MapIdx[T any](values []T, iter func(T, int) Element) Element {
	if values == nil {
		return Empty()
	}

	body := make([]Element, len(values))
	for i, v := range values {
		body[i] = iter(v, i)
	}

	return Fragment(body...)
}

// Join transfers the elements and attributes of the children to the parent. Useful for reusing Fragments.
func Join(parent Element, children ...Element) Element {
	if b, ok := parent.(*builder); ok {
		b.withBody(children)
	}

	return parent
}

// Meta returns a new html <meta> element.
func Meta(body ...Element) Element {
	return Make("meta", body...)
}

// Title returns a new html <title> element from text (can be formatted).
func Title(title string, args ...any) Element {
	return Make("title", Text(title, args...))
}

// Text returns a new html element from text (can be formatted).
func Text(txt string, args ...any) Element {
	return text(fmt.Sprintf(txt, args...))
}

// Html returns a new html <html> element.
func Html(body ...Element) Element {
	return Make("html", body...)
}

// Head returns a new html <head> element.
func Head(body ...Element) Element {
	return Make("head", body...)
}

// Link returns a new html <link rel="..." href="..."> element.
func Link(rel, href string) Element {
	return MakeSelfClosing("link", Rel(rel), Href(href))
}

// Body returns a new html <body> element.
func Body(body ...Element) Element {
	return Make("body", body...)
}

// Main returns a new html <main> element.
func Main(body ...Element) Element {
	return Make("main", body...)
}

// Span returns a new html <span> element.
func Span(body ...Element) Element {
	return Make("span", body...)
}

// H1 returns a new html <h1> element.
func H1(body ...Element) Element {
	return Make("h1", body...)
}

// H2 returns a new html <h2> element.
func H2(body ...Element) Element {
	return Make("h2", body...)
}

// H3 returns a new html <h3> element.
func H3(body ...Element) Element {
	return Make("h3", body...)
}

// H4 returns a new html <h4> element.
func H4(body ...Element) Element {
	return Make("h4", body...)
}

// Div returns a new html <div> element.
func Div(body ...Element) Element {
	return Make("div", body...)
}

// Button returns a new html <button> element.
func Button(body ...Element) Element {
	return Make("button", body...)
}

// List returns a new html <ol> or <ul> element.
func List(ordered bool, items ...Element) Element {
	if ordered {
		return Make("ol", items...)
	}

	return Make("ul", items...)
}

// ListItem returns a new html <li> element.
func ListItem(body ...Element) Element {
	return Make("li", body...)
}

// A returns a new html <a href="..."> element.
func A(href string, body ...Element) Element {
	return Make("a", append(body, Href(href))...)
}

// Img returns a new html <img src="..."/> element.
func Img(src string, attributes ...Element) Element {
	return MakeSelfClosing("img", append(attributes, Src(src))...)
}

// Br returns a new html <br/> element.
func Br() Element {
	return MakeSelfClosing("br")
}

// Hr returns a new html <hr/> element.
func Hr() Element {
	return MakeSelfClosing("hr")
}

// Input returns a new html <input> element.
func Input(body ...Element) Element {
	return Make("input", body...)
}

// Textarea returns a new html <textarea> element.
func Textarea(body ...Element) Element {
	return Make("textarea", body...)
}

// Code returns a new html <code> element.
func Code(body ...Element) Element {
	return Make("code", body...)
}

// Pre returns a new html <pre> element.
func Pre(body ...Element) Element {
	return Make("pre", body...)
}

// Script returns a new html <script src="..."> element.
func Script(src string) Element {
	return Make("script", Src(src))
}

// Form returns a new html <form> element.
func Form(body ...Element) Element {
	return Make("form", body...)
}

// Select returns a new html <select> element.
func Select(body ...Element) Element {
	return Make("select", body...)
}

// Option returns a new html <option> element.
func Option(body ...Element) Element {
	return Make("option", body...)
}

// Label returns a new html <label for="..."> element.
func Label(forName string, body ...Element) Element {
	return Join(Make("label", body...), Attr("for", forName))
}

type text string

func (t text) Render() string {
	return string(t)
}

type attribute struct {
	name  string
	value string
}

func (a *attribute) Render() string {
	if a == nil {
		return ""
	}

	return fmt.Sprintf("%s='%s'", a.name, a.value)
}

type builder struct {
	tag         string
	selfClosing bool
	fragment    bool
	attrs       []Element
	body        []Element
}

func (b *builder) Render() string {
	if b == nil || len(b.tag) == 0 {
		return ""
	}

	var sb strings.Builder

	if b.fragment {
		for _, el := range b.body {
			sb.WriteString(el.Render())
		}

		return sb.String()
	}

	fmt.Fprintf(&sb, "<%s", b.tag)

	if len(b.attrs) > 0 {
		sb.WriteByte(' ')
		for i, attr := range b.attrs {
			a := attr.(*attribute)
			sb.WriteString(a.Render())

			if i < len(b.attrs)-1 {
				sb.WriteString(" ")
			}
		}
	}

	if b.selfClosing {
		sb.WriteString("/>")
		return sb.String()
	}

	sb.WriteString(">")

	for _, el := range b.body {
		if el == nil {
			continue
		}

		sb.WriteString(el.Render())
	}

	fmt.Fprintf(&sb, "</%s>", b.tag)

	return sb.String()
}

func (b *builder) withBody(body []Element) Element {
	for _, el := range body {
		if t, ok := el.(*builder); ok {
			if t.tag == "__attr" {
				b.attrs = append(b.attrs, t.attrs[0])
			} else {
				b.body = append(b.body, el)
			}
		} else {
			b.body = append(b.body, el)
		}
	}
	return b
}

func (b *builder) withAttr(name, value string) Element {
	b.attrs = append(b.attrs, &attribute{name: name, value: value})
	return b
}

func build(t string, selfClosing, fragment bool) *builder {
	return &builder{tag: t, selfClosing: selfClosing, fragment: fragment}
}
