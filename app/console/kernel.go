package console

import (
	"github.com/go-fires/example/app/console/commands"
	"github.com/go-fires/framework/contracts/debug/recovery"
	"github.com/go-fires/framework/contracts/foundation"
)

type Kernel struct {
	app foundation.Application
}

var _ foundation.Kernel = (*Kernel)(nil)

func NewKernel(app foundation.Application) *Kernel {
	return &Kernel{
		app: app,
	}
}

func (k *Kernel) Bootstrap() {
	k.app.Boot()
}

func (k *Kernel) Handle() {
	defer func() {
		if err := recover(); err != nil {
			k.report(err)
		}
	}()

	k.Bootstrap()
	defer k.Terminate()

	commands.Execute()
}

func (k *Kernel) Terminate() {
	k.app.Terminate()
}

func (k *Kernel) report(err interface{}) {
	var handler recovery.Handler
	if k.app.Make("debug.recovery.handler", &handler) != nil {
		panic(err)
	}

	handler.Report(err)
}
