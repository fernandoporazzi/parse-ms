package main

import (
	"encoding/json"
	"fmt"

	parsems "github.com/fernandoporazzi/parse-ms"
)

func main() {
	// struct output
	fmt.Println(parsems.Parse(float64(60500.345678)))

	// json output
	js, _ := json.Marshal(parsems.Parse(float64(60500.345678)))
	fmt.Println(string(js))
}
