# go-klingon
Translate a name written in English to Klingon and find out  its species

[![Build Status](https://travis-ci.org/MarinX/go-klingon.svg?branch=master)](https://travis-ci.org/MarinX/go-klingon)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarinX/go-klingon)](https://goreportcard.com/report/github.com/MarinX/go-klingon)
[![License MIT](https://img.shields.io/badge/license-MIT-brightgreen.svg?style=flat)](LICENSE)

Note: The STAPI library does not have implementations of all STAPI REST resources[WIP]. PRs for new resources and endpoints are welcome, or you can simply implement some yourself as-you-go.

## Use it as cli tool

```go
go get github.com/MarinX/go-klingon
```

Usage:
```sh
Usage: ./go-klingon [OPTIONS] argument
  -apikey string
        stapi api key
```

## Use it as a library

### Translate

```go
go get github.com/MarinX/go-klingon/translate
```

Example:

```go
package main

import (
	"fmt"
	"github.com/MarinX/go-klingon/translate"
)

func main() {
    n := "Uhura"
    val, err := translate.New(n).Klingon()
	if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println("hex value:", val)
}
```

### STAPI

```go
go get github.com/MarinX/go-klingon/stapi
```

Example:
```go
package main

import (
	"fmt"
	"github.com/MarinX/go-klingon/stapi"
)

func main() {
    client := stapi.New("", nil)
    ch, err :=client.Character.Search(struct{
        Name string `url:"name"`
    }{"uhura"})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("%#v\n", ch)
}
```

Not all endpoints are implemented right now. 
In those case, you can use Get/Post method and point your model


```go
path := "books"
resource := new(string)
options := struct {
    Name string `url:"name"`
}{"book1"}

err := client.Get(path, options, resource)
if err != nil {
	fmt.Println(err)
	return
}
```

## Contributing
PR's are welcome. Please read [CONTRIBUTING.md](https://github.com/MarinX/go-klingon/blob/master/CONTRIBUTING.md) for more info

## License
MIT