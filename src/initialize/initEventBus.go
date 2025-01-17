package initialize

import (
	"github.com/asaskevich/EventBus"
)

func InitEventBus() EventBus.Bus {
	bus := EventBus.New()

	return bus
}
