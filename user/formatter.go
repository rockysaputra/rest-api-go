package user

type UserFormatter struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Occupation string `json:"occupation"`
	Email      string `json:"email"`
	Token      string `json:"token"`
}

func FormatUser(user User) UserFormatter {
	formatter := UserFormatter{
		ID:         int(user.ID),
		Name:       user.Name,
		Occupation: user.Occupation,
		Email:      user.Email,
		Token:      user.Token,
	}
	return formatter
}
