package test

import (
	"My-Exercise/utils"
	"log"
	"testing"
)

func TestGenerateRandomNumberToString(t *testing.T) {
	randomNum6 := utils.GenerateRandomNumberToString(6)
	randomNum10 := utils.GenerateRandomNumberToString(10)
	randomNum1 := utils.GenerateRandomNumberToString(1)
	log.Println(randomNum6)
	log.Println(randomNum10)
	log.Println(randomNum1)
}
