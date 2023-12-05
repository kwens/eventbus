/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 16:02:14
 * @Description:
 */
package pub

import (
	"context"

	"github.com/kwens/eventbus/core"
	"github.com/kwens/eventbus/errors"
	"github.com/kwens/eventbus/util"
	"github.com/kwens/eventbus/val"
)

type EventPub interface {
	Publish(ctx context.Context, eventName string, data interface{}) error
}

type Pub struct {
	*core.Core
	BeforeHandler func(val.PubInfo)
}

func (pub *Pub) Publish(ctx context.Context, eventName string, data interface{}) error {
	info := &val.PubInfo{
		EventName: eventName,
		Data:      data,
	}
	return pub.do(ctx, info)
}

func (pub *Pub) do(ctx context.Context, pubInfo *val.PubInfo, opts ...PubOption) error {
	dOpt := defaultPubOpt()
	for _, opt := range opts {
		opt(dOpt)
	}
	traceId, ok := ctx.Value(val.TraceIDKey).(string)
	if ok {
		pubInfo.EventID = traceId
	}
	if len(pubInfo.EventID) == 0 {
		pubInfo.EventID = util.GetUUID()
	}

	// 执行前回调
	if pub.BeforeHandler != nil {
		pub.BeforeHandler(*pubInfo)
	}
	switch dOpt.sendMode {
	case val.DefaultSend:
		return pub.doDefault(pubInfo)
	case val.OnlyLocal:
		return pub.doLocal(pubInfo)
	case val.OnlyRemote:
		return pub.doRemote(pubInfo)
	default:
		return errors.ErrUnknownSendMode
	}
}

// doDefault 默认发布本地和远程
func (pub *Pub) doDefault(pubInfo *val.PubInfo) error {
	if pub.OpenRemote {
		if err := pub.doRemote(pubInfo); err != nil {
			return err
		}
	}
	return pub.doLocal(pubInfo)
}

// doLocal 发布本地
func (pub *Pub) doLocal(pubInfo *val.PubInfo) error {
	if !pub.CheckWatch() {
		return errors.ErrLocalWatchClose
	}

	// 执行超级订阅者
	go pub.ExecSuper(pubInfo)
	// 执行普通订阅者(防止堵塞)
	select {
	case pub.Core.Queue <- pubInfo:
	default:
		go func() {
			pub.Core.Queue <- pubInfo
		}()
	}
	return nil
}

// doRemote 发布远程
func (pub *Pub) doRemote(pubInfo *val.PubInfo) error {
	// todo
	return nil
}
