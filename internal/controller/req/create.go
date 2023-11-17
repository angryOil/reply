package req

type Create struct {
	Writer  int    `json:"writer_id"`
	Content string `json:"content"`
}
