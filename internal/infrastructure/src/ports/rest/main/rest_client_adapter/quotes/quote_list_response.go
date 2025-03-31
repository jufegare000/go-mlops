package quotes

type QuoteListResponse struct {
	ID             int    `json:"id"`
	LiteraryWorkID int    `json:"literary_work_id"`
	Description    string `json:"description"`
}
