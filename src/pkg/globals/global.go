package globals

import (
	"database/sql"
	EventBus2 "github.com/asaskevich/EventBus"
	"github.com/redis/go-redis/v9"
)

var (
	Db       *sql.DB
	Rdb      *redis.Client
	C        *Config
	EventBus EventBus2.Bus
)
