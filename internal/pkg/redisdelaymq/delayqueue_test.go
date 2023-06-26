//go:build ignore
// +build ignore

package redisdelaymq

import (
	"testing"
	"time"

	_ "github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gctx"
	uuid "github.com/satori/go.uuid"
)

type message struct {
	ID   string
	Name string
}

// TestDelayQueue 使用流程示例: go test -v
func TestDelayQueue(t *testing.T) {
	ctx := gctx.New()

	// *** 此处配置仅测试使用, 正式使用请在manifest/config/config.yml中配置
	redisConfig := gredis.Config{
		Address: "127.0.0.1:6379",
		Db:      1,
		Pass:    "",
	}
	redisGroup := "delayqueue" // 默认: default
	gredis.SetConfig(&redisConfig, redisGroup)
	// ****************************************

	// 创建一个延迟队列dq, 使消息延迟3秒后才能被消费
	topic := "dq"
	delay := 3 * time.Second

	// 如果是默认的default分组, 可忽略此选项.
	groupOption := WithRedisGroup(redisGroup)

	dq := New(topic, delay, groupOption)

	msg1 := message{
		ID:   uuid.NewV4().String(),
		Name: "hello",
	}
	msg2 := message{
		ID:   uuid.NewV4().String(),
		Name: "world",
	}

	var msgs []interface{}
	msgs = append(msgs, msg1, msg2)

	// 生产者
	err := dq.Produce(ctx, msgs)
	if err != nil {
		t.Errorf("[PRODUCER] produce msg failed: %v", err)
		return
	}

	t.Logf("[PRODUCER] produce msg: %#v, delay: %s", msgs, delay.String())

	// 消费者
	for {
		// 消费score小于当前时间的消息
		now := time.Now().Unix()
		res, err := dq.Consume(ctx, now)
		if err != nil {
			t.Errorf("[CONSUMER] get msg failed: %v", err)
			continue
		}

		if res == nil {
			t.Logf("[CONSUME] no msg less than given '%d', try after 1 second", now)
			time.Sleep(1 * time.Second)
			continue
		}

		var msgs []message

		_ = res.Scan(&msgs)

		t.Logf("[CONSUME] got msgs: %#v", msgs)

		// 模拟处理消息所代表的任务
		for _, msg := range msgs {
			t.Logf("[CONSUME] handle msg: %#v", msg)
		}

		// 消息任务处理完成后删除消息, 避免被重复消费.
		n, err := dq.Remove(ctx, now)
		if err != nil {
			t.Errorf("[REMOVE] remove msg failed: %v", err)
			continue
		}

		t.Logf("[REMOVE] msgs removed count: %d", n)

		break
	}
}
