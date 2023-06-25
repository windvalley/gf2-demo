package redismq

import (
	"context"
	"errors"
	"fmt"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
)

type Queue struct {
	client *gredis.Redis
	topic  string
}

type Options struct {
	Group string
}

type Option func(*Options)

func New(topic string, opts ...Option) *Queue {
	options := Options{
		Group: "default",
	}

	for _, opt := range opts {
		opt(&options)
	}

	return &Queue{
		client: g.Redis(options.Group),
		topic:  topic,
	}
}

// Produce 生产消息到队列.
func (q *Queue) Produce(ctx context.Context, messages []string) error {
	if len(messages) == 0 {
		return errors.New("messages cannot be empty")
	}

	_, err := q.client.RPush(ctx, q.topic, messages)
	if err != nil {
		return err
	}

	return nil
}

// Consume 消费队列中的信息.
func (q *Queue) Consume(ctx context.Context) (string, error) {
	results, err := q.client.BLPop(ctx, 0, q.topic)
	if err != nil {
		return "", err
	}

	//nolint: gomnd
	if len(results) < 2 {
		return "", fmt.Errorf("invalid result")
	}

	message := results[1]

	return message.String(), nil
}

func WithRedisGroup(group string) Option {
	return func(o *Options) {
		o.Group = group
	}
}
