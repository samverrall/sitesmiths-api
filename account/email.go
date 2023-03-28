package account

type Email string

func NewEmail(e string) (Email, error) {
	return Email(e), nil
}
