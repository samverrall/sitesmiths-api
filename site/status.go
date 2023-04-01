package site

type Status string

const (
	StatusDevelopment Status = "development"
	StatusProduction  Status = "production"
)

func (s Status) String() string {
	return string(s)
}
