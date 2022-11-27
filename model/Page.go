package model

type Page struct {
	PageNum  int   `json:"pageNum"`
	PageSize int   `json:"pageSize"`
	Total    int64 `json:"total"`
	List     any   `json:"list"`
}

func PageOf(pageNumber, pageSize int, total int64, list any) Page {
	return Page{
		PageNum:  pageNumber,
		PageSize: pageSize,
		Total:    total,
		List:     list,
	}
}
