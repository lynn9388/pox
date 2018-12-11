# PoX

[![Go Report Card](https://goreportcard.com/badge/github.com/lynn9388/pox)](https://goreportcard.com/report/github.com/lynn9388/pox)
[![Build Status](https://travis-ci.com/lynn9388/pox.svg?branch=master)](https://travis-ci.com/lynn9388/pox)

Some simple PoX implementations.

## PoW

[![GoDoc](https://godoc.org/github.com/lynn9388/pox/pow?status.svg)](https://godoc.org/github.com/lynn9388/pox/pow)

### Example

```go
pow := NewPoW([]byte("lynn9388"), 10, 0)
pow.Compute()
if !pow.IsValid() {
	fmt.Println("PoW is not valid.")
}
```
