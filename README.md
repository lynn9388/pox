# PoX

[![Build Status](https://travis-ci.com/lynn9388/pox.svg?branch=master)](https://travis-ci.com/lynn9388/pox)

Some simple PoX implementations.

## PoW

[![GoDoc](https://godoc.org/github.com/lynn9388/pox/pow?status.svg)](https://godoc.org/github.com/lynn9388/pox/pow)

### Example

```go
func TestGetNonce(t *testing.T) {
	data := []byte("lynn9388")
	for difficulty := 1; difficulty < 6; difficulty++ {
		nonce := GetNonce(data, uint(difficulty))
		t.Log(hash(data, nonce))
	}
}
```

Output:

```text
pow_test.go:28: 0063095ed8e9f940f9dfaaa195cc57917b82a369486aa50febc71ac5cdfca21f
pow_test.go:28: 0063095ed8e9f940f9dfaaa195cc57917b82a369486aa50febc71ac5cdfca21f
pow_test.go:28: 000c59be69dd8c492e98697fce6080c0a2c51459886b286fb4cd7a49270586f2
pow_test.go:28: 0000e6079fe552a2145246f7ca966aaca598ab3aba6471643c1e96cd73c6987b
pow_test.go:28: 00000a59425185a911d033b1aa3c27e3d778659474aa8e687e65d71ca9b72e4e
```