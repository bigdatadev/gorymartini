# GoryMartini

[Riemann](https://github.com/aphyr/riemann) middleware for martini framework.

Logs the following information:
* The time the request took
* The http status code
* The requested path

## Installation

To install the package for use in your own programs:

```
go get github.com/bigdatadev/gorymartini
```

## Getting Started

First we'll need to import the library:

```go
import (
    "github.com/bigdatadev/gorymartini"
	"github.com/go-martini/martini"
)
```

Next we'll need to create the client and martini handler:

```go
c, h := gorymartini.NewGoryMartini("localhost:5555")
err := c.Connect()
if c == nil || h == nil {
    panic(err)
}
```

Don't forget to close the client connection when you're done:

```go
defer c.Close()
```

Now we simply need to use the handler:

```go
m := martini.Classic()
m.Use(h)
m.Run()
```

## Contributing

Just send me a pull request. Please take a look at the project issues and see how you can help. Here are some tips:
- please add more tests.
- please check your syntax.

## Author

Christopher Gilbert

* Web: [http://cjgilbert.me](http://cjgilbert.me)
* Twitter: [@bigdatadev](https://twitter.com/bigdatadev)
* Linkedin: [/in/christophergilbert](https://www.linkedin.com/in/christophergilbert)

## Copyright

See [LICENSE](LICENSE) document