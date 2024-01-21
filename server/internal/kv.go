package internal

import (
	"context"
	pb "github.com/wcygan/docker-to-kubernetes-example/generated/go/kv/v1"
)

type KeyValueService struct {
	pb.UnimplementedKeyValueServiceServer
	Redis *RedisClient
}

func NewKeyValueService(redis *RedisClient) *KeyValueService {
	return &KeyValueService{Redis: redis}
}

func (s *KeyValueService) Put(ctx context.Context, in *pb.PutRequest) (*pb.PutResponse, error) {
	err := s.Redis.Put(in.GetKey(), in.GetValue())
	if err != nil {
		return nil, err
	}
	return &pb.PutResponse{Success: true}, nil
}

func (s *KeyValueService) Get(ctx context.Context, in *pb.GetRequest) (*pb.GetResponse, error) {
	val, err := s.Redis.Get(in.GetKey())
	if err != nil {
		return nil, err
	}
	return &pb.GetResponse{Value: val, Found: true}, nil
}
