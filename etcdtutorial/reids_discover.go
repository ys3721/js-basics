package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var prefix = "gs:service"

func clientConnect() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	res, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
}

type ServiceInfo struct {
	Name string `json:"name"`
	Host string `json:"host"`
	Port int    `json:"port"`
}

func (s ServiceInfo) Key() string {
	return fmt.Sprintf("prefix:%s:%s:%d", s.Name, s.Host, s.Port)
}

func registerServiceByRedis(rdb *redis.Client, service ServiceInfo, ttl time.Duration) error {
	key := service.Key()
	data, _ := json.Marshal(service)

	return rdb.Set(ctx, key, data, ttl).Err()
}

func keepAliveServiceByRedis(rdb *redis.Client, service ServiceInfo, ttl time.Duration) error {
	key := fmt.Sprintf("service:%s:%s:%d", service.Name, service.Host, service.Port)
	data, _ := json.Marshal(service)

	return rdb.Set(ctx, key, data, ttl).Err()
}

func discoverServiceByRedis(rdb *redis.Client, name string) ([]ServiceInfo, error) {
	pattern := fmt.Sprintf("service:%s:*", name)
	serviceInfos := []ServiceInfo{}
	var cursor uint64
	for {
		keys, nextCursor, err := rdb.Scan(ctx, cursor, pattern, 2).Result()
		if err != nil {
			return nil, err
		}
		for _, key := range keys {
			data, err := rdb.Get(ctx, key).Result()
			if err != nil {
				log.Println(err)
				continue
			}
			var svc ServiceInfo
			err = json.Unmarshal([]byte(data), &svc)
			if err != nil {
				log.Println(err)
				continue
			}
			serviceInfos = append(serviceInfos, svc)
		}
		cursor = nextCursor
		if cursor == 0 {
			break
		}
	}
	return serviceInfos, nil
}

func callService(rdb *redis.Client, serviceName string) {
	serviceInfos, err := discoverServiceByRedis(rdb, serviceName)
	if err != nil {
		log.Fatal(err)
	}
	if len(serviceInfos) == 0 {
		log.Println("No service found")
		return
	}
	for _, service := range serviceInfos {
		fmt.Printf("Service Name: %s, Host: %s, Port: %d\n", service.Name, service.Host, service.Port)
	}
}

func doMainTest() {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	service := ServiceInfo{
		Name: "example_service",
		Host: "localhost",
		Port: 8080,
	}
	err := registerServiceByRedis(rdb, service, 10*time.Second)
	if err != nil {
		log.Fatal(err)
	}

	callService(rdb, "example_service")
}
