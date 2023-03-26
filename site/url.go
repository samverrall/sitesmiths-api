package site

type SiteURL string

func NewSiteURL(url string) SiteURL {
	return SiteURL(url)
}
