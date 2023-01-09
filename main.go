package main

import (
	"encoding/binary"
	"fmt"
)

func main() {

	b := make([]byte, 2)

	binary.LittleEndian.PutUint16(b, 12311)
	fmt.Printf("%s", b)

}
