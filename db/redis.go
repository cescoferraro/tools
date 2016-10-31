package db

import (
	"github.com/spf13/viper"
	"log"
	"time"
	"gopkg.in/redis.v5"
)

//MONGO local MongoObject
var RedisServer RedisStore

type RedisStore struct {
	Name   string
	Client *redis.Client
}

//InitMongo set MongoDB connection
func (store RedisStore) Run(done chan bool) {
	projectName := viper.GetString("project")
	url := viper.GetString("redis-host") + ":" + viper.GetString("redis-port")
	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: "descriptor8", // no password set
		DB:       0,             // use default DB
	})

	_, err := client.Subscribe(projectName + "/" + store.Name)
	for err != nil {
		time.Sleep(10 * time.Second)
		log.Println(err.Error())
		_, err = client.Subscribe(projectName + "/" + store.Name)
	}
	RedisServer.Client = client
	done <- true

}
