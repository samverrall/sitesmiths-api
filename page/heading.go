package page

import "errors"

type Heading string

func NewHeading(h string) (Heading, error) {
	if h == "" {
		return "", errors.New("invalid page heading supplied")
	}
	return Heading(h), nil
}
