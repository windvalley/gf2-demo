package rediscache

import (
	"fmt"
	"time"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
)

func ExampleRedisCache() {
	// *** 此处配置仅测试使用, 正式使用请在manifest/config/config.yml中配置
	redisConfig := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	gredis.SetConfig(&redisConfig)
	// ****************************************

	var (
		err        error
		cacheKey   = "test"
		cacheValue = "value"
		ctx        = gctx.New()
	)

	// 创建redis cache实例.
	cache := New()

	// 写缓存.
	err = cache.Set(ctx, cacheKey, cacheValue, 5*time.Second)
	if err != nil {
		panic(err)
	}

	// 读取指定缓存key的缓存时间.
	res := cache.MustGetExpire(ctx, cacheKey)
	fmt.Println(res.String())

	// 使用cache实例读取指定缓存key的值.
	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// 使用redis实例读取指定缓存key的值, 用于对比是否和cache实例一致.
	fmt.Println(g.Redis().MustDo(ctx, "GET", cacheKey).String())

	// 列出redis缓存中的所有key.
	fmt.Println(cache.MustKeys(ctx))
	// 列出redis缓存中的所有键的值.
	fmt.Println(cache.MustValues(ctx))

	// 判断redis缓存中是否存在指定的key.
	fmt.Println(cache.MustContains(ctx, cacheKey))

	// 删除指定key.
	if _, err := cache.Remove(ctx, cacheKey); err != nil {
		panic(err)
	}

	fmt.Println(cache.MustGet(ctx, cacheKey).String())

	// Output:
	// 5s
	// value
	// value
	// [test]
	// [value]
	// true
}
