# mtexport/parser

Golang Parser for [Movable Type Export Format](https://movabletype.org/documentation/appendices/import-export-format.html)

## USAGE

```go
import(
    "os"

    "github.com/soh335/mtexport/parser"
)

f, err := os.Open("path/to/file")
if err != nil {
    return err
}
defer f.Close()

stmts, err := parser.Parse(f, nil)
if err != nil {
    return err
}
```

## LICENSE

MIT
