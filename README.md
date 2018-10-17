# PoX

[![Build Status](https://travis-ci.com/lynn9388/pox.svg?branch=master)](https://travis-ci.com/lynn9388/pox)
[![Go Report Card](https://goreportcard.com/badge/github.com/lynn9388/pox)](https://goreportcard.com/report/github.com/lynn9388/pox)

Some simple PoX implementations.

## PoW

[![GoDoc](https://godoc.org/github.com/lynn9388/pox/pow?status.svg)](https://godoc.org/github.com/lynn9388/pox/pow)

### Example

```go
data := []byte("lynn9388")
targetBits := uint(10)

pow := NewPoW(data, targetBits)
pow.Compute()
if pow.Validate() == false {
	fmt.Println("failed to validate PoW")
}
```
