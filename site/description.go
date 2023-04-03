package site

type Description string 

func NewDescription(desc string) (Description, error) {
    return Description(desc), nil
}

func (d Description) String() string {
    return string(d)
}
