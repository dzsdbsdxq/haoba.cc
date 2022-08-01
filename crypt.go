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

const lockStream = "st=lDEFABCkVWXYZabc89LMmGH012345uvdefIJK6NOPyzghijQRSTUwx7nopqr"

/**
 * 可逆加密
 *
 * @param  textSteam 要加密的字符串
 * @param  password 加密私钥=解密私钥
 */
func encrypt(textSteam string, password string) (encodeStr string) {
	var k int
	var j int

	stream := []byte(lockStream)

	lockLen := len(stream)
	rand.Seed(time.Now().UnixNano())
	lockCount := rand.Intn(lockLen - 1)
	randomLock := string(stream[lockCount])

	password = func(str string) string {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}(password + randomLock)

	textSteam = base64.StdEncoding.EncodeToString([]byte(textSteam))

	for _, ts := range textSteam {
		if k == len(password) {
			k = 0
		}
		j = bytes.IndexFunc(stream, func(r rune) bool {
			return r == ts
		})
		j = (j + lockCount + int(password[k])) % (lockLen)
		encodeStr += string(stream[j])
		k++
	}
	return
}

func main() {
	s := encrypt("dzsdbsdxq", "123456")
	fmt.Println(s)
	//来个大佬写解密

}
