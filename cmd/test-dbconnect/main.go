package main

import (
	"fmt"
	"time"
)

// Schemes: http, https
// Host: localhost
// BasePath: /v1
// Version: 0.0.1
// License: MIT http://opensource.org/licenses/MIT
//
// Consumes:
// - application/json
// - application/xml
//
// Produces:
// - application/json
// - application/xml
//
//
// swagger:meta

func main() {
	xx := time.Now()
	fmt.Println(xx.String())
}
