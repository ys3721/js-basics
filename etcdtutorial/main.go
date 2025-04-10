package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

func main() {
	cltCfg := clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	}
	cli, err := clientv3.New(cltCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	key := "/services/my-service/instance-1"
	value := `{"ip":"192.168.4.54", "port":8080}`

	resp, err := cli.Put(context.Background(), key, value)
	if err != nil {
		log.Fatal(err)
	}

	getResp, err := cli.Get(context.Background(), "/services/my-service/", clientv3.WithPrefix())
	if err != nil {
		log.Fatal(err)
	}
	for _, kv := range getResp.Kvs {
		fmt.Println("key:", string(kv.Key), "value:", string(kv.Value))
	}
	fmt.Println("put result:", resp)

	watchChan := cli.Watch(context.Background(), key, clientv3.WithPrefix())
	fmt.Println("Watch for service changes...")
	for watchResp := range watchChan {
		for _, event := range watchResp.Events {
			fmt.Println("event:", event.Type, "key:", string(event.Kv.Key), "value:", string(event.Kv.Value))
			switch event.Type {
			case mvccpb.PUT:
				fmt.Println("put key:", string(event.Kv.Key), "value:", string(event.Kv.Value))
			case mvccpb.DELETE:
				fmt.Println("delete key:", string(event.Kv.Key), "value:", string(event.Kv.Value))
			}
		}
	}

}

func RegisterService(etchEndpoints []string, serviceKey string, serviceValue string, ttl int64) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   etchEndpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return err
	}
	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}
	ch, err := cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return err
	}
	_, err = cli.Put(context.Background(), serviceKey, serviceValue, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	go func() {
		for {
			ka := <-ch
			if ka == nil {
				fmt.Println("keep alive channel closed")
				break
			}
			fmt.Println("keep alive:", ka)
		}
	}()
	return nil
}
