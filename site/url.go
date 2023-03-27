package site

type URL string

func NewURL(url string) URL {
	return URL(url)
}

func (u URL) String() string {
	return string(u)
}
