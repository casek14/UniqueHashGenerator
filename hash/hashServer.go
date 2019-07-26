package hash

import (
	"context"
	"github.com/casek14/UniqueHashGenerator/etcdClient"
	"log"
)

type HashServerImpl struct {
	Client *etcdClient.EtcdClient
}

func NewHasServer(dbUrl string, dbPort string) *HashServerImpl {
	cli, err := etcdClient.NewEtcdClient(dbUrl, dbPort)
	if err != nil {
		log.Fatalf("Cannot create client: %s", err)
	}
	return &HashServerImpl{Client: cli}
}

func (s *HashServerImpl) GetHash(ctx context.Context, r *HashRequest) (*HashResponse, error) {
	hash := s.Client.CreateHash()
	return &HashResponse{Hash: hash}, nil
}
