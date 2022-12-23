package todo

type TodoCreateRequest struct {
	Title    string `validate:"required" json:"title"`
	Body     string `validate:"required" json:"body"`
	DateTime string `validate:"required" json:"date_time"`
}
