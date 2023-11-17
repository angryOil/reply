package res

type GetCountList struct {
	BoardId    int `json:"board_id"`
	ReplyCount int `json:"reply_count"`
}

type CountListDto struct {
	Content []GetCountList `json:"content"`
}
