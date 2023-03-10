package util

/*
 *
 *  SiGG-Proof-Einstein-Rosen-Bridge-Theory
 *
 */

import "math"

type Pagination struct {
	Total       uint        `json:"total"`
	PerPage     uint        `json:"per_page"`
	CurrentPage uint        `json:"current_page"`
	LastPage    uint        `json:"last_page"`
	From        uint        `json:"from"`
	To          uint        `json:"to"`
	Data        interface{} `json:"data"`
}

func Paginate(page, pageSize, count uint, data interface{}) *Pagination {
	var lastPage uint = 1

	to := page * pageSize

	if to > count {
		to = count
	}

	from := (page-1)*pageSize + 1

	if count == 0 || from > count {
		return &Pagination{PerPage: pageSize, CurrentPage: page, LastPage: lastPage, Data: data}
	}

	lastPage = uint(math.Ceil(float64(count) / float64(pageSize)))

	return &Pagination{
		Total: count, PerPage: pageSize, CurrentPage: page,
		LastPage: lastPage, From: from, To: to, Data: data,
	}
}

func CanPaginate(page, pageSize, count uint) bool {
	if page == 0 {
		page = 1
	}

	from := (page-1)*pageSize + 1
	if count == 0 || from > count {
		return false
	}

	return true
}

func PurePageArgs(page uint, pageSize uint) (uint, uint) {
	if page == 0 {
		page = 1
	}
	if pageSize == 0 {
		pageSize = 20
	}

	if pageSize > 100 {
		pageSize = 100
	}

	return page, pageSize
}

func EmptyPagination(page, pageSize uint) *Pagination {
	return Paginate(page, pageSize, 0, make([]interface{}, 0))
}
