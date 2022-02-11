package main

import (
	"sample/redis"
	"sync"
)

func main(){
	//server := chttp.CreateServer("192.168.0.40", "8080")
	//server.Run()

	var wg sync.WaitGroup
	client := redis.NewClient()
	defer client.Close()


	wg.Add(1)
	go redis.Consumer(client, 1, &wg)
	go redis.Consumer(client, 2, &wg)
	go redis.Producer(client, &wg)
	//go redis.Consumer(client, 2, &wg)

	wg.Wait()
}

