/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 15:58:45
 * @Description:
 */
package val

import "context"

type EventCallBack func(context.Context, interface{}) error

const (
	DefaultMaxQueue = 1000
)

var (
	EventCtxKey = struct{}{}
	TraceIDKey  = struct{}{}
)

type SendMode int

const (
	DefaultSend SendMode = iota // 默认发送模式，本地&远程
	OnlyLocal
	OnlyRemote
)

const DefalutSubName = "defaultSubscriber"

type PubInfo struct {
	EventID   string      `json:"EventId"`   // 事件id
	EventName string      `json:"EventName"` // 事件名称
	Data      interface{} `json:"Data"`      // 数据
}

type SubInfo struct {
	EventName string
	SubName   string
	Exec      EventCallBack
}
