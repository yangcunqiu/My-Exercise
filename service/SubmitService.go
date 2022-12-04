package service

import (
	"My-Exercise/global"
	"My-Exercise/model"
	"My-Exercise/model/dto"
	"My-Exercise/model/entity"
	"My-Exercise/model/query"
	"My-Exercise/ranner"
	"My-Exercise/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"io"
	"log"
	"strconv"
	"sync"
	"time"
)

func SubmitList(c *gin.Context) {
	submitListQuery := new(query.SubmitListQuery)
	_ = c.ShouldBindJSON(submitListQuery)

	var total int64
	submitList := make([]entity.Submit, 0)
	pageNum, pageSize, offset := model.PageParams(c)
	tx := entity.ListSubmit(*submitListQuery)
	tx.Count(&total).Offset(offset).Limit(pageSize).Find(&submitList)

	utils.Success(c, model.PageOf(pageNum, pageSize, total, submitList))
}

func ProblemSubmit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	bytes, _ := io.ReadAll(c.Request.Body)

	problem := new(entity.Problem)
	err := global.DB.Model(new(entity.Problem)).First(problem, id).Error
	if err == gorm.ErrRecordNotFound {
		utils.Fail(c, model.ErrorCodeOf(90001, "问题不存在"))
	}

	value, _ := c.Get("user")
	user := value.(*entity.User)
	// 保存代码
	codePath, err := utils.SaveCode(bytes, user.Name)
	if err != nil {
		utils.Fail(c, model.ErrorCodeOf(90002, "代码保存失败"))
	}

	var status int
	var passCount *int = new(int)
	*passCount = 0
	mu := &sync.Mutex{}
	// 判断代码
	testCaseList := GetTestCaseByProblemId(id)
	answerError := make(chan struct{})
	compileError := make(chan struct{})
	maxMemError := make(chan struct{})
	for _, testCase := range testCaseList {
		codeDTO := &dto.CodeDTO{
			Path:   codePath,
			Input:  testCase.Input,
			Output: testCase.Output,
			MaxMem: problem.MaxMemory,
		}
		go ranner.RannerCode(codeDTO, passCount, mu, answerError, compileError, maxMemError)
	}

	select {
	case <-answerError:
		status = 2
	case <-compileError:
		status = 3
	case <-maxMemError:
		status = 5
	case <-time.After(time.Millisecond * time.Duration(problem.Timeout)):
		if *passCount == len(testCaseList) {
			status = 1
		} else {
			log.Printf("code ranner timeout-[%v], %v", problem.Timeout, time.Millisecond*time.Duration(problem.Timeout))
			status = 4
		}
	}
	submit := &entity.Submit{
		ProblemId: id,
		UserId:    int(user.Id),
		CodePath:  codePath,
		Status:    status,
	}

	// 保存记录
	global.DB.Model(new(entity.Submit)).Create(submit)
	utils.Success(c, status)
}
