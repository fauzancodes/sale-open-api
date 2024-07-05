package res

type ReqPaging struct {
	Page   int         `default:"1"`
	Search string      `default:""`
	Limit  int         `default:"10"`
	Offset int         `default:"0"`
	Sort   string      `default:"ASC"`
	Order  string      `default:"id"`
	Custom interface{} `default:""`
}

type ResPaging struct {
	TotalData         int         `json:"total_data"`
	TotalDataFiltered int         `json:"total_data_filtered"`
	CurrentPage       int         `default:"1" json:"current_page"`
	TotalPage         int         `default:"0" json:"total_page"`
	Next              bool        `default:"false" json:"next"`
	Back              bool        `default:"false" json:"back"`
	Search            string      `default:"" json:"search"`
	Limit             int         `default:"10" json:"limit"`
	Offset            int         `default:"0" json:"offset"`
	Sort              string      `default:"ASC" json:"sort"`
	Order             string      `default:"id" json:"order"`
	Error             string      `json:"error"`
	Status            int         `default:"200" json:"status"`
	Messages          string      `default:"Success" json:"message"`
	Data              interface{} `default:"[]" json:"data"`
}
