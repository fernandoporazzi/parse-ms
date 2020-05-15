# parse-ms
Parse milliseconds into a struct

![](https://github.com/fernandoporazzi/parse-ms/workflows/Go/badge.svg)


## Install

```
$ go get -u github.com/fernandoporazzi/parse-ms
```

## API

**type** `ParsedMilliseconds`
```go
type ParsedMilliseconds struct {
	Days         int `json:"days"`
	Hours        int `json:"hours"`
	Minutes      int `json:"minutes"`
	Seconds      int `json:"seconds"`
	Milliseconds int `json:"milliseconds"`
	Microseconds int `json:"microseconds"`
	Nanoseconds  int `json:"nanoseconds"`
}
```

**func** `Parse`
```go
func Parse(m float64) ParsedMilliseconds
```
Parse returns a struct of type `ParsedMilliseconds` based on the given `float64` input.

## Usage

```go
package main

import (
	"encoding/json"
	"fmt"

	parsems "github.com/fernandoporazzi/parse-ms"
)

func main() {
	// struct output
	fmt.Println(parsems.Parse(60500.345678))

	// json output
	js, _ := json.Marshal(parsems.Parse(60500.345678))
	fmt.Println(string(js))
}
```

## License

[MIT License](https://github.com/fernandoporazzi/parse-ms/blob/master/LICENSE)
