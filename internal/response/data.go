package response

type Data struct {
	Data interface{} `json:"data"`
}

type DataWithMessage struct {
	Data WithMessage `json:"data"`
}
