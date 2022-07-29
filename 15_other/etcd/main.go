package main

import (
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"golang.org/x/net/context"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	// 创建连接
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"47.108.93.159:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return
	}
	defer func(client *clientv3.Client) {
		_ = client.Close()
	}(cli)

	// 监听
	ch := make(chan int) // 同步监听的创建
	go func(ch chan int) {
		watchChan := cli.Watch(context.TODO(), "/message")
		ch <- 1
		for res := range watchChan {
			for _, event := range res.Events {
				fmt.Printf("-WATCH- Type: %s , Key: %s , Value: %s \n", event.Type, event.Kv.Key, event.Kv.Value)
			}
		}
	}(ch)
	<-ch

	// 写
	_, err = cli.Put(context.TODO(), "/message", "Hello")
	if err != nil {
		return
	}

	// 读
	get, err := cli.Get(context.TODO(), "/", clientv3.WithPrefix()) // 查看以key为前缀的
	if err != nil {
		return
	}
	for _, kv := range get.Kvs {
		fmt.Printf("-GET- %s : %s \n", kv.Key, kv.Value)
	}

	// 删除
	_, err = cli.Delete(context.TODO(), "/message")
	if err != nil {
		return
	}

	// 租约
	fmt.Println("=================================")
	// 注册,ttl指的是最少多久过期,不是实际值
	lease, err := cli.Grant(context.TODO(), 1)
	if err != nil {
		return
	}
	// 自动续约
	_, err = cli.KeepAlive(context.TODO(), lease.ID)
	if err != nil {
		return
	}
	// 使用携带租约的 KV
	_, err = cli.Put(context.TODO(), "lease_test1", "hi", clientv3.WithLease(lease.ID))
	if err != nil {
		return
	}
	// 查看
	fmt.Println("解除前:")
	leaseRes, err := cli.Get(context.TODO(), "lease_test1")
	if err != nil {
		return
	}
	for _, kv := range leaseRes.Kvs {
		fmt.Printf("lease:%v,key:%s,v:%s\n", kv.Lease, kv.Key, kv.Value)
	}

	// 查看使用了多久，哪些key使用了该租约
	fmt.Println("哪些key使用了该租约:")
	live, err := cli.TimeToLive(context.TODO(), lease.ID, clientv3.WithAttachedKeys())
	if err != nil {
		return
	}
	fmt.Println(live.GrantedTTL)
	fmt.Println(live.TTL)
	for _, key := range live.Keys {
		fmt.Println("key:", string(key))
	}

	// 解除
	_, err = cli.Revoke(context.TODO(), lease.ID)
	if err != nil {
		return
	}

	// 过期后再查看
	fmt.Println("解除后:")
	leaseRes, err = cli.Get(context.TODO(), "lease_test1")
	if err != nil {
		return
	}
	for _, kv := range leaseRes.Kvs {
		fmt.Printf("lease:%v,key:%s,v:%s\n", kv.Lease, kv.Key, kv.Value)
	}

	time.Sleep(time.Second)
}
