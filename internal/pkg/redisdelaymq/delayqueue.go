package redisdelaymq

import (
	"context"
	"errors"
	"strconv"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

// DelayQueue 延迟队列.
type DelayQueue struct {
	client *gredis.Redis
	topic  string
	delay  time.Duration
}

// Options 可选配置项.
type Options struct {
	Group string
}

// Option 配置项.
type Option func(*Options)

// New 新建延迟队列实例.
func New(topic string, delay time.Duration, opts ...Option) *DelayQueue {
	options := Options{
		Group: "default",
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &DelayQueue{
		client: g.Redis(options.Group),
		topic:  topic,
		delay:  delay,
	}
}

// Produce 将一个或多个消息放到延迟队列.
func (d *DelayQueue) Produce(ctx context.Context, messages []interface{}) error {
	if len(messages) == 0 {
		return errors.New("messages cannot be empty")
	}

	// 将任务加入有序集合，使用当前时间加上延迟时间作为分数.
	score := float64(time.Now().Unix()) + d.delay.Seconds()

	var members []gredis.ZAddMember
	for _, message := range messages {
		member := gredis.ZAddMember{
			Score:  score,
			Member: message,
		}
		members = append(members, member)
	}

	_, err := d.client.ZAdd(ctx, d.topic, &gredis.ZAddOption{NX: true}, members[0], members[1:]...)

	return err
}

// Consume 获取score小于给定时间戳(一般为消费的当前时间戳)的所有消息.
// NOTE: 多实例同时消费时，需要使用分布式锁，避免重复消费.
func (d *DelayQueue) Consume(ctx context.Context, timestamp int64) (gvar.Vars, error) {
	results, err := d.client.ZRange(ctx, d.topic, 0, timestamp, gredis.ZRangeOption{
		ByScore:    true,
		WithScores: false,
	})
	if err == nil {
		return results, nil
	}

	// 一些redis版本不支持如上命令方式，通过ZRANGEBYSCORE兼容.
	result, err := d.client.Do(ctx, "ZRANGEBYSCORE", d.topic, 0, timestamp)
	if err != nil {
		return nil, err
	}

	return result.Vars(), nil
}

// Remove 删除score小于给定时间戳(一般为传给Consume方法的时间戳)的所有消息.
func (d *DelayQueue) Remove(ctx context.Context, timestamp int64) (int64, error) {
	min := "-inf"
	max := strconv.FormatInt(timestamp, 10)

	// n 为成功删除的消息数量
	n, err := d.client.ZRemRangeByScore(ctx, d.topic, min, max)

	return n, err
}

// Remove 删除指定消息.
func (d *DelayQueue) RemoveByMsg(ctx context.Context, message interface{}) (int64, error) {
	n, err := d.client.ZRem(ctx, d.topic, message)

	return n, err
}

// WithRedisGroup 设置使用配置的哪组redis服务器(或集群).
func WithRedisGroup(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}
