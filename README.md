# expressions
An example named integer expression interpreter in Go using an ANTLR grammar.

This uses the Visitor pattern approach from https://github.com/cyberfox/antlr4

With that build of antlr4, you can:

```bash
java -jar $PATH_TO_ANTLR4JAR -Dlanguage=Go -no-listener -visitor Expressions.g4 -o parser
go test
```
