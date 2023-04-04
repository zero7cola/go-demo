package main

import (
	"fmt"
	"go1/bootstrap"
	"go1/pkg/cahce"
)

func main() {

	bootstrap.SetupRedis()

	if cahce.NewCache().SetCache("name", "hello 222") {
		name := cahce.NewCache().GetCache("name", false)

		fmt.Println(name)
	}

}
