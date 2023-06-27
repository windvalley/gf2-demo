package kafka

import (
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	kafka "github.com/segmentio/kafka-go"
)

var (
	reader *kafka.Reader
	writer *kafka.Writer

	once sync.Once
)

type Options struct {
	Brokers []string
	Topic   string
	GroupID string
}

type Option func(*Options)

// NewReader new kafka reader.
func NewReader(ctx context.Context, opts ...Option) (*kafka.Reader, error) {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.Brokers == nil {
		brokers, err := g.Cfg().Get(ctx, "kafka.brokers")
		if err != nil {
			return nil, err
		}
		options.Brokers = brokers.Strings()
	}
	if options.Topic == "" {
		topic, err := g.Cfg().Get(ctx, "kafka.topic")
		if err != nil {
			return nil, err
		}
		options.Topic = topic.String()
	}
	if options.GroupID == "" {
		groupID, err := g.Cfg().Get(ctx, "kafka.group_id")
		if err != nil {
			return nil, err
		}
		options.GroupID = groupID.String()
	}

	once.Do(func() {
		reader = kafka.NewReader(kafka.ReaderConfig{
			Brokers:     options.Brokers,
			GroupID:     options.GroupID,
			Topic:       options.Topic,
			StartOffset: kafka.FirstOffset,
		})
	})

	if reader == nil {
		return nil, errors.New("kafka reader is nil")
	}

	return reader, nil
}

// NewWriter new kafka writer.
func NewWriter(ctx context.Context, opts ...Option) (*kafka.Writer, error) {
	options := Options{}
	for _, opt := range opts {
		opt(&options)
	}

	if options.Brokers == nil {
		brokers, err := g.Cfg().Get(ctx, "kafka.brokers")
		if err != nil {
			return nil, err
		}
		options.Brokers = brokers.Strings()
	}
	if options.Topic == "" {
		topic, err := g.Cfg().Get(ctx, "kafka.topic")
		if err != nil {
			return nil, err
		}
		options.Topic = topic.String()
	}

	once.Do(func() {
		writer = kafka.NewWriter(kafka.WriterConfig{
			Brokers: options.Brokers,
			Topic:   options.Topic,
		})
	})

	if writer == nil {
		return nil, errors.New("kafka writer is nil")
	}

	return writer, nil
}

// GetReader get kafka reader.
func GetReader() *kafka.Reader {
	return reader
}

// GetWriter get kafka writer.
func GetWriter() *kafka.Writer {
	return writer
}

// WithBrokers set kafka brokers.
func WithBrokers(brokers []string) Option {
	return func(o *Options) {
		o.Brokers = brokers
	}
}

// WithTopic set kafka topic.
func WithTopic(topic string) Option {
	return func(o *Options) {
		o.Topic = topic
	}
}

// WithGroupID set kafka group id.
func WithGroupID(groupID string) Option {
	return func(o *Options) {
		o.GroupID = groupID
	}
}
