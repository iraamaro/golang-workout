---
tags: ['golang', 'studies', 'programming']
date: '2025-02-02'
title: 'the-go-programing-language'
author: ['DONOVAN, Alan A.A. KERNIGHAN, Brian W.'] 
---

# 01-tutorial

Use of `asdf` with `go1.8.7`

Instalation, attr, set global 1.8.7, config GOPATH

```sh
export GOPATH=$HOME/Documents/golang
```

BEST setup [here](https://archive.is/5EF5m).

Use of `go get`

```sh
go get gopl.io/ch1/echo1

$GOPATH/bin/echo1

```
**compile**

`go build -v -work -x *.go`

**compile and execute and gen temp binary**

`go help run`

*By default, 'go run' runs the compiled binary directly:
'a.out arguments...'.
If the -exec flag is given, 'go run' invokes the binary using xprog:
'xprog a.out arguments...'.
If the -exec flag is not given, GOOS or GOARCH is different from the
system default, and a program named go_$GOOS_$GOARCH_exec can be found
on the current search path, 'go run' invokes the binary using that program,
for example 'go_js_wasm_exec a.out arguments...'. This allows execution of
cross-compiled programs when a simulator or other execution method is
available.*

`go run -v 8.go`

**review of slices like `Args[:]`**

Benchmark in dir *08* using `go test =v bench=.`

