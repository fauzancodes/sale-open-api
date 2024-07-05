package utils

import (
	"sale-open-api/res"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

func PopulatePaging(c echo.Context, custom string) (param res.ReqPaging) {
	customval := c.QueryParam(custom)
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	if limit == 0 {
		limit = 10
	}
	offset, _ := strconv.Atoi(c.QueryParam("offset"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page == 0 && offset == 0 {
		page = 1
		offset = 0
	}
	if page >= 1 && offset == 0 {
		offset = (page - 1) * limit
	}
	draw, _ := strconv.Atoi(c.QueryParam("draw"))
	if draw == 0 {
		draw = 1
	}
	order := c.QueryParam("sort")
	if strings.ToLower(order) == "asc" {
		order = "ASC"
	} else {
		order = "DESC"
	}
	sort := c.QueryParam("order")
	if sort == "" {
		sort = "id"
	}
	param = res.ReqPaging{
		Search: c.QueryParam("search"),
		Order:  order,
		Limit:  limit,
		Offset: offset,
		Sort:   sort,
		Custom: customval,
		Page:   page}
	return
}

func PopulateResPaging(param *res.ReqPaging, data interface{}, totalResult int64, totalFiltered int64) (output res.ResPaging) {
	totalPages := int(totalResult) / param.Limit
	if int(totalResult)%param.Limit > 0 {
		totalPages++
	}

	currentPage := param.Offset/param.Limit + 1
	next := false
	back := false
	if currentPage < totalPages {
		next = true
	}
	if currentPage <= totalPages && currentPage > 1 {
		back = true
	}

	output = res.ResPaging{
		Status:            200,
		Data:              data,
		Search:            param.Search,
		Order:             param.Order,
		Limit:             param.Limit,
		Offset:            param.Offset,
		Sort:              param.Sort,
		Next:              next,
		Back:              back,
		TotalData:         int(totalResult),
		TotalDataFiltered: int(totalFiltered),
		CurrentPage:       currentPage,
		TotalPage:         totalPages,
	}
	return
}
