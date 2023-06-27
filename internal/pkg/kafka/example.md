# kafka

Kafka client usage examples.

## kafka reader

```go
package main

import (
  "gf2-demo/internal/pkg/kafka"
)

func main() {
	done := make(chan struct{})
	ctx := gctx.New()
	ctx, cancel := context.WithCancel(ctx)

	reader, err := kafka.NewReader(ctx, kafka.WithBrokers([]string{
		"127.0.0.1:9092",
	}), kafka.WithTopic("test"), kafka.WithGroupID("test"))
	if err != nil {
		panic(err)
	}
	defer reader.Close()

	g.Log().Debug(ctx, "consumer up and running...")
	go consumeTasks(ctx, done)

	sigterm := make(chan os.Signal, 1)
			signal.Notify(
				sigterm,
				syscall.SIGHUP,
				syscall.SIGINT,
				syscall.SIGTERM,
				syscall.SIGQUIT,
	)

	select {
	case <-ctx.Done():
			g.Log().Info(ctx, "terminating: context canceled")
	case sig := <-sigterm:
			g.Log().Infof(ctx, "terminating: %s signal captured", sig)
	}

	cancel()

	<-done
}

func consumeTasks(ctx context.Context, done chan struct{}) {
	defer func(done chan struct{}) {
		done <- struct{}{}
	}(done)

	concurrency := g.Cfg().MustGet(ctx, "kafka.consumer_concurrency").Int()

	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		wg.Add(1)
		go taskDaemon(ctx, &wg, i)
	}

	wg.Wait()
}

func taskDaemon(ctx context.Context, wg *sync.WaitGroup, i int) {
	defer wg.Done()

	reader := kafka.GetReader()

	for {
		g.Log().Debugf(ctx, "goroutine %d runing...", i)

		select {
		case <-ctx.Done():
			g.Log().Debugf(ctx, "received cancel signal, goroutine %d done", i)
			return
		default:
			msg, err := reader.FetchMessage(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "[TaskMessageReceive] read message from kafka failed: %s", err)
				continue
			}

			g.Log().Infof(ctx,
				"[TaskMessageReceive] received message at topic: %s, partition: %d, offset: %d, value: %s, goroutine: %d",
				msg.Topic,
				msg.Partition,
				msg.Offset,
				msg.Value,
				i,
			)

			var t model.TaskQueue
			if err := gconv.Struct(msg.Value, &t); err != nil {
				g.Log().Errorf(ctx,
					"[TaskMessageReceive] unmarshal message failed, message at partition: %d, offset: %d, value: %s, error: %s",
					msg.Partition,
					msg.Offset,
					msg.Value,
					err,
				)

				commitMessage(ctx, msg)
				return
			}

			// 使之后代码逻辑中打印的日志的Trace-ID和web api请求的日志的Trace-Id一致, 便于排查问题
			ctx, err := gtrace.WithUUID(ctx, t.TraceID)
			if err != nil {
				g.Log().Warningf(ctx, "[TaskMessageReceive] inject trace_id %s into ctx failed: %s", t.TraceID, err)
			}

			// 处理任务逻辑
			task := service.Daemon().BuildTask(ctx, t)
			service.Daemon().Run(ctx, task)

			commitMessage(ctx, msg)
		}
	}
}

// 当前任务执行完再向kafka提交消息offset,
// 这样当任务执行中途程序意外退出的情况下,
// 程序重新运行后, 还会继续消费上次没成功的消息.
func commitMessage(ctx context.Context, msg kafkago.Message) {
	reader := kafka.GetReader()

	if err := reader.CommitMessages(ctx, msg); err != nil {
		g.Log().Errorf(
			ctx,
			"[TaskMessageCommit] commit message failed, message at partition: %d, offset: %d, value: %s, error: %s",
			msg.Partition,
			msg.Offset,
			msg.Value,
			err,
		)

		return
	}

	g.Log().Infof(
		ctx,
		"[TaskMessageCommit] commit message success, message at partition: %d, offset: %d, value: %s",
		msg.Partition,
		msg.Offset,
		msg.Value,
	)
}
```

## kafka writer

```go
import (
	kafkago "github.com/segmentio/kafka-go"

	"gf2-demo/internal/pkg/kafka"
)

func main() {
	writer, err := kafka.NewReader(ctx, kafka.WithBrokers([]string{
		"127.0.0.1:9092",
	}), kafka.WithTopic("test"))
	if err != nil {
		panic(err)
	}
	defer writer.Close()

	// 从ctx获取'Trace-Id'
	traceID := gctx.CtxId(ctx)

	// 将任务信息写入kafka
	msg := msgData{
		TaskID:   "xxx",
		TraceID:  traceID,
	}
	msgJson := gconv.String(msg)

	if err := writer.WriteMessages(ctx, kafkago.Message{
		Value: []byte(msgJson),
	}); err != nil {
		// handle error
	}
}
```
