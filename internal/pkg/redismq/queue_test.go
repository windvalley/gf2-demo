//go:build ignore
// +build ignore

package redismq

import (
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	uuid "github.com/satori/go.uuid"
)

type Message struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// 使用流程示例: go test -v
func TestQueue(t *testing.T) {
	ctx := gctx.New()

	// *** 此处配置仅测试使用, 正式使用请在manifest/config/config.yml中配置
	redisConfig := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	redisGroup := "queue" // 默认: default
	gredis.SetConfig(&redisConfig, redisGroup)
	// ****************************************

	topic := "queue"

	// 如果是默认的default分组, 可忽略此选项.
	groupOption := WithRedisGroup(redisGroup)

	q := New(topic, groupOption)

	msg1 := Message{
		ID:   uuid.NewV4().String(),
		Name: "hello",
	}
	msg2 := Message{
		ID:   uuid.NewV4().String(),
		Name: "world",
	}

	var msgs []string
	msg1Json := gconv.String(msg1)
	msg2Json := gconv.String(msg2)
	msgs = append(msgs, msg1Json, msg2Json)

	// 生产者
	err := q.Produce(ctx, msgs)
	if err != nil {
		t.Errorf("[PRODUCER] produce msgs failed: %v", err)
		return
	}
	t.Logf("[PRODUCER] produce msgs: %#v", msgs)

	err = q.Produce(ctx, []string{msg1Json})
	if err != nil {
		t.Errorf("[PRODUCER] produce msgs failed: %v", err)
		return
	}
	t.Logf("[PRODUCER] produce msg: %#v", msg1Json)

	go func() {
		time.Sleep(1 * time.Second)
		err = q.Produce(ctx, []string{msg2Json})
		if err != nil {
			t.Errorf("[PRODUCER] produce msgs failed: %v", err)
			return
		}
		t.Logf("[PRODUCER] produce msg: %#v", msg2Json)
	}()

	// 消费者, 并发5个消费端.
	for i := 0; i < 5; i++ {
		go func(n int) {
			for {
				res, err := q.Consume(ctx)
				if err != nil {
					t.Errorf("[CONSUMER-%d] get msg failed: %v", n, err)
					return
				}

				var msgs []string

				err = gconv.Structs(res, &msgs)
				if err != nil {
					t.Errorf("[CONSUMER-%d] get msgs failed, msgs: %s, error: %v", n, res, err)
					return
				}

				t.Logf("[CONSUMER-%d] got msgs: %#v", n, msgs)

				for _, msg := range msgs {
					var m Message
					err = gconv.Struct(msg, &m)
					if err != nil {
						t.Errorf("[CONSUMER] deserialize msg failed, msg: %s, error: %v", msg, err)
						continue
					}

					t.Logf("[CONSUMER] deserialized msg: %#v", m)
				}
			}
		}(i)
	}

	time.Sleep(3 * time.Second)
}
