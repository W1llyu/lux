package xredis

import (
	"time"
	"github.com/garyburd/redigo/redis"
)

type Client struct {
	conn redis.Conn
}

func GetClient() *Client {
	return GetPool().Get()
}

func GetNamedClient(name string) *Client {
	return GetNamedPool(name).Get()
}

func (c *Client) Close() error {
	return c.conn.Close()
}

func (c *Client) Err() error {
	return c.conn.Err()
}

func (c *Client) Hkeys(key string) ([]string, error) {
	return redis.Strings(c.conn.Do("HKEYS", key))
}

func (c *Client) Hget(key, field string) (string, error) {
	return redis.String(c.conn.Do("HGET", key, field))
}

func (c *Client) Set(key, value string) error {
	_, err := c.conn.Do("SET", key, value)
	return err
}

func (c *Client) Get(key string) (string, error) {
	return redis.String(c.conn.Do("GET", key))
}

func (c *Client) Expire(key string, duration time.Duration) error {
	_, err := c.conn.Do("EXPIRE", key, duration.Seconds())
	return err
}