package mredis

import (
	"fmt"

	"github.com/go-redis/redis"
	"github.com/kailashjanakiraman/go-cloud/serviceboiler/constants"
)

func NewRedisClient() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     constants.RedisAddress,
		Password: constants.RedisPassword, // no password set
		DB:       constants.RedisDBIndex,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	return client

}

// func main() {

// 	client := NewRedisClient()
// 	err := client.Set("key", "value", 0).Err()
// 	if err != nil {
// 		panic(err)
// 	}

// 	val, err := client.Get("key").Result()
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println("key", val)

// 	val2, err := client.Get("key2").Result()
// 	if err == redis.Nil {
// 		fmt.Println("key2 does not exist")
// 	} else if err != nil {
// 		panic(err)
// 	} else {
// 		fmt.Println("key2", val2)
// 	}
// 	// Output: key value
// 	// key2 does not exist
// }
