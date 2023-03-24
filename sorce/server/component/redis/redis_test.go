package redis

import (
	"testing"
	"time"
)

func TestRedis(t *testing.T) {
	instance.Set("key", "value", time.Second*3600)
}
