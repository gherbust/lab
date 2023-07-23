package infrastructure

type Response struct {
	OriginalText string `json:"original_text"`
	NewText      string `json:"new_text"`
}
