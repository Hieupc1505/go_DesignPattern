package main

import (
	"github.com/hieupc05/goDesignPattern/builder/exp2/internal"
)

func main() {
	service := internal.NewServiceBuilder()

	director := internal.NewServiceBuilderDirector()

	result := director.BuildService(service)

	result.DoService()

}
