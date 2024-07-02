package main

import (
	"fmt"

	"github.com/sircoon4/module-publish-test-go/test"
	"github.com/sircoon4/module-publish-test-go/testother"
)

func main() {
	fmt.Println(test.Test())
	fmt.Println(testother.TestOther())
}
