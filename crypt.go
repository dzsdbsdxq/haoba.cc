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
	ran := rand.New(rand.NewSource(time.Now().Unix()))
	lockCount := ran.Intn(lockLen)

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
		encodeStr = encodeStr + string(stream[j])
		k++
	}
	return encodeStr + randomLock
}

/**
 * 可逆解密
 *
 * @param  textStream 要解密的字符串
 * @param  password 加密私钥=解密私钥
 */
func decrypt(textSteam string, password string) (decodeStr string) {
	var k int
	var j int
	stream := []byte(lockStream)

	lockLen := len(stream)

	// 截取随机密锁值
	textLen := len(textSteam)
	randomLock := textSteam[textLen-1]
	// 获得随机密码值的位置
	lockCount := bytes.IndexFunc(stream, func(r rune) bool {
		return r == rune(randomLock)
	})
	password = func(str string) string {
		h := md5.New()
		h.Write([]byte(str))
		return hex.EncodeToString(h.Sum(nil))
	}(password + string(randomLock))

	// 开始对字符串解密
	textSteam = textSteam[:textLen-1]
	for _, ts := range textSteam {
		if k == len(password) {
			k = 0
		}
		j = bytes.IndexFunc(stream, func(r rune) bool {
			return r == ts
		})

		j = j - lockCount - int(password[k])

		for {
			if j >= 0 {
				break
			}
			j = j + lockLen
		}
		decodeStr += string(lockStream[j])
		k++
	}
	//fmt.Println(decodeStr)
	decodeByte, _ := base64.StdEncoding.DecodeString(decodeStr)
	decodeStr = string(decodeByte)
	return
}

func main() {
	s := encrypt("dzsdbsdxq", "123456")
	fmt.Println(s)
	d := decrypt(s, "123456")
	fmt.Println(d)
}
