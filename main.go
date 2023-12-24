package main

import (
	"bytes"
	"encoding/binary"
	"log"
)

func main() {
	key, b := make([]byte, 2), make([]byte, 2)

	binary.LittleEndian.PutUint16(b, 143)
	binary.LittleEndian.PutUint16(key, 43)

	log.Println(bytes.Compare(b, key))
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
