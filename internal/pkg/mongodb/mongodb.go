package mongodb

import (
	"context"
	"errors"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/qiniu/qmgo"
)

var (
	client *qmgo.Client
	once   sync.Once
)

type Options struct {
	Address  string
	User     string
	Password string
}

type Option func(*Options)

// New init mongodb client.
func New(ctx context.Context, opts ...Option) (*qmgo.Client, error) {
	options := Options{}
	for _, opt := range opts {
		opt(&Options{})
	}

	if options.Address == "" {
		address, err := g.Cfg().Get(ctx, "mongodb.address")
		if err != nil {
			return nil, err
		}
		options.Address = address.String()
	}

	if options.User == "" {
		user, err := g.Cfg().Get(ctx, "mongodb.user")
		if err != nil {
			return nil, err
		}
		options.User = user.String()
	}
	if options.Password == "" {
		password, err := g.Cfg().Get(ctx, "mongodb.password")
		if err != nil {
			return nil, err
		}
		options.Password = password.String()
	}

	uri := "mongodb://" + options.Address
	if options.User != "" && options.Password != "" {
		uri = "mongodb://" + options.Address + ":" + options.User + "@" + options.Password
	}

	var err error

	once.Do(func() {
		client, err = qmgo.NewClient(
			ctx,
			&qmgo.Config{Uri: uri},
		)
	})

	if err != nil {
		return nil, err
	}
	if client == nil {
		return nil, errors.New("mongodb client is nil")
	}

	return client, nil
}

// GetClient get mongodb client.
func GetClient() *qmgo.Client {
	return client
}

// WithAddress set mongodb address.
func WithAddress(address string) Option {
	return func(o *Options) {
		o.Address = address
	}
}

// WithUser set mongodb user.
func WithUser(user string) Option {
	return func(o *Options) {
		o.User = user
	}
}

// WithPassword set mongodb password.
func WithPassword(password string) Option {
	return func(o *Options) {
		o.Password = password
	}
}
