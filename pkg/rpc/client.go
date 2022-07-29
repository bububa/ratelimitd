package rpc

import (
	"fmt"
	"time"

	"github.com/smallnest/rpcx/client"
	"github.com/smallnest/rpcx/protocol"
	"github.com/smallnest/rpcx/share"

	"github.com/bububa/ratelimitd/conf"
	"github.com/bububa/ratelimitd/pkg/logger"
	"github.com/bububa/ratelimitd/pkg/rpc/codec"
)

func NewProtobufClientPool(cfg conf.ClientConfig, poolSize int) (*client.XClientPool, error) {
	share.Codecs[protocol.SerializeType(4)] = &codec.Msgpack{}
	share.Codecs[protocol.SerializeType(5)] = &codec.Protobuf{}
	return newClientPool(cfg, poolSize, protocol.SerializeType(5))
}

func newClientPool(cfg conf.ClientConfig, poolSize int, serializeType protocol.SerializeType) (*client.XClientPool, error) {
	option := client.DefaultOption
	option.SerializeType = serializeType
	option.Heartbeat = true
	option.HeartbeatInterval = time.Second
	discover, err := client.NewPeer2PeerDiscovery(fmt.Sprintf("tcp@%s:%d", cfg.Host, cfg.Port), "")
	if err != nil {
		logger.Error().Err(err).Send()
		return nil, err
	}
	name := cfg.Name
	if name == "" {
		name = "ratelimitd"
	}
	return client.NewXClientPool(poolSize, name, client.Failtry, client.RandomSelect, discover, option), nil
}

func GetClientFromPool(pool *client.XClientPool) client.XClient {
	clt := pool.Get()
	return clt
}
