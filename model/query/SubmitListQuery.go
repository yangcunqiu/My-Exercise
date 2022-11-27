package query

type SubmitListQuery struct {
	ProblemIds []int `json:"problemIds"`
	UserIds    []int `json:"userIds"`
	Status     []int `json:"status"`
}
