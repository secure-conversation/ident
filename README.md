ident
=====

[![Go Doc](https://pkg.go.dev/badge/github.com/secure-conversation/ident.svg)](https://pkg.go.dev/github.com/secure-conversation/ident)
[![Go Report Card](https://goreportcard.com/badge/github.com/secure-conversation/ident)](https://goreportcard.com/report/github.com/secure-conversation/ident)

Create and manage unique identifiers based on uuidv7, with that implementation provided from `github.com/samborkent/uuidv7`

It is expected that the identifiers will be stored, so generating ids based on uuidv7 will reduce the cost of table scans and indexing.

Example:

```go
package main

import (
  "encoding/hex"
  "fmt"

  "github.com/secure-conversation/ident"
)

func main() {

  type myID [16]byte

  id := ident.NewID[myID]()

  fmt.Println(hex.EncodeToString(id[:]))
}
```
