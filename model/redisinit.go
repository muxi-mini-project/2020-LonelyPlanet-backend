package model

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

//Redis
type Redis struct {
	Self redis.Conn
}

var RedisDb *Redis

func (rdb *Redis) Init() {
	newDb, err := redis.Dial(viper.GetString("redis.network"), viper.GetString("redis.addr"))
	if err != nil {
		fmt.Println(err)
	}
	RedisDb = &Redis{Self: newDb}
}

func (rdb *Redis) Close() error {
	if err := RedisDb.Self.Close(); err != nil {
		return err
	}
	return nil
}

