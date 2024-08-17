package requests

type CreateUser struct {
	Name  string `validate:"required, min=3, max=191"`
	Email string `validate:"required, email"`
}
