package utils

import (
	"log"
	"testing"
)

func TestGenerateMD5(t *testing.T) {
	caseStr := "123456"
	md5 := GenerateMD5(caseStr)
	log.Println(md5)
}
