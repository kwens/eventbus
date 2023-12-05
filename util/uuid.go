/**
 * @Author: kwens
 *
 * @Date: 2023-06-05 17:07:58
 * @Description:
 */
package util

import (
	"encoding/hex"

	uuid "github.com/satori/go.uuid"
)

type uuID [32]byte

// GetUUID 获取UUid
func GetUUID() string {
	buf := uuID{}
	hex.Encode(buf[:], uuid.NewV4().Bytes())
	return string(buf[:])
}
