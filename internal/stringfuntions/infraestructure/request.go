package infraestructure

type Request struct {
	Action []string `json:"action"`
	Text   string   `json:"text"`
}
