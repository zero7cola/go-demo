package bootstrap

import (
	"fmt"
	"go1/pkg/redis"
)

func SetupRedis() {
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", "127.0.0.1", "6379"),
		"",
		"",
		1,
	)
}
