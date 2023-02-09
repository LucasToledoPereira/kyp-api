package pets_commands

// command Request to create a pet
type CreatePetCommandRequest struct {
	// Pet Name
	Name        string `json:"name"`
	Description string `json:"description"`
}
