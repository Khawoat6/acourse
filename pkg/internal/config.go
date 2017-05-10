package internal

import (
	"time"

	"github.com/acoshift/configfile"
	"github.com/garyburd/redigo/redis"
)

var config = configfile.NewReader("config")

var (
	redisPrimaryAddr   = config.String("redis_primary_addr")
	redisPrimaryDB     = config.Int("redis_primary_db")
	redisPrimaryPass   = config.String("redis_primary_pass")
	redisSecondaryAddr = config.String("redis_secondary_addr")
	redisSecondaryDB   = config.Int("redis_secondary_db")
	redisSecondaryPass = config.String("redis_secondary_pass")
	xsrfSecret         = config.String("xsrf_secret")
)

var (
	primaryPool = &redis.Pool{
		IdleTimeout: 10 * time.Minute,
		MaxIdle:     10,
		MaxActive:   100,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisPrimaryAddr,
				redis.DialDatabase(redisPrimaryDB),
				redis.DialPassword(redisPrimaryPass),
			)
		},
	}
	secondaryPool = &redis.Pool{
		IdleTimeout: 10 * time.Minute,
		MaxIdle:     10,
		MaxActive:   100,
		Wait:        true,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", redisSecondaryAddr,
				redis.DialDatabase(redisSecondaryDB),
				redis.DialPassword(redisSecondaryPass),
			)
		},
	}
)

func init() {
	time.Local = time.UTC
}

// GetPrimaryDB returns primary redis connection from pool, use for store app data
func GetPrimaryDB() redis.Conn {
	return primaryPool.Get()
}

// GetSecondaryDB returns secondary redis connection from pool, use for store session
func GetSecondaryDB() redis.Conn {
	return secondaryPool.Get()
}

// GetSecondaryPool returns secondary redis pool
func GetSecondaryPool() *redis.Pool {
	return secondaryPool
}

// GetXSRFSecret returns xsrf secret
func GetXSRFSecret() string {
	return xsrfSecret
}
