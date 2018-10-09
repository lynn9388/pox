# PoX

[![Build Status](https://travis-ci.com/lynn9388/pox.svg?branch=master)](https://travis-ci.com/lynn9388/pox)

Some simple PoX implementations.

## PoW

[![GoDoc](https://godoc.org/github.com/lynn9388/pox/pow?status.svg)](https://godoc.org/github.com/lynn9388/pox/pow)

### Example

```go
data := []byte("lynn9388")
for difficulty := 1; difficulty < 6; difficulty++ {
    nonce := GetNonce(data, uint(difficulty))
    fmt.Println(Hash(data, nonce))
}
```

Output:

```text
0c9c05b5cabd4e096b03c1c34a43506ed14d587073b413038bb847182d48cdd3
00acbf00115bbac7df7101a4c85e05e4876a66704455ca03bd80f033d77a9236
000d528129f382d45e19bdfb7a1d28435545beda56980616104c19bcac0acc95
00009762cf50633d878ebee17dd87123efcee3b23c6a6e42cb1d25252584ecf1
0000063b1492354413af28fd46f54c74a15520981f2c20b57c8e062b2ef7f3c8
```