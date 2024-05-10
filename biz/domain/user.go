package domain

import "time"


type Gender string

const (
	Male Gender = "MALE"
	Female Gender = "FEMALE"
)

var GetGender = map[string]Gender{
	"MALE": Male,
	"FEMALE": Female,
}


type User struct {
	ID string `json:"id"`
	Username string `json:"username"`
	Email string `json:"email"`
	Dob time.Time `json:"date_of_birth"`
	Gender Gender `json:"gender"`
	Password string `json:"password"`
	CreatedTime time.Time `json:"created_time"`
	UpdatedTime time.Time `json:"updated_time"`
}




