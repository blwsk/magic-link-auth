package data

import (
  "os"
  "fmt"
  "gopkg.in/redis.v5"
)

var RedisClient *redis.Client

func RedisClientInit() {
  redisAddress := os.Getenv("REDIS_ADDRESS")

  fmt.Println(redisAddress)

  RedisClient := redis.NewClient(&redis.Options{
    Addr: redisAddress,
    Password: "",
    DB: 0,
  })

  pong, err := RedisClient.Ping().Result()
  fmt.Println(pong, err)
}
