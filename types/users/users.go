package users // import "github.com/BryanKMorrow/aqua-sdk-go/types/users"

// User represents a local Aqua user
type User struct {
	ID              string   `json:"id"` // Username
	Password        string   `json:"password,omitempty"`
	PasswordConfirm string   `json:"passwordConfirm,omitempty"`
	Roles           []string `json:"roles,omitempty"`
	Name            string   `json:"name,omitempty"` // Display Name
	Email           string   `json:"email,omitempty"`
	FirstTime       bool     `json:"first_time,omitempty"`
}
