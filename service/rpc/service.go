package rpc

import (
	"context"
	"os"
	"path/filepath"
	"sync"

	"github.com/golang/protobuf/proto"

	"github.com/bububa/ratelimitd/pb"
)

type Service struct {
	lock *sync.RWMutex
	mp   map[string]*Limiter
	db   string
}

func NewService(ctx context.Context, storagePath string) (*Service, error) {
	s := &Service{
		lock: new(sync.RWMutex),
		mp:   make(map[string]*Limiter),
		db:   filepath.Join(storagePath, "limiters.db"),
	}
	if err := s.load(ctx); err != nil {
		return nil, err
	}
	return s, nil
}

func (s *Service) NewLimiter(ctx context.Context, req *pb.Limiter, ret *pb.Limiter) error {
	s.lock.Lock()
	s.mp[req.GetName()] = NewLimiter(req)
	s.lock.Unlock()
	return nil
}

func (s *Service) RemoveLimiter(ctx context.Context, req *pb.Limiter, ret *pb.Limiter) error {
	s.lock.Lock()
	delete(s.mp, req.GetName())
	s.lock.Unlock()
	return nil
}

func (s *Service) Take(ctx context.Context, req *pb.Limiter, ret *pb.Limiter) error {
	s.lock.RLock()
	limiter, found := s.mp[req.GetName()]
	s.lock.RUnlock()
	if found {
		limiter.Take()
	}
	return nil
}

func (s *Service) List(ctx context.Context, req *pb.Limiter, ret *pb.LimiterList) error {
	s.lock.RLock()
	for _, limiter := range s.mp {
		ret.List = append(ret.List, limiter.Config())
	}
	s.lock.RUnlock()
	return nil
}

func (s *Service) Flush(ctx context.Context, req *pb.Limiter, ret *pb.Limiter) error {
	list := new(pb.LimiterList)
	if err := s.List(ctx, req, list); err != nil {
		return err
	}
	bs, _ := proto.Marshal(list)
	f, err := os.OpenFile(s.db, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer f.Close()
	_, err = f.Write(bs)
	return err
}

func (s *Service) load(ctx context.Context) error {
	bs, err := os.ReadFile(s.db)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	limiters := new(pb.LimiterList)
	if err := proto.Unmarshal(bs, limiters); err != nil {
		return err
	}
	s.lock.Lock()
	defer s.lock.Unlock()
	for _, cfg := range limiters.GetList() {
		s.mp[cfg.GetName()] = NewLimiter(cfg)
	}
	return nil
}
