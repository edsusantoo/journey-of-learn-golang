package todo

type TodoResponse struct {
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Body      string `json:"body"`
	DateTime  string `json:"date_time"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
