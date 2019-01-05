package db

import (
	"github.com/go-redis/redis"
	"github.com/vmihailenco/msgpack"
)

type Db struct {
	client *redis.Client
}

func NewDb(address string, dbNumber int) (*Db, error) {
	db := &Db{
		client: redis.NewClient(&redis.Options{
			Addr: address,
			DB:   dbNumber,
		}),
	}

	_, err := db.client.Ping().Result()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func (d *Db) Get(key string, out interface{}) (interface{}, error) {
	val, err := d.client.Get(key).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil || val == "" {
		return nil, err
	}

	err = msgpack.Unmarshal([]byte(val), out)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (d *Db) Set(key string, value interface{}) (interface{}, error) {
	b, err := msgpack.Marshal(value)
	if err != nil {
		return nil, err
	}
	return d.client.Set(key, b, 0).Result()
}

func (d *Db) SAdd(key string, values ...interface{}) error {
	_, err := d.client.SAdd(key, values...).Result()
	return err
}

func (d *Db) SMembers(key string) ([]string, error) {
	return d.client.SMembers(key).Result()
}
