package main

import (
	"github.com/linj-disanbo/hello"
	helloV2 "github.com/linj-disanbo/hello/v2"
	helloV3 "github.com/linj-disanbo/hello/v3"
)

func main() {
	hello.Hello("v1")
	helloV2.Hello("v2")
	helloV3.Hello("v3", "v3")
}
