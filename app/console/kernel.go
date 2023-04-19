package console

import (
	"github.com/go-fires/example/app/console/commands"
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

func (k Kernel) Bootstrap() {
	k.app.Boot()
}

func (k Kernel) Handle() {
	k.Bootstrap()

	commands.Execute()
}

func (k Kernel) Terminate() {
	k.app.Terminate()
}
