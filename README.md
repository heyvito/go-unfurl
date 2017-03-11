# go-unfurl
[![Build Status](https://travis-ci.org/victorgama/go-unfurl.svg?branch=master)](https://travis-ci.org/victorgama/go-unfurl)
[![Go Report Card](https://goreportcard.com/badge/github.com/victorgama/go-unfurl)](https://goreportcard.com/report/github.com/victorgama/go-unfurl)
[![codecov](https://codecov.io/gh/victorgama/go-unfurl/branch/master/graph/badge.svg)](https://codecov.io/gh/victorgama/go-unfurl)
[![GoDoc](https://godoc.org/github.com/victorgama/go-unfurl?status.svg)](https://godoc.org/github.com/victorgama/go-unfurl)

**go-unfurl** is a simple library that follow all redirects of a given URL.

```go
package main

import (
  "fmt"

  "github.com/victorgama/go-unfurl"
)

func main() {
  client := unfurl.NewClient()
  res, err := client.Process("http://goo.gl/g0DEfq")
  if err != nil {
    panic(err)
  }
  fmt.Println(res)
}

// => https://github.com
```

## Installing
1. Download and install it:
```
$ go get -u github.com/victorgama/go-unfurl
```
2. Import it in your code:
```
import "github.com/victorgama/go-unfurl"
```

## Usage

By default, this library will follow 20 redirects and return an `ErrTooManyRedirects` when this value is surpassed.
```go
client := unfurl.NewClient() // 20 max redirects
// ...or you can define a custom value
client := unfurl.NewClientWithOptions(unfurl.Options{MaxHops: 27})
```

## License
```
MIT License

Copyright (c) 2016 Victor Gama

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

```
