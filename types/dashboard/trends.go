package dashboard // import "github.com/BryanKMorrow/aqua-sdk-go/types/dashboard"

// Trends returns the v1/dashboard/<trend>/trends request
type Trends []struct {
	Date     int `json:"date"`
	Total    int `json:"total"`
	Ok       int `json:"ok"`
	High     int `json:"high"`
	Medium   int `json:"medium"`
	Low      int `json:"low"`
	Critical int `json:"critical"`
}
