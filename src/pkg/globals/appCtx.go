package globals

import (
	"database/sql"
	EventBus2 "github.com/asaskevich/EventBus"
)

type AppCtx struct {
	db       *sql.DB
	config   *Config
	eventBus EventBus2.Bus
}

type Options func(a *AppCtx)

func NewAppCtx(op ...Options) *AppCtx {
	appCtx := new(AppCtx)
	for _, options := range op {
		options(appCtx)
	}
	return appCtx
}

func NewDefaultAppCtx() *AppCtx {
	return NewAppCtx(WithOptionDb(Db), WithOptionConfig(C), WithOptionEventBus(EventBus))
}

func (a *AppCtx) GetEventBus() EventBus2.Bus {
	return a.eventBus
}

func (a *AppCtx) GetDb() *sql.DB {
	return a.db
}

func WithOptionDb(db *sql.DB) Options {
	return func(a *AppCtx) {
		a.db = db
	}
}

func WithOptionConfig(c *Config) Options {
	return func(a *AppCtx) {
		a.config = c
	}
}

func WithOptionEventBus(e EventBus2.Bus) Options {
	return func(a *AppCtx) {
		a.eventBus = e
	}
}
