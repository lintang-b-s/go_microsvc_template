package domain

type Gender string

const (
	Male   Gender = "Male"
	Female Gender = "Female"
)

var GetGender = map[string]Gender{
	"Male":   Male,
	"Female": Female,
}

type User struct {
	ID       uint64 `json:"id"`
	Username string `json:"user_name"`
	Email    string `json:"email"`
	Age      uint64 `json:"age"`
	Gender   Gender `json:"gender"`
	Password string `json:"password"`
	Address  string `json:"address"`
}
