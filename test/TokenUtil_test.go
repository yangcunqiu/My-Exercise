package test

import (
	"My-Exercise/model/entity"
	"My-Exercise/utils"
	"log"
	"testing"
)

func TestGenerateToken(t *testing.T) {
	caseUser := entity.User{
		Id:   1,
		Name: "ycq",
	}
	token, _ := utils.GenerateToken(caseUser.Id, caseUser.Name)
	log.Println(token)
}

func TestParseToken(t *testing.T) {
	caseTokenString := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJJZCI6MSwiTmFtZSI6InljcSJ9.wQyw_FKePEiq-2PkjRW3O-Vk8Mzp8vFosh6rKEMDDT4"
	userClaims, ok := utils.ParseToken(caseTokenString)
	if ok {
		log.Println("校验成功\n", userClaims)
	} else {
		log.Println("校验失败")
	}
}
