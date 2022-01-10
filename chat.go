//usr/bin/go run $0 $@ ; exit
package main

import (
	"context"
	"flag"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

func main() {
	redisHost := flag.String("host", "localhost", "Redis host")
	redisPort := flag.String("port", "6379", "Redis port")
	password := flag.String("password", "my_password", "Redis password")
	channel := flag.String("channel", "linux_adm", "Channel to pub/sub")
	pub := flag.Bool("pub", false, "Publish message")
	flag.Parse()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     *redisHost + ":" + *redisPort,
		Password: *password,
	})

	if *pub {
		log.Println("Starting as publisher")
		for range time.Tick(5 * time.Second) {
			log.Println("Sending message to channel")
			redisClient.Publish(context.Background(), *channel, "Hello, world!")
		}
	} else {
		log.Println("Starting as subscriber")
		subsriber := redisClient.Subscribe(context.Background(), *channel)
		for {
			msg, err := subsriber.Receive(context.Background())
			if err != nil {
				panic(err)
			}
			log.Printf("> %s\n", msg)
		}
	}
}
