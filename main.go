// Copyright 2014 The Go Authors.  All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Goobjdump dumps structural information about a Go object file.
// It does not include the actual data bytes of each symbol, under
// the assumption that using the compiler or assembler -S flag will
// take care of those bits.
package main // import "rsc.io/goobjdump"

import (
	"encoding/json"
	"log"
	"os"

	"rsc.io/goobjdump/goobj" // copied from cmd/internal/goobj
)

func main() {
	f, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	pkg, err := goobj.Parse(f, "main")
	if err != nil {
		log.Fatal(err)
	}
	js, err := json.MarshalIndent(pkg, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	os.Stdout.WriteString(string(js) + "\n")
}
