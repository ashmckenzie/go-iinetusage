# go-iinetusage

[![License](https://img.shields.io/badge/license-MIT-yellowgreen.svg?style=flat-square)][license]
[![Wercker](https://app.wercker.com/status/466a2faab09dfc4f21b3e5d6d771c9a2/s/master "wercker status")][wercker]

[license]: LICENSE.txt
[wercker]: https://app.wercker.com/project/byKey/466a2faab09dfc4f21b3e5d6d771c9a2

A golang library to grab your IINet usage.

## Install

```bash
$ go get -u github.com/ashmckenzie/go-iinetusage/...
```

## Example

```go
package main

import (
  "log"
  "os"

  "github.com/ashmckenzie/go-iinetusage"
)

func main() {
  iinet := iinetusage.New(os.Getenv("IINET_USERNAME"), os.Getenv("IINET_PASSWORD"))
  usage, err := iinet.GetUsage()
  if err != nil {
    log.Fatal(err)
  }
  log.Printf("%v", usage)
}
```

## License

The MIT License (MIT)

Copyright (c) 2016 Ash McKenzie.

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
