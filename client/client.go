package client

import (
	"context"
	"time"

	"github.com/smallnest/rpcx/client"

	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pb"
	"github.com/bububa/ratelimitd/pkg/rpc"
)

type Client struct {
	pool *client.XClientPool
}

func NewClient(cfg conf.ClientConfig, poolSize int) (*Client, error) {
	pool, err := rpc.NewProtobufClientPool(cfg, poolSize)
	if err != nil {
		return nil, err
	}
	return &Client{
		pool: pool,
	}, nil
}

func (c *Client) Close() {
	c.pool.Close()
}

func (c *Client) NewLimiter(ctx context.Context, name string, rate int, interval time.Duration) error {
	clt := rpc.GetClientFromPool(c.pool)
	req := new(pb.Limiter)
	req.Name = name
	req.Rate = int32(rate)
	req.Interval = int64(interval)
	return clt.Call(ctx, "NewLimiter", req, nil)
}

func (c *Client) RemoveLimiter(ctx context.Context, name string) error {
	clt := rpc.GetClientFromPool(c.pool)
	req := new(pb.Limiter)
	req.Name = name
	return clt.Call(ctx, "RemoveLimiter", req, nil)
}

func (c *Client) Take(ctx context.Context, name string) (time.Duration, error) {
	clt := rpc.GetClientFromPool(c.pool)
	req := new(pb.Limiter)
	req.Name = name
	ret := new(pb.Limiter)
	err := clt.Call(ctx, "Take", req, ret)
	return time.Duration(ret.Interval), err
}

func (c *Client) List(ctx context.Context, ret *pb.LimiterList) error {
	clt := rpc.GetClientFromPool(c.pool)
	return clt.Call(ctx, "List", new(pb.Limiter), ret)
}
