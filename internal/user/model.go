package user

import "time"

type Gender string
const (
	GenderMale Gender   = "male"
	GenderFemale Gender = "female"
)

type LookingFor string
const (
	LookingMale LookingFor   = "male"
	LookingFemale LookingFor = "female"
	LookingBoth LookingFor   = "both"
)

type User struct {
	ID          string      `json:"id" db:"id"`
	Email       string      `json:"email" db:"email"`
    Password    string      `json:"-" db:"password"`
	Name        string      `json:"name" db:"name"`
	Age         int         `json:"age" db:"age"`
	Gender      Gender      `json:"gender" db:"gender"`
	Description string      `json:"description" db:"description"`
	LookingFor  LookingFor  `json:"looking_for" db:"looking_for"`
	CreatedAt   time.Time   `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time   `json:"updated_at" db:"updated_at"`
}
