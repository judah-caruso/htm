package htm_test

import (
	"testing"

	. "github.com/judah-caruso/htm"
)

func TestRendering(t *testing.T) {
	cases := []struct {
		given    Element
		expected string
	}{
		{given: Div(H1(Text("first"))), expected: "<div><h1>first</h1></div>"},
		{given: Fragment(H1(Text("second"))), expected: "<h1>second</h1>"},
		{given: Fragment(Class("foo"), H1(Text("third"))), expected: "<h1>third</h1>"},
		{given: Div(Id("id"), Class("class"), Text("fourth")), expected: `<div id="id" class="class">fourth</div>`},
		{given: If(false, Text("true"), Text("false")), expected: "false"},
		{given: If(true, Text("true"), Text("false")), expected: "true"},
		{given: Link(".", "."), expected: `<link rel="." href="."/>`},
	}

	for _, c := range cases {
		given := c.given.Render()
		if c.expected != given {
			t.Fatalf("expected %q, given %q", c.expected, given)
		}
	}
}
