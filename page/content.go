package page

type Content string

func NewContent(c string) (Content, error) {
	return Content(c), nil
}

func (c Content) String() string {
	return string(c)
}
