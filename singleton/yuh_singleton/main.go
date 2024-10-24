package main

import (
	"fmt"

	"github.com/hieupc05/goDesignPattern/singleton/yuh/singleton"
)

func main() {
	for i := 0; i < 10; i++ {
		go func() {
			fmt.Printf("%p\n", singleton.GetIntance())
		}()
	}
}
