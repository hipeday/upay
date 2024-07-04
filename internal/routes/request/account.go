package request

type SignInPayload struct {
	Username string `json:"username" validate:"required" message:"Username is illegal"`
	Password string `json:"password" validate:"required" message:"Password is illegal"`
}

func (s SignInPayload) Validate() error {
	return Validate(s)
}
