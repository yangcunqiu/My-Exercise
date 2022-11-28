package utils

import (
	"github.com/jordan-wright/email"
	"log"
	"net/smtp"
	"testing"
)

func TestSendEmail(t *testing.T) {
	em := email.NewEmail()
	// 设置 sender 发送方 的邮箱 ， 此处可以填写自己的邮箱
	em.From = "ycq <903421928@qq.com>"

	// 设置 receiver 接收方 的邮箱  此处也可以填写自己的邮箱， 就是自己发邮件给自己
	em.To = []string{"yangcunqiuup@163.com"}

	// 设置主题
	em.Subject = "Test"

	// 简单设置文件发送的内容，暂时设置成纯文本
	em.HTML = []byte("<b>hello world， 咱们用 golang 发个邮件！！<b>")

	// 设置服务器相关的配置
	err := em.Send("smtp.qq.com:25", smtp.PlainAuth("", "903421928@qq.com", "jrwfxdmbzphzbcgj", "smtp.qq.com"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("send successfully ... ")
}
