[![Build Status](https://travis-ci.org/deelawn/structsync.svg?branch=master)](https://travis-ci.org/deelawn/structsync)
[![Coverage Status](https://coveralls.io/repos/github/deelawn/structsync/badge.svg?branch=master)](https://coveralls.io/github/deelawn/structsync?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/deelawn/structsync)](https://goreportcard.com/report/github.com/deelawn/structsync)
[![License](https://img.shields.io/github/license/deelawn/structsync.svg)](https://github.com/deelawn/structsync/blob/master/LICENSE)

# structsync
Use field tags to sync one struct to another

Example:
```
package main

import (
	"fmt"

	"github.com/deelawn/structsync"
)

type srcStruct struct {
	S *string  `sync:"s"`
	I int      `sync:"i"`
	B bool     `sync:"b"`
	F *float64 `sync:"f"`
}

type dstStruct struct {
	AnInt   *int    `sync:"i"`
	AString string  `sync:"s"`
	AFloat  float64 `sync:"f"`
}

func main() {

	floatVal := float64(123.123)
	validDst := dstStruct{}
	validSrc := srcStruct{
		I: 19,
		B: true,
		F: &floatVal,
	}

	// structsync.Tag = "sync"
	if err := structsync.StructSync(validSrc, &validDst, structsync.Tag); err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("%+v\n", validDst)

}
```




