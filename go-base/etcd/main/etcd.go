package main

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"log"
	"time"
)

func main() {
	// 设置etcd服务器的地址和端口
	etcdConfig := clientv3.Config{
		Endpoints:   []string{"124.70.197.10:2379"},
		DialTimeout: 5 * time.Second,
	}

	// 创建etcd客户端
	client, err := clientv3.New(etcdConfig)
	if err != nil {
		log.Fatalf("Failed to create etcd client: %v", err)
	}
	defer client.Close()

	// 测试etcd连接
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = client.Get(ctx, "test_key")
	if err != nil {
		log.Fatalf("Failed to connect to etcd: %v", err)
	} else {
		fmt.Println("Connected to etcd successfully!")
	}
}
