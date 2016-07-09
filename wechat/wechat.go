package wechat

import (
	"crypto/sha1"
	"fmt"
	"io"
	"sort"

	log "github.com/Sirupsen/logrus"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	wcToken            string
	wcAppID            string
	wcAppSecret        string
	wcEncoding_AES_Key string
	access_token       string
)

func Init() {
	wcToken = viper.GetString("wechat.token")
	wcAppID = viper.GetString("wechat.app_id")
	wcAppSecret = viper.GetString("wechat.app_secret")
	wcEncoding_AES_Key = viper.GetString("wechat.encoding_aes_key")

	if wcToken == "" {
		log.Fatal("invalid config wechat.token")
	}
	if wcAppID == "" {
		log.Fatal("invalid config wechat.app_id")
	}
	if wcAppSecret == "" {
		log.Fatal("invalid config wechat.app_secret")
	}

	log.WithFields(log.Fields{
		"wcToken":            wcToken,
		"wcAppID":            wcAppID,
		"wcAppSecret":        wcAppSecret,
		"wcEncoding_AES_Key": wcEncoding_AES_Key,
	}).Debug("get wechat config.")
	//	RegDB()
}

func RegDB() {
	orm.RegisterModel(new(TextMsg),
		new(ImageMsg),
		new(LinkMsg),
		new(LocationMsg),
		new(VideoMsg),
		new(VoiceMsg))
}

// Auth 服务器验证
func Auth(signature, timestamp, nonce, echostr string) bool {
	tmps := []string{wcToken, timestamp, nonce}
	sort.Strings(tmps)
	tmpStr := tmps[0] + tmps[1] + tmps[2]

	tmp := func(data string) string {
		t := sha1.New()
		io.WriteString(t, data)
		return fmt.Sprintf("%x", t.Sum(nil))
	}(tmpStr)

	if tmp == signature {
		return true
	}
	log.Warning("auth receive: ", tmp)
	log.Warning("auth should be: ", signature)
	return false
}
