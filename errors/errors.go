/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 16:36:16
 * @Description:
 */
package errors

import (
	"errors"

	"github.com/kwens/eventbus/log"
	"github.com/kwens/eventbus/val"
)

type ErrHandler func(info *val.PubInfo, err error)

var DefalutErrHandler ErrHandler = func(info *val.PubInfo, err error) {
	log.DefaultLogger.Log(log.LevelError, "[EventBus] info:%+v err:%+v", info, err)
}

var (
	ErrEventNoExist    = errors.New("event not exist")
	ErrNoExpectEvent   = errors.New("no expect event")
	ErrUnknownSendMode = errors.New("unknown send mode")
	ErrLocalWatchClose = errors.New("local event watch is close")
	ErrSubNameExist    = errors.New("subName existed")
	ErrSuperNameEmpty  = errors.New("super subscribe name is empty")
)
