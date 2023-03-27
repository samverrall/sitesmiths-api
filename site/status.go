package site

type Status string

const (
	StatusDevelopment Status = "development"
)

func (s Status) String() string {
	return string(s)
}
