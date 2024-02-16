package response

type Error struct {
	Error interface{} `json:"error"`
}

type ErrorWithMessage struct {
	Error WithMessage `json:"error"`
}
