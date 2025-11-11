package user

import "time"

type CreateUserRequest struct {
	Email       string     `json:"email"`
	Password    string     `json:"password"`
	Name        string     `json:"name"`
	Age         int        `json:"age"`
	Gender      Gender     `json:"gender"`
	Description string     `json:"description"`
	LookingFor  LookingFor `json:"looking_for"`
}

func (rq *CreateUserRequest) ToUser() *User {
	return &User {
		Email:       rq.Email,
		Password:    rq.Password,
		Name:        rq.Name,
		Age:         rq.Age,
		Gender:      rq.Gender,
		Description: rq.Description,
		LookingFor:  rq.LookingFor,
	}
}

type UserResponse struct {
	ID          string     `json:"id"`
	Email       string     `json:"email"`
	Name        string     `json:"name"`
	Age         int        `json:"age"`
	Gender      Gender     `json:"gender"`
	Description string     `json:"description"`
	LookingFor  LookingFor `json:"looking_for"`
	CreatedAt   time.Time  `json:"created_at"`
}

func FromUser(user *User) *UserResponse {
	return &UserResponse {
		ID:          user.ID,
		Email:       user.Email,
		Name:        user.Name,
		Age:         user.Age,
		Gender:      user.Gender,
		Description: user.Description,
		LookingFor:  user.LookingFor,
		CreatedAt:   user.CreatedAt,
	}
}
