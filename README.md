# ginie

`ginie` is a simple Web framework, just like minimum viable version of Gin.

## Feature

- wrapped the basic interface and types for ginie data structure

- support return JSON/HTML/String types

- support dynamic route and route group control

- extensible middleware

- support static resource service

- support HTML template rendering

## Usage

```go
package main

import "github.com/i0Ek3/ginie"

func main() {
    r := ginie.New()

    r.GET("/", func(c *ginie.Context) {
        c.String(http.StatusOK, "Hi there, this is Ginie here!\n")
    })

    r.Run(":8888")
} 
```

## Credit

[geektutu](https://github.com/geektutu)
