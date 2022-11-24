package entity

import "My-Exercise/model"

type Submit struct {
	Id        uint
	ProblemId int    `gorm:"comment:题目id"`
	UserId    int    `gorm:"comment:提交用户id"`
	CodePath  string `gorm:"type:varchar(255);comment:提交代码存储地址"`
	status    int    `gorm:"comment:状态(0-待判断, 1-答案正确, 2-答案错误, 3-提交超时, 4-超出最大内存限制)"`
	model.BaseInfo
}
