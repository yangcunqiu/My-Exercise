package test

import (
	"My-Exercise/utils"
	"log"
	"testing"
)

func TestGenerateMD5(t *testing.T) {
	caseStr := "123456"
	md5 := utils.GenerateMD5(caseStr)
	log.Println(md5)
}
