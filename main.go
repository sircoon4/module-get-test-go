package main

import (
	"fmt"

	"github.com/sircoon4/bencodex-go"
	"github.com/sircoon4/module-publish-test-go/test"
	"github.com/sircoon4/module-publish-test-go/testother"
)

func main() {
	fmt.Println(test.Test())
	fmt.Println(testother.TestOther())

	mapData := map[string]any{"a": 1, "b": "2", "c": 3}
	encoded, err := bencodex.Encode(mapData)
	if err != nil {
		fmt.Println("error while encoding: ", err)
	}
	fmt.Println(encoded)
	decoded, err := bencodex.Decode(encoded)
	if err != nil {
		fmt.Println("error while decoding: ", err)
	}
	fmt.Printf("Value: %v Type: %T\n", decoded, decoded)
}
