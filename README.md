# forcedomain

Martini middleware/handler for redirect based on environment variable.

## Usage

Just add the middleware and run it with **DOMAIN** env set to your redirect url.

~~~ go
package main

import (
    "github.com/go-martini/martini"
    "github.com/mytrile/forcedomain"
)

func main() {
    m := martini.Classic()
    m.Use(forcedomain.ForceDomainRedirect())
    m.Get("/", func() string {
      return "Hello world!"
    })
    m.Run()
}

~~~

    $ DOMAIN=http://example.com main


## Meta

* Author  : Dimitar Kostov
* Email   : mitko.kostov@gmail.com
* Website : [http://mytrile.github.com](http://mytrile.github.com)
* Twitter : [http://twitter.com/mytrile](http://twitter.com/mytrile)