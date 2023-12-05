# eventbus

一个简单的事件总线，目前仅支持简单本地事件发布和订阅,暂不带工厂

# Usage

基本使用

```go

var UserEvent val.EventCallBack = func(ctx context.Context, data interface{}) error {
	fmt.Println("UserEvent Doing...")
	fmt.Printf("data is:%+v", data)
	return nil
}

func main() {
	// init
	eb := eventbus.New()

	// subscirbe
	if err := eb.Subscribe("user.login", UserEvent); err != nil {
		panic(err)
	}

	// publish
	if err := eb.Publish(context.Background(), "user.login", "hello world"); err!= nil {
		panic(err)
	}
}
```


# Val

```go

import "eventbus/val" // 注意在这个包里面

var (
	EventCtxKey = struct{}{} // 事件上下文key
	TraceIDKey  = struct{}{} // 链路追逐上下文key
)

type SendMode int

const (
	DefaultSend SendMode = iota // 默认发送模式，本地&远程
	OnlyLocal
	OnlyRemote
)


```

# 选项

```go
// EventBusOption
type eventBusOption struct {
	MaxQueue   int
	ErrHandler errors.ErrHandler
}
// WithEventMaxQueue 最大队列数
func WithEventMaxQueue(max int)
// WithErrHandler 错误处理方法
func WithErrHandler(h errors.ErrHandler)

// pub option
// eventbus/pub 在这个包里面
type pubOption struct {
	sendMode val.SendMode // 发送模式
}
func WithPubSendMode(sendMode val.SendMode)

// sub option
// eventbus/sub 在这个包里面
type subOption struct {
	subName string
}
func WithSubName(name string)
```

# TODO

- [ ] logger
- [ ] 远程事件服务，比如：nsq、nats等等