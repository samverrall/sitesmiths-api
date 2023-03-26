package site

type SiteName string

func NewSiteName(name string) SiteName {
	return SiteName(name)
}
