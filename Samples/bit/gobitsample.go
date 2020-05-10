package main

import (
	"fmt"
	"math/bits"
)

// https://qiita.com/cia_rana/items/2df8fb14517ab74af4c7

func main() {
	fmt.Println(bits.Len64(4097))
	fmt.Println(bits.RotateLeft(1, 10))
}
