# HTM

An incredibly simple HTML creation tool for [Go](https://go.dev).

## Installation

```sh
go get github.com/judah-caruso/htm
```

## General Usage

```go
import (
   "fmt"
   "github.com/judah-caruso/htm"
)

// head returns the website's <head> element.
func head(title string) htm.Element {
   return htm.Head(
      htm.Meta(htm.Attr("charset", "utf-8")),
      htm.Meta(
         htm.Attr("name", "viewport"),
         htm.Attr("content", "width=device-width, initial-scale=1"),
      ),
      htm.Link("stylesheet", "foo.css"),
      htm.Title(title),
   )
}

type link struct {
   url, name string
}

// navbar returns an ordered list of links.
func navbar(links ...link) htm.Element {
   return htm.List(true, htm.Map(links, func(l link) htm.Element {
      return htm.ListItem(A(l.url, htm.Text(l.name)))
   }))
}

// Home returns the website's homepage.
func Home() htm.Element {
   return htm.Html(
      head("Home"),
      htm.Body(htm.Class("app"),
         navbar(link{"/", "about"}),
         htm.H1(htm.Class("page-title"),
            htm.Text("This is a small example!")
         ),
         // ...
      ),
   )
}

// User returns the website's user profile page.
func User(name string, loggedIn bool) htm.Element {
   return htm.Html(
      head(fmt.Sprintf("User @%s", name)),
      htm.Body(
         navbar(link{"/", "home"}, link{"/about", "about"}),
         htm.If(loggedIn,
            htm.H1(htm.Textf("Welcome back %s", name))
            htm.A("/login", htm.H1(htm.Text("Please login"))),
         ),
         // ...
      ),
   )
}

func main() {
   // ...
   out := Home().Render() // send to client, etc.
}
```

## Advanced Usage

```go
// CustomElement returns a new html <custom-el> element.
func CustomElement(body ...htm.Element) htm.Element {
   return htm.Make("custom-el", body)
}

// MyElement is a custom htm.Element that has its own state.
type MyElement struct {
   username string
   loggedIn bool
   // ...
}

func (m MyElement) Render() string {
   base := func(body ...htm.Element) string {
      return htm.Join(
         htm.Div(htm.Id("my-element"), htm.Class("element")),
         body...,
      ).Render()
   }

   if m.loggedIn {
      return base(
         htm.H1(htm.Textf("Welcome back, %s", m.username)),
      )
   } else {
      return base(
         htm.H1(htm.Text("Please login")),
      )
   }
}
```

## License

Public Domain. See [LICENSE](LICENSE) for details.
