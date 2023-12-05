/**
 * @Author: kwens
 *
 * @Date: 2023-06-06 08:54:08
 * @Description:
 */
package pub

import "github.com/kwens/eventbus/val"

type PubOption func(*pubOption)

type pubOption struct {
	sendMode val.SendMode
}

func defaultPubOpt() *pubOption {
	return &pubOption{}
}

func WithPubSendMode(sendMode val.SendMode) PubOption {
	return func(po *pubOption) {
		po.sendMode = sendMode
	}
}
