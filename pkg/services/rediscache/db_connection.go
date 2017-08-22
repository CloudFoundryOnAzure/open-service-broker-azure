package rediscache

import (
	"fmt"

	"github.com/go-redis/redis"
)

func getDBConnection(pc *redisProvisioningContext) error {

	client := redis.NewClient(&redis.Options{
		Addr:     pc.ServerName,
		Password: pc.PrimaryKey, // primary key generated by Azure
		DB:       0,             // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		return fmt.Errorf("error connecting to the database: %s", err)
	}

	return err
}