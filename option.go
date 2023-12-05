/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 15:56:17
 * @Description:
 */
package eventbus

import (
	"github.com/kwens/eventbus/errors"
	"github.com/kwens/eventbus/val"
)

type EventBusOption func(*eventBusOption)
type eventBusOption struct {
	MaxQueue   int
	ErrHandler errors.ErrHandler
}

func defaultEventBusOpt() *eventBusOption {
	return &eventBusOption{
		MaxQueue:   val.DefaultMaxQueue,
		ErrHandler: errors.DefalutErrHandler,
	}
}

// WithEventMaxQueue 最大队列数
func WithEventMaxQueue(max int) EventBusOption {
	return func(ebo *eventBusOption) {
		ebo.MaxQueue = max
	}
}

// WithErrHandler 错误处理方法
func WithErrHandler(h errors.ErrHandler) EventBusOption {
	return func(ebo *eventBusOption) {
		ebo.ErrHandler = h
	}
}
