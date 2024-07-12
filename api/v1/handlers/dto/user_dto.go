package dto

type UpdateUserRequestDto struct {
	Email string `json:"email" binding:"required,email"`
}
