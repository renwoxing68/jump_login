package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"github.com/go-ini/ini"
	"log"
	"time"
)

var Cfg *ini.File

type TfActorAuthenticator struct {
	Issuer    string
	Algorithm string
	Period    int64
	Digits    int
}

func main() {
	var err error
	Cfg, err = ini.Load("conf/conf")
	if err != nil {
		log.Fatal("Fail to Load ‘conf/conf’:", err)
	}

	//直接读取
	prefix := Cfg.Section("").Key("PREFIX").MustString("")
	secret := Cfg.Section("").Key("SECRET").MustString("")
	suffix := Cfg.Section("").Key("SUFFIX").MustString("")
	authClient := TfActorAuthenticator{
		Issuer:    "",
		Algorithm: "SHA1",
		Period:    30,
		Digits:    6,
	}
	token := prefix + authClient.TotpString(secret) + suffix
	fmt.Println(token)
}

func (tfa *TfActorAuthenticator) TotpString(secret string) string {
	// base32编码秘钥：K
	key := make([]byte, base32.StdEncoding.DecodedLen(len(secret)))
	base32.StdEncoding.Decode(key, []byte(secret))

	// 根据当前时间计算随机数：C
	message := make([]byte, 8)
	binary.BigEndian.PutUint64(message, uint64(time.Now().Unix()/tfa.Period))

	// 使用sha1对K和C做hmac得到20个字节的密串：HMAC-SHA-1(K, C)
	hmacsha1 := hmac.New(sha1.New, key)
	hmacsha1.Write(message)
	hash := hmacsha1.Sum([]byte{})

	// 从20个字节的密串取最后一个字节，取该字节的低四位
	offset := hash[len(hash)-1] & 0xF
	truncatedHash := hash[offset : offset+4]

	// 按照大端方式组成一个整数
	bin := (binary.BigEndian.Uint32(truncatedHash) & 0x7FFFFFFF) % 1000000

	// 将数字转成特定长度字符串，不够的补0
	return fmt.Sprintf(`%0*d`, tfa.Digits, bin)
}
