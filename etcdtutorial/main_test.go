package main

import (
	"log"
	"testing"
	"time"

	"github.com/redis/go-redis/v9"
)

func TestRegisterService(t *testing.T) {
	RegisterService([]string{"localhost:2379"}, "/services/hello/127.0.0.1:50051", "127.0.0.1:50051", 5)
}

func TestRedisClient(t *testing.T) {
	clientConnect()
}

func TestMain0(t *testing.T) {
	main0()
}

func TestRegisterServiceByReids(t *testing.T) {
	service := ServiceInfo{
		Name: "my-service",
		Host: "localhost",
		Port: 8080,
	}
	rbd := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	err := registerServiceByRedis(rbd, service, 10*time.Second)
	if err != nil {
		t.Error(err)
	}
	log.Println("服务注册成功")
}

func TestDoMainTest(t *testing.T) {
	doMainTest()
}
