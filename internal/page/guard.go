package page

import "github.com/google/uuid"

type Guard interface {
	CanCreatePage(accountID uuid.UUID)
}
