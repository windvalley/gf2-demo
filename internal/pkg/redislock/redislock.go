package redislock

import (
	"context"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	uuid "github.com/satori/go.uuid"
)

// Lock 分布式锁.
type Lock struct {
	client                 *gredis.Redis
	logger                 *glog.Logger
	key                    string
	value                  string
	expireSeconds          int64
	renewalIntervalSeconds int64
	timeoutSeconds         int64
	stopRenewalCh          chan struct{}
}

// Options 可选配置项.
type Options struct {
	LoggerName     string
	ExpireSeconds  int64
	RenewalSeconds int64
	TimeoutSeconds int64
}

// Option 配置项.
type Option func(*Options)

// New 创建一个新的分布式锁.
func New(lockKey string, opts ...Option) *Lock {
	options := Options{
		LoggerName:     "",
		ExpireSeconds:  60,
		RenewalSeconds: 20,
		TimeoutSeconds: 0,
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Lock{
		client:                 g.Redis(),
		key:                    lockKey,
		value:                  uuid.NewV4().String(),
		expireSeconds:          options.ExpireSeconds,
		renewalIntervalSeconds: options.RenewalSeconds,
		timeoutSeconds:         options.TimeoutSeconds,
		stopRenewalCh:          make(chan struct{}),
		logger:                 g.Log(options.LoggerName),
	}
}

// Acquire 获取锁.
func (l *Lock) Acquire(ctx context.Context) bool {
	ok, err := l.client.SetNX(ctx, l.key, l.value)
	if err != nil {
		l.logger.Errorf(ctx, "acquire lock failed, key: %s, value: %s, error: %v", l.key, l.value, err)
		return false
	}
	if !ok {
		return false
	}

	if _, err := l.client.Expire(ctx, l.key, l.expireSeconds); err != nil {
		l.logger.Warningf(ctx, "set lock expiration failed, key: %s, value: %s, error: %v", l.key, l.value, err)
	}

	// 成功获取到锁，启动续约 goroutine
	go l.startRenewal(ctx)

	return ok
}

// Release 释放锁.
func (l *Lock) Release(ctx context.Context) bool {
	// 关闭续约 goroutine
	close(l.stopRenewalCh)

	// 验证锁是否归属当前任务
	value, err := l.client.Get(ctx, l.key)
	if err != nil {
		//nolint: goconst
		l.logger.Errorf(ctx, "release lock failed, key: %s, value: %s, error: %v", l.key, l.value, err)
		return false
	}
	if value.String() != l.value {
		l.logger.Errorf(ctx, "release lock failed, key: %s, value: %s, error: lock is not owned by the current task", l.key, l.value)
		return false
	}

	// 使用 DEL 命令释放锁
	if _, err := l.client.Del(ctx, l.key); err != nil {
		l.logger.Errorf(ctx, "release lock failed, key: %s, value: %s, error: %v", l.key, l.value, err)
		return false
	}

	return true
}

// ForceRelease 不考虑锁的归属, 强制释放锁.
func (l *Lock) ForceRelease(ctx context.Context) bool {
	// 关闭续约 goroutine
	close(l.stopRenewalCh)

	// 使用 DEL 命令释放锁
	if _, err := l.client.Del(ctx, l.key); err != nil {
		l.logger.Errorf(ctx, "release lock failed, key: %s, value: %s, error: %v", l.key, l.value, err)
		return false
	}

	return true
}

// 续约锁.
func (l *Lock) startRenewal(ctx context.Context) {
	ticker := time.NewTicker(time.Duration(l.renewalIntervalSeconds) * time.Second)
	defer ticker.Stop()

	timeoutCh := make(chan struct{})
	if l.timeoutSeconds > 0 {
		go func() {
			time.Sleep(time.Duration(l.timeoutSeconds) * time.Second)
			close(timeoutCh)
		}()
	}

	for {
		select {
		// 续约，重置锁的过期时间
		case <-ticker.C:
			if _, err := l.client.Expire(ctx, l.key, l.expireSeconds); err != nil {
				l.logger.Warningf(ctx, "renew lock expiration failed, key: %s, value: %s, error: %v", l.key, l.value, err)
			}
		// 续约被停止，结束 goroutine
		case <-l.stopRenewalCh:
			return
		// 任务超时时间到, 结束 goroutine
		case <-timeoutCh:
			return
		}
	}
}

// WithLoggerName 使用指定的logger, 如果存在子项目logger, 则需要使用此函数提供指定的logger名称.
// 默认是gf框架配置文件中配置的default logger(apiserver).
func WithLoggerName(loggerName string) Option {
	return func(o *Options) {
		o.LoggerName = loggerName
	}
}

// WithExpireSeconds 设置锁的过期时间. 为分布式锁设置过期时间的意义在于防止程序崩溃或其他原因导致锁始终存在(死锁).
// 默认60秒.
func WithExpireSeconds(expireSeconds int64) Option {
	return func(o *Options) {
		o.ExpireSeconds = expireSeconds
	}
}

// WithRenewalSeconds 设置间隔多少秒对锁进行续约.
// 默认20秒.
func WithRenewalSeconds(renewalSeconds int64) Option {
	return func(o *Options) {
		o.RenewalSeconds = renewalSeconds
	}
}

// WithTimeoutSeconds 为续约goroutine设置超时时间, 一般是实际任务执行的超时时间.
// 默认不超时.
func WithTimeoutSeconds(timeoutSeconds int64) Option {
	return func(o *Options) {
		o.TimeoutSeconds = timeoutSeconds
	}
}
