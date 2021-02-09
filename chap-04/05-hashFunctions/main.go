package main

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

const (
	example1 = "first example"
	example2 = "second example"
)

func main() {
	firstHash := sha256.New()
	firstHash.Write([]byte(example1))

	res1 := firstHash.Sum(nil)
	fmt.Printf("%x\n", res1)

	secondHash := sha256.New()
	secondHash.Write([]byte(example2))

	res2 := secondHash.Sum(nil)
	fmt.Printf("%x\n", res2)

	fmt.Println("Are two hash equal?")
	fmt.Println(bytes.Equal(res1, res2))
}
