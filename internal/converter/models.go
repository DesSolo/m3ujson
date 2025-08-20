package converter

type Channel struct {
	ID     uint64 `json:"id"`
	Name   string `json:"name"`
	Source string `json:"source"`
}
