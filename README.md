## JSON config parser and encoder for Go

Installation:

```bash
go get github.com/deepch/config
```

### Examples

This package works `JSON`.

For the simplest example, consider some TOML file as just a list of keys
and values:

```json
{
  "Age": 25,
  "Cats": [
    "Cauchy",
    "Plato"
  ],
  "Pi": 3.14,
  "Perfection": [
    6,
    28,
    496,
    8128
  ],
  "DOB": "1987-07-05T05:45:00.000Z"
}
```

Which could be defined in Go as:

```go
type Config struct {
  Age int
  Cats []string
  Pi float64
  Perfection []int
  DOB time.Time // requires `import time`
}
```

And then decoded with:

```go
var conf Config
if _, err := config.Decode(jsonFileString, &conf); err != nil {
  // handle error
}
```

And then encode with:

```go
var conf Config
if _, err := config.Encode(jsonFileString, &conf); err != nil {
  // handle error
}
```

You can also use struct tags if your struct field name doesn't map to a JSON
key value directly:

```json
some_key_NAME = "wat"
```

```go
type json struct {
  ObscureKey string `json:"some_key_NAME"`
}
```


You can line and char error 
```bash
Decode Config Syntax Error File config/file.conf 
Line 5 Char 1 Error invalid character '1' looking for beginning of object key string
```
