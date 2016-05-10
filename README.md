[![wercker status](https://app.wercker.com/status/5f1378668f9aea20d76df49b71e19aa8/s/master "wercker status")](https://app.wercker.com/project/bykey/5f1378668f9aea20d76df49b71e19aa8)

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
