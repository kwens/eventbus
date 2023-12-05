/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 16:31:59
 * @Description:
 */
package sub

import (
	"sync"

	"github.com/kwens/eventbus/core"
	"github.com/kwens/eventbus/errors"
	"github.com/kwens/eventbus/val"
)

type EventSub interface {
	Subscribe(eventName string, callback val.EventCallBack, opts ...SubOption) error
	UnSubscribe(eventName string, subName string)
}

type Sub struct {
	*core.Core
	Consumers sync.Map // [预设]远程订阅消费者
}

func (sub *Sub) Subscribe(eventName string, callback val.EventCallBack, opts ...SubOption) error {
	execInfo := &val.SubInfo{
		EventName: eventName,
		Exec:      callback,
	}
	return sub.doSub(execInfo, opts...)
}

func (sub *Sub) doSub(subInfo *val.SubInfo, opts ...SubOption) error {
	dOpt := defaultSubOpt()
	for _, opt := range opts {
		opt(dOpt)
	}
	subInfo.SubName = dOpt.subName
	if len(subInfo.SubName) == 0 {
		subInfo.SubName = val.DefalutSubName
	}
	// 已存在事件订阅，添加回调到最后
	if event, ok := sub.EventFunc.Load(subInfo.EventName); ok {
		if eventFuncs, ok := event.([]*val.SubInfo); ok {
			for _, v := range eventFuncs {
				if v.SubName == subInfo.SubName {
					return errors.ErrSubNameExist
				}
			}
			eventFuncs = append(eventFuncs, subInfo)
			sub.EventFunc.Store(subInfo.EventName, eventFuncs)
			return nil
		}
	}
	// 不存在事件订阅
	sub.EventFunc.Store(subInfo.EventName, []*val.SubInfo{subInfo})
	return nil
}

func (sub *Sub) UnSubscribe(eventName string, subName string) {
	if len(subName) == 0 {
		subName = val.DefalutSubName
	}
	eventFuncs, ok := sub.EventFunc.Load(eventName)
	if ok {
		if subInfo, ok := eventFuncs.([]*val.SubInfo); ok {
			for i, sub := range subInfo {
				if sub.SubName == subName {
					subInfo = append(subInfo[:i], subInfo[i+1:]...)
					break
				}
			}
			if len(subInfo) == 0 {
				sub.DelEvent(eventName)
				return
			}
			sub.EventFunc.Store(eventName, subInfo)
		}
	}
}

// DelEvent 删除事件
func (sub *Sub) DelEvent(eventName string) {
	sub.EventFunc.Delete(eventName)
}

func (sub *Sub) SubscribeSuper(superName string, callback val.EventCallBack) error {
	if len(superName) == 0 {
		return errors.ErrSuperNameEmpty
	}
	sub.Core.SuperSub.Store(superName, callback)
	return nil
}

func (sub *Sub) SubscribeGlobal(globalName string, callback val.EventCallBack) error {
	if len(globalName) == 0 {
		return errors.ErrSuperNameEmpty
	}
	sub.Core.GlobalSub.Store(globalName, callback)
	return nil
}
