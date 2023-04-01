package account

type Name string

func NewName(n string) (Name, error) {
	return Name(n), nil
}

func (n Name) String() string {
	return string(n)
}
