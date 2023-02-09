package user_commands

// command Request to create a user
type CreateUserCommandRequest struct {
	// User Name
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
