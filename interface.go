package iredis

import(
	"github.com/garyburd/redigo/redis"
)

type Pool interface {
	Get() redis.Conn
}