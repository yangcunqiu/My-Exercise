package test

import (
	"My-Exercise/utils"
	"log"
	"testing"
)

func TestSendEmail(t *testing.T) {
	code := "123456"
	htmlStr := "<b>您的验证码是: " + code + "<b>"
	err := utils.SendEmail("TestUtil", htmlStr, "yangcunqiuup@163.com")
	if err != nil {
		log.Println(err)
	} else {
		log.Println("send email success")
	}
}
