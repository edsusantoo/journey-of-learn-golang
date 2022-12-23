package todo

type TodoUpdateRequest struct {
	ID       int    `validate:"required" json:"id"`
	Title    string `validate:"required" json:"title"`
	Body     string `validate:"required" json:"body"`
	DateTime string `validate:"required" json:"date_time"`
}
