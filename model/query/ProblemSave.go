package query

type ProblemSave struct {
	Title          string
	Content        string
	Timeout        int
	MaxMemory      int
	CategoryIdList []uint
	TestCaseList   []TestCaseSave
}
