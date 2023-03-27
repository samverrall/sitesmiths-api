package site

type Name string

func NewName(name string) Name {
	return Name(name)
}

func (n Name) String() string {
	return string(n)
}
