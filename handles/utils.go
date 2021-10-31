package handles

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

func MakeMD5(st string) string {
	m := md5.New()
	m.Write([]byte(st+"abcd"+time.Now().String()))
	return hex.EncodeToString(m.Sum(nil))
}
