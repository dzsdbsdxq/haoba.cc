package main

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

func encrypt(textSteam string, password string) (encodeStr string) {
	var k int
	var j int
	
	// 随机找一个数字，并从密锁串中找到一个密锁值
	lockStream := []byte("st=lDEFABCkVWXYZabc89LMmGH012345uvdefIJK6NOPyzghijQRSTUwx7nopqr")

	lockLen := len(lockStream)
	rand.Seed(time.Now().UnixNano())
	lockCount := rand.Intn(lockLen - 1)
	randomLock := string(lockStream[lockCount])
	// 结合随机密锁值生成MD5后的密码
	password = func(str string) string {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}(password + randomLock)
	// 开始对字符串加密
	textSteam = base64.StdEncoding.EncodeToString([]byte(textSteam))

	for _, ele := range textSteam {
		if k == len(password) {
			k = 0
		}
		j = bytes.IndexFunc(lockStream, func(r rune) bool {
			return r == ele
		})
		j = (j + lockCount + int(password[k])) % (lockLen)
		encodeStr += string(lockStream[j])
		k++
	}
	return
}

func main() {
	s := encrypt("jackielee", "123456")
	fmt.Println(s)
	//来个大佬写解密

}
