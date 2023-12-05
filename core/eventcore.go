/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 15:55:17
 * @Description:
 */
package core

import (
	"context"
	"strings"
	"sync"

	"github.com/kwens/eventbus/errors"
	"github.com/kwens/eventbus/val"

	"go.uber.org/atomic"
)

type EventCore interface {
	Watch()
	CheckWatch() bool
	Exec(pubInfo *val.PubInfo)
}

type Core struct {
	close      atomic.Bool // 关闭判断
	isWatch    atomic.Bool // 是否开启监听
	OpenRemote bool        // [预设]是否开启支持远程事件服务，后续可以支持远程事件服务，比如nsq、nats等

	EventFunc  sync.Map // 订阅事件方法
	SuperSub   sync.Map // 超级订阅者执行事件
	GlobalSub  sync.Map // 全局订阅者执行事件
	ErrHandler errors.ErrHandler

	Queue  chan *val.PubInfo
	Cancel chan struct{}
}

func (eb *Core) CheckWatch() bool {
	return eb.isWatch.Load()
}

func (eb *Core) Watch() {
	// 校验开启情况，已经开启就
	if !eb.isWatch.Load() {
		if !eb.isWatch.CompareAndSwap(false, true) {
			return
		}
	} else {
		return
	}

	go func() {
		for {
			select {
			case pubInfo := <-eb.Queue:
				// exec
				go eb.Exec(pubInfo)
			case <-eb.Cancel:
				return
			}
		}
	}()
}

// ExecPub 发布执行动作
func (eb *Core) Exec(pubInfo *val.PubInfo) {
	var (
		err      error
		execData []*val.SubInfo
		execFunc *val.SubInfo
	)
	if execData, err = eb.getEventSub(pubInfo.EventName); err == nil {
		ctx := context.WithValue(context.Background(), val.EventCtxKey, pubInfo)
		for _, execFunc = range execData {
			if err = execFunc.Exec(ctx, pubInfo.Data); err != nil {
				eb.ErrHandler(pubInfo, err)
			}
		}
	} else {
		eb.ErrHandler(pubInfo, err)
	}
}

func (eb *Core) ExecSuper(pubInfo *val.PubInfo) {
	eb.SuperSub.Range(func(key, value any) bool {
		if strings.Contains(pubInfo.EventName, key.(string)) {
			if execFunc, ok := value.(val.EventCallBack); ok {
				ctx := context.WithValue(context.Background(), val.EventCtxKey, pubInfo)
				go execFunc(ctx, pubInfo.Data)
			}
		}
		return true
	})
	eb.GlobalSub.Range(func(key, value any) bool {
		if execFunc, ok := value.(val.EventCallBack); ok {
			ctx := context.WithValue(context.Background(), val.EventCtxKey, pubInfo)
			go execFunc(ctx, pubInfo.Data)
		}
		return true
	})
}

func (eb *Core) Stop() {
	if !eb.close.Load() {
		if eb.close.CompareAndSwap(false, true) {
			close(eb.Cancel)
		}
	}
}

func (eb *Core) getEventSub(eventName string) ([]*val.SubInfo, error) {
	if data, ok := eb.EventFunc.Load(eventName); ok {
		if subInfo, ok := data.([]*val.SubInfo); ok {
			return subInfo, nil
		}
		return nil, errors.ErrNoExpectEvent
	}
	return nil, errors.ErrEventNoExist
}
