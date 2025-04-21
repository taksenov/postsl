/*
Copyright © 2025 taksenov@gmail.com
*/

// Package customerrors -- error handlers.
package customerrors

import (
	"bytes"
	"log"
)

func HandleErr(err error, fn string) {
	var buf bytes.Buffer
	logger := log.New(&buf, "[ERROR]: ", log.Lmsgprefix)

	logger.Print("func: ", fn, "; message: ", err)
	log.Fatal(&buf)
}
