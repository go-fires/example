package main

import (
	"github.com/go-fires/example/app/console"
	"github.com/go-fires/example/bootstrap"
)

func main() {
	kernel := console.NewKernel(bootstrap.App())
	defer kernel.Terminate()

	kernel.Handle()
}
