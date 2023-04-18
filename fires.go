package main

import (
	"fmt"
	"github.com/go-fires/example/bootstrap"
	"github.com/go-fires/framework/facade"
)

func main() {
	app := bootstrap.App()
	defer app.Terminate()

	fmt.Println(facade.Hash().Make("123456"))

	app.Boot()
}
