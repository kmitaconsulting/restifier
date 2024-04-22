package cache

import (
	"kmitaconsulting/restifier/data"

	"github.com/go-redis/redis"
)

type Cache struct {
	redisClient *redis.Client
}

func NewCache() *Cache {
	client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })

    // Ping Redis to check if the connection is working
    _, err := client.Ping().Result()
    if err != nil {
        panic(err)
    }

	return &Cache{redisClient: client}
}

func (c *Cache) InsertRequest(ID string) {
	c.redisClient.HSet(ID, "data", "")
	c.redisClient.HSet(ID, "status", false)
}

func (c *Cache) InsertResponse(response *data.SomeObject) {
	c.redisClient.HSet(response.ID, "data", response.SomeData)
	c.redisClient.HSet(response.ID, "status", response.Status)
}

func (c *Cache) DeleteResponse(ID string) {
	c.redisClient.Del(ID)
}