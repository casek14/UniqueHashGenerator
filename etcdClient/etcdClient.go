package etcdClient

import (
	"context"
	"go.etcd.io/etcd/clientv3"
	"log"
	"math/rand"
	"time"
)

type EtcdClient struct {
	Config clientv3.Config
	client *clientv3.Client
}

// Create etcd client from given settings, returns error when client creation fails
// etcdUrl string - Etcd database
//
//
//
func NewEtcdClient(etcdUrl string, etcdPort string) (*EtcdClient, error) {
	var etcdClient EtcdClient
	endpoint := etcdUrl + ":" + etcdPort
	conf := clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
		//Username:etcdUser,
		//Password:etcdpasswd,
	}
	etcdClient.Config = conf
	cli, err := clientv3.New(etcdClient.Config)
	if err != nil {
		return nil, err
	}

	etcdClient.client = cli
	return &etcdClient, nil
}

func (c *EtcdClient) CheckKeyExistence(hash string) bool {
	kv := clientv3.NewKV(c.client)
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	response, err := kv.Get(ctx, hash)
	cancel()
	if err != nil {
		log.Printf("Unable to get key %s. ERROR: %s\n", hash, err)
	}
	for _, ev := range response.Kvs {
		if string(ev.Key) == hash {
			return false
		}
	}
	return true
}

func createHash() string {
	rand.Seed(time.Now().UnixNano())
	n := rand.Intn(len(Symbols) - 1)
	char1 := Symbols[n]
	n = rand.Intn(len(Symbols) - 1)
	char2 := Symbols[n]
	n = rand.Intn(len(Symbols) - 1)
	char3 := Symbols[n]
	n = rand.Intn(len(Symbols) - 1)
	char4 := Symbols[n]
	n = rand.Intn(len(Symbols) - 1)
	char5 := Symbols[n]
	return char1 + char2 + char3 + char4 + char5
}

func (c *EtcdClient) CreateHash() string {
	var hash string

	hash = createHash()

	ok := c.CheckKeyExistence(hash)
	log.Println(ok)
	for {
		if ok {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			log.Printf("CREATING HASH IN DB: %s\n", hash)
			_, err := c.client.Put(ctx, hash, "used")
			if err != nil {
			}
			cancel()
			break
		} else {
			log.Printf("HASH %s, already exists\n", hash)
		}

	}

	return hash
}
