package query

type ProblemSave struct {
	Id             uint
	Title          string
	Content        string
	Timeout        int
	MaxMemory      int
	CategoryIdList []uint
	TestCaseList   []TestCaseSave
}
