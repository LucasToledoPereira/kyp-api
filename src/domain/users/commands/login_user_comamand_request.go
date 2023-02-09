package user_commands

// command Request to create a user
type LoginUserCommandRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
