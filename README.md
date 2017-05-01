# expressions
An example named integer expression interpreter in Go using an ANTLR grammar.



```bash
mkdir testdir
cd testdir
export GOPATH=`pwd`
go get -v 	github.com/millergarym/expressions
cd src/github.com/millergarym/expressions
go generate gen.go
go test -v
```
