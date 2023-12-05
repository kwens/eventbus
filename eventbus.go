/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 17:56:01
 * @Description:
 */
package eventbus

import (
	"github.com/kwens/eventbus/core"
	"github.com/kwens/eventbus/errors"
	"github.com/kwens/eventbus/pub"
	"github.com/kwens/eventbus/sub"
	"github.com/kwens/eventbus/val"
)

type EventBus interface {
	core.EventCore
	pub.EventPub
	sub.EventSub
}

type Bus struct {
	*core.Core
	*pub.Pub
	*sub.Sub
}

func New(opts ...EventBusOption) EventBus {
	event := defaultEventBus()

	dOpt := defaultEventBusOpt()
	for _, opt := range opts {
		opt(dOpt)
	}
	event.Queue = make(chan *val.PubInfo, dOpt.MaxQueue)
	// 开启监听
	go event.Watch()
	return event
}

func defaultEventBus() *Bus {
	eventCore := &core.Core{
		Queue:      make(chan *val.PubInfo, val.DefaultMaxQueue),
		Cancel:     make(chan struct{}),
		ErrHandler: errors.DefalutErrHandler,
	}
	return &Bus{
		Core: eventCore,
		Pub: &pub.Pub{
			Core: eventCore,
		},
		Sub: &sub.Sub{
			Core: eventCore,
		},
	}
}

func (b *Bus) Start() error { return nil }
func (b *Bus) Stop() error  { return nil }
