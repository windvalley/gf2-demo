package rediscache

import (
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcache"
)

// Options 可选配置项.
type Options struct {
	Group string
}

// Option 配置项.
type Option func(*Options)

func New(opts ...Option) *gcache.Cache {
	options := Options{
		Group: "default",
	}

	for _, opt := range opts {
		opt(&options)
	}

	cache := gcache.New()
	redis := g.Redis(options.Group)

	cache.SetAdapter(gcache.NewAdapterRedis(redis))

	return cache
}

// WithRedisGroup 设置使用配置的哪组redis服务器(或集群).
func WithRedisGroup(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}
