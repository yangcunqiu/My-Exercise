package entity

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"My-Exercise/model/query"
	"gorm.io/gorm"
)

type Submit struct {
	Id        uint     `json:"id"`
	ProblemId int      `json:"problemId" gorm:"comment:题目id"`
	Problem   *Problem `json:"problem"`
	UserId    int      `json:"userId" gorm:"comment:提交用户id"`
	User      *User    `json:"user"`
	CodePath  string   `json:"codePath" gorm:"type:varchar(255);comment:提交代码存储地址"`
	Status    int      `json:"status" gorm:"comment:状态(0-待判断, 1-答案正确, 2-答案错误, 3-提交超时, 4-超出最大内存限制)"`
	model.BaseInfo
}

func (s Submit) TableName() string {
	return "submit"
}

func ListSubmit(query query.SubmitListQuery) *gorm.DB {
	tx := global.DB.Debug().Model(new(Submit)).Preload("Problem", func(DB *gorm.DB) *gorm.DB {
		return DB.Omit("content")
	}).Preload("User")
	if query.ProblemIds != nil {
		tx.Where("problem_id in ?", query.ProblemIds)
	}
	if query.UserIds != nil {
		tx.Where("user_id in ?", query.UserIds)
	}
	if query.Status != nil {
		tx.Where("status in ?", query.Status)
	}
	return tx

}
