package goque

import (
	"encoding"
	"github.com/garyburd/redigo/redis"
)

const (
	name = "goque"
)

type Queue struct {
	conn redis.Conn
}

func New(c Config) *Queue {
	conn, err := redis.Dial("tcp", ":6379")
	if err != nil {
		return nil
	}
	q := new(Queue)
	q.conn = conn
	return new(Queue)
}

func (q *Queue) Push(e encoding.TextMarshaler) (err error) {
	encodedObject, err := e.MarshalText()
	if err != nil {
		return
	}
	_, err = q.conn.Do("RPUSH", "queue:" + name, encodedObject)
	return
}
