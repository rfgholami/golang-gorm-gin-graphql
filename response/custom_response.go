package response

type Response struct {
	Code     int         `json:"code"`
	Status   string      `json:"status"`
	Message  string      `json:"message"`
	Error    string      `json:"error"`
	RowCount int         `json:"rowCount"`
	Data     interface{} `json:"data"`
}
